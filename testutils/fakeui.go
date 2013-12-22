package testutils

import (
	"github.com/dahenson/agenda/ui"
	. "github.com/dahenson/agenda/types"
)

type FakeUi struct {
	entryText string
	addItemCallback ui.AddItemCallback
	items []*Item
	errMsg string
}

func NewFakeUi() *FakeUi {
	ui := new(FakeUi)
	ui.addItemCallback = func(){}
	ui.items = make([]*Item, 0)
	return ui
}

func (ui *FakeUi) ClearEntryText() {
	ui.entryText = ""
}

func (ui *FakeUi) GetEntryText() string {
	return ui.entryText
}

func (ui *FakeUi) SetEntryText(newText string) {
	ui.entryText = newText
}

func (ui *FakeUi) PressAddButton() {
	ui.addItemCallback()
}

func (ui *FakeUi) SetAddItemCallback(callback ui.AddItemCallback) {
	ui.addItemCallback = callback
}

func (ui *FakeUi) AddItem(item *Item) {
	ui.items = append(ui.items, item)
}

func (ui *FakeUi) Items() []*Item {
	return ui.items
}

func (ui *FakeUi) CurrentErrorMessage() string {
	return ui.errMsg
}

func (ui *FakeUi) NotifyError(msg string) {
	ui.errMsg = msg
}

func (ui *FakeUi) Run() {}
