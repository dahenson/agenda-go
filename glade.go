package main

import (
	"github.com/dahenson/agenda/gtkui"
	"github.com/dahenson/agenda/gtkui/widgets"
	"github.com/weberc2/gotk3/gtk"
	"log"
)

const (
	entryName          = "itemTextEntry"
	listStoreName      = "itemsListStore"
	windowName         = "mainWindow"
	toggleRendererName = "completeToggleRenderer"
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

	renderer := loadToggleRenderer(b)
	return widgets.NewListStore(ls, renderer)
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

func loadToggleRenderer(b *gtk.Builder) *gtk.CellRendererToggle {
	obj, err := b.GetObject(toggleRendererName)
	if err != nil {
		log.Fatal(err)
	}
	tog, ok := obj.(*gtk.CellRendererToggle)
	if !ok {
		log.Fatalf("Failed to type-assert '%s' to gtk.CellRendererToggle", toggleRendererName)
	}
	return tog
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
