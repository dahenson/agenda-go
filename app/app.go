package app

import (
	"github.com/dahenson/agenda/itemstore"
	"github.com/dahenson/agenda/ui"
	. "github.com/dahenson/agenda/types"
)

type App struct {
	is itemstore.ItemStore
	ui ui.Ui
}

func NewApp (is itemstore.ItemStore, ui ui.Ui) *App {
	a := &App{is: is, ui: ui}
	ui.SetAddItemCallback(a.OnAddItem)
	ui.SetToggleItemCallback(a.OnToggleItem)
	return a
}

func (a *App) OnAddItem() {
	text := a.ui.GetEntryText()
	item := NewItem(text)
	if err := a.is.AddItem(item); err != nil {
		a.ui.NotifyError("Failed to add item")
		return
	}
	a.ui.AddItem(item)
	a.ui.ClearEntryText()
}

func (a *App) OnToggleItem(id string, toggled bool) {
	if err := a.is.MarkComplete(id, !toggled); err != nil {
		a.ui.NotifyError("Oops, there seems to be a problem saving your change!")
		return
	}
	a.ui.SetToggled(id, !toggled)
}

func (a *App) LoadItems() error {
	items, err := a.is.Items()
	if err != nil {
		return err
	}

	for _, item := range items {
		a.ui.AddItem(item)
	}

	return nil
}
