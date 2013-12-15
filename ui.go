package main

import (
	"fmt"
	"github.com/conformal/gotk3/gtk"
	"log"
)

func checkTypeAssertion(objName string, ok bool) error {
	if !ok {
		return fmt.Errorf("Object '%s' not found in ui file", objName)
	}
	return nil
}

func mainWindow(builder *gtk.Builder, mainWindowName string, w, h int) (*gtk.Window, error) {
	obj, err := builder.GetObject(mainWindowName)
	if err != nil {
		return nil, err
	}
	win, ok := (obj).(*gtk.Window)
	if err := checkTypeAssertion(mainWindowName, ok); err != nil {
		return nil, err
	}
	win.Connect("destroy", gtk.MainQuit)
	win.SetDefaultSize(w, h)
	win.SetPosition(gtk.WIN_POS_CENTER)
	return win, nil
}

func itemTextEntry(builder *gtk.Builder, objName string) (*gtk.Entry, error) {
	obj, err := builder.GetObject(objName)
	if err != nil {
		return nil, err
	}
	entry, ok := (obj).(*gtk.Entry)
	if err := checkTypeAssertion(objName, ok); err != nil {
		return nil, err
	}
	return entry, nil
}

func itemsListStore(builder *gtk.Builder, itemsListStoreName string) (*gtk.ListStore, error) {
	obj, err := builder.GetObject(itemsListStoreName)
	if err != nil {
		return nil, err
	}
	ls, ok := (obj).(*gtk.ListStore)
	if err := checkTypeAssertion(itemsListStoreName, ok); err != nil {
		return nil, err
	}
	return ls, nil
}

type App struct {
	mainWindow     *gtk.Window
	itemTextEntry  *gtk.Entry
	itemsListStore *gtk.ListStore
	store          ItemStore
}

func (a *App) Run() {
	a.mainWindow.ShowAll()
	gtk.Main()
}

// Grabs the contents of the addItemEntry, creates a new Item, and represents that item
// with a new widget.
func (a *App) AddItem() {
	text, _ := a.itemTextEntry.GetText()
	item := NewItem(text)
	if err := a.store.AddItem(item); err != nil {
		log.Println("ItemStore failed to add item:", item, "err:", err)
		return
		// TODO: Notify user if ItemStore fails to add item
	}
	a.AddToTreeView(item)
	a.itemTextEntry.SetText("")
	// TODO make a TreeViewItem or whatever they're called and add it to the treeview
}

func (a *App) AddToTreeView(item Item) {
	var iter gtk.TreeIter
	a.itemsListStore.Append(&iter)
	a.itemsListStore.Set(&iter, []int{0}, []interface{}{item.Text()})
}

func (a *App) LoadItems() error {
	items, err := a.store.Items()
	if err != nil {
		return err
	}

	for _, item := range items {
		a.AddToTreeView(item)
	}
	return nil
}

func NewApp(uiFileName string, store ItemStore, defaultWidth, defaultHeight int) (*App, error) {
	gtk.Init(nil)
	builder, _ := gtk.BuilderNew()
	if err := builder.AddFromFile(uiFileName); err != nil {
		return nil, err
	}

	win, err := mainWindow(builder, "mainWindow", defaultWidth, defaultHeight)
	if err != nil {
		return nil, err
	}

	entry, err := itemTextEntry(builder, "itemTextEntry")
	if err != nil {
		return nil, err
	}

	ls, err := itemsListStore(builder, "itemsListStore")
	if err != nil {
		return nil, err
	}

	a := &App{
		mainWindow:     win,
		itemTextEntry:  entry,
		itemsListStore: ls,
		store:          store,
	}

	entry.Connect("activate", a.AddItem)
	entry.Connect("icon-release", a.AddItem)

	if err := a.LoadItems(); err != nil {
		return nil, err
	}

	return a, nil
}
