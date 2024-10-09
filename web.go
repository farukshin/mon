package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func startWeb() {

	http.HandleFunc("/api/sensors/add", app.apiSensorsAdd)
	http.HandleFunc("/api/sensors/edit", app.apiSensorsEdit)
	http.HandleFunc("/api/sensors/delete", app.apiSensorsDelete)
	http.HandleFunc("/api/sensors/list", app.apiSensorsList)

	http.HandleFunc("/", webHome)
	http.HandleFunc("/sensors", webSensors)
	http.HandleFunc("/sensors/", webSensors)
	http.HandleFunc("/api", webApi)

	err := http.ListenAndServe(":1616", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func webHome(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./templates/home.tmpl",
		"./templates/base.tmpl",
		"./templates/footer.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func webApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "api")
}

func webSensors(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./templates/sensors.tmpl",
		"./templates/base.tmpl",
		"./templates/footer.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, app)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
