package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("ui.glade")
	check(err)
	file, err := os.Create("glade_xml.go")
	check(err)
	defer file.Close()
	fmt.Fprintf(file, "package widgets\n\nvar gladestr = `%s`", string(data))
}
