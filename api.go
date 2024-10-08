package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func (app *application) apiSensorsAdd(w http.ResponseWriter, r *http.Request) {

	var p sensor
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Fprintf(w, "{ok:false}")
		return
	}
	app.addSensor(&p)
	json := "{ok:true, uid:" + p.UID + "}"
	fmt.Fprintf(w, json)

}

func (app *application) apiSensorsEdit(w http.ResponseWriter, r *http.Request) {

	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprintf(w, "{ok:false}")
		return
	}
	_uid, ok := data["uid"]
	if !ok {
		fmt.Fprintf(w, "{ok:false}")
		return
	}
	s := app.getSensorByUID(_uid)
	if s == nil {
		fmt.Fprintf(w, "{ok:false}")
		return
	}
	for key, value := range data {
		if key == "target" {
			reflect.ValueOf(s).Elem().FieldByName("Target").SetString(value)
		}
	}
	json := "{ok:true, uid:" + s.UID + "}"
	fmt.Fprintf(w, json)
}

func (app *application) apiSensorsDelete(w http.ResponseWriter, r *http.Request) {

	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprintf(w, "{ok:false}")
		return
	}
	_uid, ok := data["uid"]
	if !ok {
		fmt.Fprintf(w, "{ok:false}")
		return
	}

	err = app.deleteSensorByID(_uid)
	if err != nil {
		fmt.Fprintf(w, "{ok:false}")
		return
	}

	json := "{ok:true}"
	fmt.Fprintf(w, json)
}
