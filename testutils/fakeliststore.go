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
func (ls *FakeListStore) AddItem(item *Item) { ls.items = append(ls.items, item) }
func (ls *FakeListStore) Get(itemId string) *Item {
	for _, item := range ls.items {
		if item.Id == itemId {
			return item
		}
	}
	return nil
}
func (ls *FakeListStore) SetItemComplete(itemId string, complete bool) error {
	item := ls.Get(itemId)
	if item == nil {
		return fmt.Errorf("Item '%s' not found", itemId)
	}
	item.Complete = complete
	return nil
}
