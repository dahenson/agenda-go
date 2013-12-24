package ui

import (
	. "github.com/dahenson/agenda/types"
)

type AddItemCallback func()
type ToggleItemCallback func(id string, complete bool)

type Ui interface {
	SetAddItemCallback(callback AddItemCallback)
	SetToggleItemCallback(callback ToggleItemCallback)
	GetEntryText() string
	ClearEntryText()
	AddItem(item *Item)
	SetToggled(id string, toggled bool)
	NotifyError(msg string)
	Run()
}
