package testutils

import (
	"fmt"
	. "github.com/dahenson/agenda/types"
)

type FakeListStore struct {
	items []*Item
}

func NewFakeListStore() *FakeListStore       { return &FakeListStore{items: []*Item{}} }
func (ls *FakeListStore) Items() []*Item     { return ls.items }
func (ls *FakeListStore) Len() int           { return len(ls.items) }
func (ls *FakeListStore) AddItem(item *Item) { ls.items = append(ls.items, item) }
func (ls *FakeListStore) getIndex(itemId string) (int, error) {
	for i, item := range ls.items {
		if item.Id() == itemId {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Item id not found: %s", itemId)
}
func (ls *FakeListStore) Get(itemId string) (*Item, error) {
	i, err := ls.getIndex(itemId)
	if err != nil {
		return nil, err
	}
	return ls.items[i], nil
}
func (ls *FakeListStore) SetItemComplete(itemId string, complete bool) error {
	item, err := ls.Get(itemId)
	if err != nil {
		return err
	}
	item.SetComplete(complete)
	return nil
}

func (ls *FakeListStore) RemoveItem(id string) error {
	i, err := ls.getIndex(id)
	if err != nil {
		return err
	}

	ls.items = append(ls.items[:i], ls.items[i+1:]...)
	return nil
}
