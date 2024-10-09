package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func cliSensors() {
	if os.Args[2] == "list" {
		cliSensorsList()
	} else if os.Args[2] == "add" {
		cliSensorsAdd()
	} else if os.Args[2] == "delete" {
		cliSensorsDelete()
	} else if os.Args[2] == "edit" {
		cliSensorsEdit()
	} else {
		sensorsHome()
	}
}

func argsToJSON(arg []string) string {
	var sb strings.Builder
	//sb.WriteString("{")
	for _, s := range arg {
		ss := strings.Split(s, "=")
		if len(ss) == 2 {
			s0 := ss[0]
			s0 = strings.Replace(s0, "-", "", -1)
			if sb.Len() != 0 {
				sb.WriteString(",")
			}
			sb.WriteString(fmt.Sprintf("\"%s\":\"%s\"", s0, ss[1]))
		}
	}

	return fmt.Sprintf(`{%s}`, sb.String())
	//return `{"foo":"bar"}`
}

func cliSensorsAdd() {
	j := argsToJSON(os.Args[3:])
	data := []byte(j)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:1616/api/sensors/add", "application/json", r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(bodyBytes))
	}
}

func cliSensorsEdit() {
	j := argsToJSON(os.Args[3:])
	data := []byte(j)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:1616/api/sensors/edit", "application/json", r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(bodyBytes))
	}
}

func cliSensorsDelete() {
	j := argsToJSON(os.Args[3:])
	data := []byte(j)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:1616/api/sensors/delete", "application/json", r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(bodyBytes))
	}
}

func cliSensorsList() {
	fmt.Println(cliSensorsListStr())
}

func sensorsHome() {
	fmt.Println(sensorsHomeStr())
}

func sensorsHomeStr() string {
	var sb strings.Builder
	sb.WriteString("Приложение: mon\n")
	sb.WriteString("\tСистема мониторинга\n\n")

	sb.WriteString("Строка запуска: mon sensors [Опции]\n\n")

	sb.WriteString("Опции:\n")
	sb.WriteString("list - список всех сенсоров\n")

	sb.WriteString("Пример запуска:\n")
	sb.WriteString("./mon sensors list\n")
	return sb.String()
}

func cliSensorsListStr() string {
	resp, err := http.Get("http://localhost:1616/api/sensors/list")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return ""
		}
		var data []sensor
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			return ""
		}
		var sb strings.Builder
		for _, s := range data {
			sb.WriteString(fmt.Sprintf("%s %s %s\n", s.UID, s.Kind, s.Target))
		}
		return sb.String()
	}
	return ""
}
