package main

import (
	"os"
	"log"
	. "github.com/dahenson/agenda/app"
	"github.com/dahenson/agenda/itemstore"
	. "github.com/dahenson/agenda/fs"
)

func main() {
	path := GetPath()
	if err := os.MkdirAll(path, 0744); err != nil {
		log.Fatal(err)
	}

	ui := load("ui.glade")
	is := itemstore.NewItemStore(path + "default.txt")
	app := NewApp(is, ui)
	app.LoadItems()
	ui.Run()
}
