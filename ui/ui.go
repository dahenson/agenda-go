package ui

import (
	. "github.com/dahenson/agenda/types"
	. "github.com/dahenson/agenda/ui/uicallbacks"
)

type Ui interface {
	GetEntryText() string
	ClearEntryText()
	Items() []*Item
	Len() int
	AddItem(item *Item)
	SetAddCallback(callback AddCallback)
	SetItemComplete(id string, complete bool) error
	SetToggleCallback(callback ToggleCallback)
	NotifyError(v interface{})
	RemoveItem(id string) error
	Run()
}

type ListStore interface {
	Items() []*Item
	AddItem(item *Item)
	SetItemComplete(itemId string, complete bool) error
	RemoveItem(id string) error
	Len() int
}

type Entry interface {
	SetText(text string)
	GetText() string
}

type AddButton interface {
	SetCallback(callback AddCallback)
}

type ToggleButton interface {
	SetCallback(callback ToggleCallback)
}

type ui struct {
	entry Entry
	ListStore
	addButton AddButton
	toggleButton ToggleButton
	runFn func()
}

func NewUi(entry Entry, liststore ListStore, addButton AddButton, toggleButton ToggleButton, runfn func()) Ui {
	return &ui{
		entry: entry,
		ListStore: liststore,
		addButton: addButton,
		toggleButton: toggleButton,
		runFn: runfn,
	}
}

func (ui *ui) GetEntryText() string {
	return ui.entry.GetText()
}

func (ui *ui) ClearEntryText() {
	ui.entry.SetText("")
}

func (ui *ui) AddItem(item *Item) {
	ui.ListStore.AddItem(item)
}

func (ui *ui) SetAddCallback(callback AddCallback) {
	ui.addButton.SetCallback(callback)
}

func (ui *ui) SetToggleCallback(callback ToggleCallback) {
	ui.toggleButton.SetCallback(callback)
}

func (ui *ui) SetItemComplete(itemId string, complete bool) error {
	if err := ui.ListStore.SetItemComplete(itemId, complete); err != nil {
		return err
	}
	return nil
}

func (ui *ui) NotifyError(v interface{}) {
	// TODO
}

func (ui *ui) Run() {
	ui.runFn()
}
