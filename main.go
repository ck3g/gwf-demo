package main

import "github.com/ck3g/gwf"

type application struct {
	App *gwf.GWF
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
