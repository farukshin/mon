package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	version       string
	ConfCatalog   string
	ConfFile      string
	DataCatalog   string
	LogCatalog    string
	Sensors       []sensor
	Notifications []notification
}

func (app *application) run() {

	if len(os.Args) < 1 || isArgs("--help") || isArgs("-h") {
		help_home()
	} else if isArgs("--version") || isArgs("-v") {
		getVersion()
	} else if os.Args[1] == "start" {
		start()
	} else if os.Args[1] == "sensors" {
		cliSensors()
	} else if os.Args[1] == "notify" {
		cliNotify()
	} else {
		help_home()
	}
}

func isArgsAll(ar string) bool {
	mas := strings.Split(ar, ",")
	res := true
	for _, a := range mas {
		res = res && isArgs(a)
	}
	return res
}

func start() {

	app.init()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		app.close()
		os.Exit(1)
	}()

	jobs := make(chan *sensor)
	results := make(chan sensorResultMessage)

	go startWorker(jobs, results)
	go startTimer(jobs)
	go startResultsCalcer(results)
	go startWeb()

	for true {
		time.Sleep(time.Duration(5) * time.Second)
	}

}

func getVersion() {
	fmt.Printf("mon %s\n", app.version)
}

func isArgs(a1 string) bool {

	_, err := getArgs(a1)
	return err == nil

}

func getArgs(a1 string) (string, error) {

	for _, s := range os.Args[1:] {
		if s == a1 {
			return "", nil
		}
		for i := 0; i < len(s); i++ {
			if s[i] == '=' && i > 0 {
				v := s[:i]
				if v == a1 {
					return s[i+1:], nil
				}
			}
		}

	}
	return "", errors.New("Не найдено флага " + a1)
}

func help_home() {
	fmt.Println(helpHomeStr())
}

func helpHomeStr() string {
	var sb strings.Builder
	sb.WriteString("Приложение: mon\n")
	sb.WriteString("\tСистема мониторинга\n\n")

	sb.WriteString("Строка запуска: mon [КОМАНДА] [Опции]\n\n")
	sb.WriteString("Доступные команды:\n")
	sb.WriteString("  start - запуск системы мониторинга\n\n")

	sb.WriteString("Опции:\n")
	sb.WriteString("-h --help - вызов справки\n")
	sb.WriteString("-v --version - версия приложения\n\n")

	sb.WriteString("Пример запуска:\n")
	sb.WriteString("./mon --version\n")
	sb.WriteString("./mon start")
	return sb.String()
}

func startResultsCalcer(results <-chan sensorResultMessage) {
	for r := range results {
		sensor := app.getSensorByUID(r.sensorUID)
		sensor.History = append(sensor.History, r.result)
		sensor.calcNewTic()
		//curSuccess := sensor.check(&r)
		//log := fmt.Sprintf("%s Job %s target = %s result = %s", successStr(curSuccess), sensor.name, sensor.target, r.result)
		//if sensor.successStatus != curSuccess {
		//sendMessageTG(log)
		//}
		//sensor.successStatus = curSuccess
		//println(log)
	}
}

func successStr(success bool) string {
	if success {
		return "SUCCESS"
	}
	return "FAIL"
}

func startTimer(jobs chan<- *sensor) {
	for true {
		for _, r := range app.Sensors {
			jobs <- &r
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func startWorker(jobs <-chan *sensor, results chan<- sensorResultMessage) {
	for j := range jobs {
		results <- checkJob(j)
	}
}

func testSensors() []sensor {
	res := []sensor{}
	sen1 := sensor{
		UID:    "cc1c0cb3-3d53-48c9-af84-7bfff76f0f1d",
		Name:   "test1",
		Kind:   "httpcode",
		Target: "http://localhost:1717",
		Time:   60,
		Life:   1000,
		Expect: sensorResult{
			resInt: 200,
		},
	}
	sen2 := sensor{
		UID:    "01dc579c-0884-4fe1-902a-1571a5c4bd85",
		Name:   "test2",
		Kind:   "httpcode",
		Target: "http://farukshin.com",
		Time:   60,
		Life:   1000,
		Expect: sensorResult{
			resInt: 200,
		},
	}
	res = append(res, sen1)
	res = append(res, sen2)
	return res

}

func createAppCatalogs() error {

	cat := []string{app.ConfCatalog, app.DataCatalog, app.LogCatalog}

	for _, catalog := range cat {
		if _, err := os.Stat(catalog); err != nil {
			if os.IsNotExist(err) {
				err = os.Mkdir(catalog, 0777)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}

func (app *application) close() {
	app.saveConfig()
}
