package main

import (
	"demoapp/data"
	"demoapp/handlers"

	"github.com/ck3g/gwf"
)

type application struct {
	App      *gwf.GWF
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
