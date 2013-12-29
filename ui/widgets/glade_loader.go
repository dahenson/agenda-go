package widgets

import (
	"github.com/dahenson/agenda/ui"
	"github.com/weberc2/gotk3/gtk"
	"log"
)

const (
	entryName          = "itemTextEntry"
	listStoreName      = "itemsListStore"
	windowName         = "mainWindow"
	toggleRendererName = "completeToggleRenderer"
)

var b *gtk.Builder

func init() {
	gtk.Init(nil)
	b, _ = gtk.BuilderNew()
	if err := b.AddFromString(gladestr); err != nil {
		log.Fatal(err)
	}
}

func loadEntry() *Entry {
	obj, err := b.GetObject(entryName)
	if err != nil {
		log.Fatal(err)
	}
	entry, ok := obj.(*gtk.Entry)
	if !ok {
		log.Fatalf("Failed to type-assert '%s' to gtk.Entry", entryName)
	}
	return NewEntry(entry)
}

func loadListStore() *ListStore {
	obj, err := b.GetObject(listStoreName)
	if err != nil {
		log.Fatal(err)
	}
	ls, ok := obj.(*gtk.ListStore)
	if !ok {
		log.Fatalf("Failed to type-assert '%s' to gtk.ListStore", listStoreName)
	}

	return NewListStore(ls)
}

func loadWindow() *gtk.Window {
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

func loadToggleRenderer() *gtk.CellRendererToggle {
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

func Load() ui.Ui {
	ls := loadListStore()
	toggleButton := NewToggleButton(loadToggleRenderer(), ls)
	entry := loadEntry()
	addButton := entry
	win := loadWindow()
	return ui.NewUi(entry, ls, addButton, toggleButton, func() {
		win.ShowAll()
		gtk.Main()
	})
}
