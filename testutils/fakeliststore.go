package testutils

import (
	. "github.com/dahenson/agenda/types"
	"github.com/dahenson/agenda/ui"
)

type FakeListStore struct {
	items []*Item
	toggledCallback ui.ToggleItemCallback
}

func NewFakeListStore() *FakeListStore {
	return &FakeListStore{items:[]*Item{}, toggledCallback: func(_ string, _ bool) {}}
}

func (ls *FakeListStore) AddItem(item *Item) {
	ls.items = append(ls.items, item)
}

func (ls *FakeListStore) Items() []*Item {
	return ls.items
}

func (ls *FakeListStore) SetOnToggled(callback ui.ToggleItemCallback) {
	ls.toggledCallback = callback
}

func (ls *FakeListStore) Toggle(id string, toggled bool) {
	ls.SetToggled(id, toggled)
	ls.toggledCallback(id, toggled)
}

func (ls *FakeListStore) SetToggled(id string, toggled bool) {
	for _, item := range ls.items {
		if item.Id == id {
			item.Complete = toggled
		}
	}
}
