package main

import (
	"demoapp/handlers"

	"github.com/ck3g/gwf"
)

type application struct {
	App      *gwf.GWF
	Handlers *handlers.Handlers
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
