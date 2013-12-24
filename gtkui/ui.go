package gtkui

import (
	. "github.com/dahenson/agenda/types"
	"github.com/dahenson/agenda/ui"
	"github.com/weberc2/gotk3/gtk"
)

type ListStore interface {
	AddItem(item *Item)
	SetOnToggled(callback ui.ToggleItemCallback)
	SetToggled(id string, toggled bool)
}

type Window interface {
	ShowAll()
}

type Ui struct {
	liststore ListStore
	entry     *gtk.Entry
	win       Window
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

func (ui *Ui) SetToggleItemCallback(callback ui.ToggleItemCallback) {
	ui.liststore.SetOnToggled(callback)
}

// Updates the UI of the item identified by 'id' with the state 'toggled';
// Does not trigger the OnToggled callback
func (ui *Ui) SetToggled(id string, toggled bool) {
	ui.liststore.SetToggled(id, toggled)
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
