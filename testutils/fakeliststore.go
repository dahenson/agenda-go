package testutils

import (
	. "github.com/dahenson/agenda/types"
)

type FakeListStore struct {
	items []*Item
}

func NewFakeListStore() *FakeListStore {
	return &FakeListStore{make([]*Item, 0)}
}

func (ls *FakeListStore) AddItem(item *Item) {
	ls.items = append(ls.items, item)
}

func (ls *FakeListStore) Items() []*Item {
	return ls.items
}
