package main

import (
	"demoapp/data"
	"demoapp/handlers"
	"log"
	"os"

	"github.com/ck3g/gwf"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init GWF
	g := &gwf.GWF{}
	err = g.New(path)
	if err != nil {
		log.Fatal(err)
	}

	g.AppName = "demoapp"

	myHandlers := &handlers.Handlers{
		App: g,
	}

	app := &application{
		App:      g,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)

	return app
}
