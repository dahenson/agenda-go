package app

import (
	"github.com/dahenson/agenda/itemstore"
	. "github.com/dahenson/agenda/types"
	"github.com/dahenson/agenda/ui"
	"log"
)

type App struct {
	itemstore itemstore.ItemStore
	ui        ui.Ui
}

func NewApp(is itemstore.ItemStore, ui ui.Ui) *App {
	a := &App{itemstore: is, ui: ui}
	ui.SetAddCallback(a.OnAddItem)
	ui.SetToggleCallback(a.OnToggleItem)
	return a
}

func (a *App) OnAddItem() {
	item := NewItem(a.ui.GetEntryText())
	if err := a.itemstore.Save(append(a.ui.Items(), item)); err != nil {
		a.ui.NotifyError(err)
		log.Println(err)
		return
	}
	a.ui.AddItem(item)
	a.ui.ClearEntryText()
}

func (a *App) OnToggleItem(id string, toggled bool) {
	a.ui.SetItemComplete(id, !toggled)
	if err := a.itemstore.Save(a.ui.Items()); err != nil {
		// reset UI on error
		a.ui.SetItemComplete(id, toggled)
	}
}

func (a *App) Load() {
	items, err := a.itemstore.Load()
	if err != nil {
		a.ui.NotifyError(err)
		return
	}

	for _, item := range items {
		a.ui.AddItem(item)
	}
}
