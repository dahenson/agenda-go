package ui

import (
	. "github.com/dahenson/agenda/types"
)

type AddItemCallback func()

type Ui interface {
	SetAddItemCallback(callback AddItemCallback)
	GetEntryText() string
	ClearEntryText()
	AddItem(item *Item)
	NotifyError(msg string)
	Run()
}
