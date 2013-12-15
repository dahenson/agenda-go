package main

import (
	"log"
)

func main() {
	app, err := NewApp("ui.glade", NewItemStore("default.txt"), 300, 400)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
