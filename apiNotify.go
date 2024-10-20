package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) apiNotifyAdd(w http.ResponseWriter, r *http.Request) {

	var p notification
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Fprintf(w, "{ok:false}")
		return
	}
	app.addNotify(&p)
	json := "{ok:true, uid:" + p.UID + "}"
	fmt.Fprintf(w, json)

}

func (app *application) apiNotifyList(w http.ResponseWriter, r *http.Request) {

	var jsonData []byte
	jsonData, err := json.Marshal(app.Notifications)
	if err != nil {
		fmt.Fprintf(w, "{ok:false}")
		return
	}
	fmt.Fprintf(w, string(jsonData))
}

func (app *application) apiNotifyEdit(w http.ResponseWriter, r *http.Request) {

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
	s := app.getNotifyByUID(_uid)
	if s == nil {
		fmt.Fprintf(w, "{ok:false}")
		return
	}
	for key, value := range data {
		if key == "name" {
			s.Name = value
		}
	}
	json := "{ok:true, uid:" + s.UID + "}"
	fmt.Fprintf(w, json)
}

func (app *application) apiNotifyDelete(w http.ResponseWriter, r *http.Request) {

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

	err = app.deleteNotifyByID(_uid)
	if err != nil {
		fmt.Fprintf(w, "{ok:false}")
		return
	}

	json := "{ok:true}"
	fmt.Fprintf(w, json)
}
