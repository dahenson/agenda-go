package widgets

import (
	"github.com/dahenson/agenda/ui"
	"github.com/conformal/gotk3/gtk"
	"log"
)

const (
	entryName          = "itemTextEntry"
	listStoreName      = "itemsListStore"
	windowName         = "mainWindow"
	toggleRendererName = "completeToggleRenderer"
)

func initBuilder() *gtk.Builder {
	b, _ := gtk.BuilderNew()
	if err := b.AddFromString(gladestr); err != nil {
		log.Fatal("Failed to add from widget from string:", gladestr, ":", err)
	}
	return b
}

func loadEntry(b *gtk.Builder) *Entry {
	obj, err := b.GetObject(entryName)
	if err != nil {
		log.Fatal("Error getting entry object:", err)
	}
	entry, ok := obj.(*gtk.Entry)
	if !ok {
		log.Fatalf("Failed to type-assert '%s' to gtk.Entry", entryName)
	}
	return NewEntry(entry)
}

func loadListStore(b *gtk.Builder) *ListStore {
	obj, err := b.GetObject(listStoreName)
	if err != nil {
		log.Fatal("Failed to get liststore object by name:", listStoreName, ":", err)
	}
	ls, ok := obj.(*gtk.ListStore)
	if !ok {
		log.Fatalf("Failed to type-assert '%s' to gtk.ListStore", listStoreName)
	}
	return NewListStore(ls)
}

func loadWindow(b *gtk.Builder) *gtk.Window {
	obj, err := b.GetObject(windowName)
	if err != nil {
		log.Fatal("Failed to get window object by name:", windowName, ":", err)
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
		log.Fatal("Failed to get CellRendererToggle object by name:", toggleRendererName, ":", err)
	}
	tog, ok := obj.(*gtk.CellRendererToggle)
	if !ok {
		log.Fatalf("Failed to type-assert '%s' to gtk.CellRendererToggle", toggleRendererName)
	}
	return tog
}

func init() {
	gtk.Init(nil)
}

func Load() ui.Ui {
	b := initBuilder()
	ls := loadListStore(b)
	toggleButton := NewToggleButton(loadToggleRenderer(b), ls)
	entry := loadEntry(b)
	addButton := entry
	win := loadWindow(b)
	return ui.NewUi(entry, ls, addButton, toggleButton, func() {
		win.ShowAll()
		gtk.Main()
	})
}
