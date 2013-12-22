package main

import (
	"github.com/conformal/gotk3/gtk"
	"github.com/dahenson/agenda/gtkui"
	"github.com/dahenson/agenda/gtkui/widgets"
	"log"
)

const (
	entryName = "itemTextEntry"
	listStoreName = "itemsListStore"
	windowName = "mainWindow"
)

func loadEntry(b *gtk.Builder) *gtk.Entry {
	obj, err := b.GetObject(entryName)
	if err != nil {
		log.Fatal(err)
	}
	entry, ok := obj.(*gtk.Entry)
	if !ok {
		log.Fatalf("Failed to type-assert '%s' to gtk.Entry", entryName)
	}
	return entry
}

func loadListStore(b *gtk.Builder) gtkui.ListStore {
	obj, err := b.GetObject(listStoreName)
	if err != nil {
		log.Fatal(err)
	}
	ls, ok := obj.(*gtk.ListStore)
	if !ok {
		log.Fatalf("Failed to type-assert '%s' to gtk.ListStore", listStoreName)
	}
	return widgets.NewListStore(ls)
}

func loadWindow(b *gtk.Builder) gtkui.Window {
	obj, err := b.GetObject(windowName)
	if err != nil {
		log.Fatal(err)
	}
	win, ok := obj.(*gtk.Window)
	if !ok {
		log.Fatalf("Failed to type-assert '%s' to gtk.Window", windowName)
	}
	win.Connect("destroy", gtk.MainQuit)
	return win
}

func load(gladeFile string) *gtkui.Ui {
	b, _ := gtk.BuilderNew()
	if err := b.AddFromFile(gladeFile); err != nil {
		log.Fatal(err)
	}
	ls := loadListStore(b)
	entry := loadEntry(b)
	win := loadWindow(b)
	return gtkui.NewUi(ls, entry, win)
}
