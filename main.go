package main

import (
	. "github.com/dahenson/agenda/app"
	"github.com/dahenson/agenda/itemstore"
)

func main() {
	ui := load("ui.glade")
	is := itemstore.NewItemStore("default.txt")
	app := NewApp(is, ui)
	app.LoadItems()
	ui.Run()
}
