package main

import (
	"log"
)

func main() {
	is, err := NewFSItemStore("default.txt")
	if err != nil {
		log.Fatal(err)
	}
	app, err := NewApp("ui.glade", is, 300, 400)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
