package main

import (
	"fmt"
	"github.com/conformal/gotk3/glib"
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

func addItemButton(builder *gtk.Builder, addItemButtonName string) (*gtk.Button, error) {
	obj, err := builder.GetObject(addItemButtonName)
	if err != nil {
		return nil, err
	}
	btn, ok := (obj).(*gtk.Button)
	if err := checkTypeAssertion(addItemButtonName, ok); err != nil {
		return nil, err
	}
	return btn, nil
}

func itemsListStore(builder *gtk.Builder, itemsTreeViewName string) (*gtk.ListStore, error) {
	obj, err := builder.GetObject(itemsTreeViewName)
	if err != nil {
		return nil, err
	}
	tv, ok := (obj).(*gtk.TreeView)
	if err := checkTypeAssertion(itemsTreeViewName, ok); err != nil {
		return nil, err
	}

	renderer, _ := gtk.CellRendererTextNew()
	col, _ := gtk.TreeViewColumnNewWithAttribute("Item", renderer, "text", 0)
	tv.AppendColumn(col)

	ls, _ := gtk.ListStoreNew(glib.TYPE_STRING)
	tv.SetModel(ls)
	return ls, nil
}

type App struct {
	mainWindow     *gtk.Window
	itemTextEntry  *gtk.Entry
	addItemButton  *gtk.Button
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
		log.Println("ItemStore failed to add item:", item)
		// TODO: Notify user if ItemStore fails to add item
	}
	var iter gtk.TreeIter
	a.itemsListStore.Append(&iter)
	a.itemsListStore.Set(&iter, []int{0}, []interface{}{text})
	// TODO make a TreeViewItem or whatever they're called and add it to the treeview
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

	btn, err := addItemButton(builder, "addItemButton")
	if err != nil {
		return nil, err
	}

	ls, err := itemsListStore(builder, "itemsTreeView")
	if err != nil {
		return nil, err
	}

	a := &App{
		mainWindow:     win,
		itemTextEntry:  entry,
		addItemButton:  btn,
		itemsListStore: ls,
		store:          store,
	}

	btn.Connect("clicked", a.AddItem)

	return a, nil
}
