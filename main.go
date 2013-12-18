package main

import (
	"log"
	"os"
)

func main() {
	path := getPath()
	if err := os.MkdirAll(path, 0744); err != nil {
		log.Fatal(err)
	}
	app, err := NewApp("ui.glade", NewFSItemStore(path + "default.txt"), 300, 400)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
