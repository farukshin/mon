package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configs struct {
	Version string   `json:"version"`
	Sensors []sensor `json:"sensors"`
}

func (app *application) init() error {

	app.DataCatalog = "./data/"
	app.LogCatalog = "./logs/"
	app.ConfCatalog = "./conf/"
	app.ConfFile = "./conf/conf.json"

	err := createAppCatalogs()
	if err != nil {
		return err
	}
	fInfo, err := os.OpenFile("logs/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer fInfo.Close()

	fError, err := os.OpenFile("logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer fError.Close()

	infoLog := log.New(fInfo, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(fError, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.infoLog = infoLog
	app.errorLog = errorLog

	/*
		t := app.testSensors()
		for _, sensor := range t {
			app.Sensors = append(app.Sensors, sensor)
		}
	*/

	if _, err := os.Stat(app.ConfFile); err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(app.ConfFile)
			if err != nil {
				return err
			}
			app.saveConfig()
		} else {
			return err
		}
	} else {
		app.loadConfig()
	}
	return nil
}

func (app *application) saveConfig() error {
	conf := Configs{}
	conf.Version = app.version
	for _, sen := range app.Sensors {
		conf.Sensors = append(conf.Sensors, sen)
	}
	var jsonData []byte
	jsonData, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(app.ConfFile, jsonData, 0644)
	return err
}

func (app *application) loadConfig() error {
	data, err := ioutil.ReadFile(app.ConfFile)
	if err != nil {
		return err
	}
	conf := Configs{}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	for _, s := range conf.Sensors {
		app.Sensors = append(app.Sensors, s)
	}
	return nil
}
