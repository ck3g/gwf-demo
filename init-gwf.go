package main

import (
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

	g.InfoLog.Println("Debug is set to ", g.Debug)

	app := &application{
		App: g,
	}

	return app
}
