package main

var app application

func main() {

	app := &application{
		version: "v0.0.1",
	}

	app.run()
}
