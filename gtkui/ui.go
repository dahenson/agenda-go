package gtkui

import (
	. "github.com/dahenson/agenda/types"
	"github.com/dahenson/agenda/ui"
	"github.com/conformal/gotk3/gtk"
)

type ListStore interface {
	AddItem(item *Item)
}

type Window interface {
	ShowAll()
}

type Ui struct {
	liststore ListStore
	entry *gtk.Entry
	win Window
}

func init() {
	gtk.Init(nil)
}

func NewUi(ls ListStore, entry *gtk.Entry, win Window) *Ui {
	return &Ui{liststore: ls, entry: entry, win: win}
}

func (ui *Ui) AddItem(item *Item) {
	ui.liststore.AddItem(item)
}

func (ui *Ui) GetEntryText() string {
	text, _ := ui.entry.GetText()
	return text
}

func (ui *Ui) SetAddItemCallback(callback ui.AddItemCallback) {
	ui.entry.Connect("activate", callback)
	ui.entry.Connect("icon-release", callback)
}

func (ui *Ui) ClearEntryText() {
	ui.entry.SetText("")
}

func (ui *Ui) NotifyError(msg string) {
	// TODO implement me
}

func (ui *Ui) Run() {
	ui.win.ShowAll()
	gtk.Main()
}
