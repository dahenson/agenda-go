package main

import (
	"os"
	"log"
	. "github.com/dahenson/agenda/app"
	"github.com/dahenson/agenda/itemstore"
	. "github.com/dahenson/agenda/fs"
	"github.com/dahenson/agenda/ui/widgets"
)

func main() {
	path := GetPath()
	if err := os.MkdirAll(path, 0744); err != nil {
		log.Fatal(err)
	}

	ui := widgets.Load()
	is := itemstore.NewItemStore(path + "default.txt")
	app := NewApp(is, ui, 5)
	app.Load()
	ui.Run()
}
