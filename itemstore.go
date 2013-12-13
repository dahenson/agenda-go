package main

import ()

type ItemStore interface {
	AddItem(item Item) error
	RemoveItem(item Item) error
}

type InMemoryItemStore struct {
	items []Item
}

func (is *InMemoryItemStore) AddItem(item Item) error {
	is.items = append(is.items, item)
	return nil
}

func (is *InMemoryItemStore) RemoveItem(item Item) error {
	for i := 0; i < len(is.items); i++ {
		if is.items[i].Text() == item.Text() {
			is.items = append(is.items[:i-1], is.items[i+1:]...)
		}
	}
	return nil
}

func NewInMemoryItemStore() *InMemoryItemStore {
	return new(InMemoryItemStore)
}
