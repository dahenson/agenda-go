package app

import (
	"github.com/dahenson/agenda/itemstore"
	. "github.com/dahenson/agenda/types"
	"github.com/dahenson/agenda/ui"
	"log"
)

type App struct {
	itemstore         itemstore.ItemStore
	ui                ui.Ui
	maxCompletedItems int
}

func NewApp(is itemstore.ItemStore, ui ui.Ui, maxCompletedItems int) *App {
	a := &App{itemstore: is, ui: ui, maxCompletedItems: maxCompletedItems}
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

func findFirstCompletedItem(items []*Item) *Item {
	var firstCompleted *Item = nil
	for _, item := range items {
		if item.Complete() {
			if firstCompleted == nil || item.CompletedBefore(firstCompleted) {
				firstCompleted = item
			}
		}
	}
	return firstCompleted
}

func (a *App) removeFirstCompletedItem() error {
	// find first completed item
	firstCompletedItem := findFirstCompletedItem(a.ui.Items())

	// if `firstCompletedItem` is nil, then there are no completed items, so do nothing
	if firstCompletedItem == nil {
		return nil
	}
	// remove it from model
	if err := a.ui.RemoveItem(firstCompletedItem.Id()); err != nil {
		return err
	}
	// save the itemstore
	return a.itemstore.Save(a.ui.Items())
}

func (a *App) countCompletedItems() int {
	count := 0
	for _, item := range a.ui.Items() {
		if item.Complete() {
			count++
		}
	}
	return count
}

func (a *App) fewerThanMaxCompletedItems() bool {
	return a.countCompletedItems() < a.maxCompletedItems
}

func (a *App) OnToggleItem(id string, toggled bool) {
	// if we're toggling from untoggled to toggled and the current number of
	// completed items is >= maxCompletedItems, remove the first completed item
	if !toggled && !a.fewerThanMaxCompletedItems() {
		if err := a.removeFirstCompletedItem(); err != nil {
			a.ui.NotifyError(err)
		}
	}
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
