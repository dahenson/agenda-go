package main

import (
	"strings"
)

type ItemStore interface {
	AddItem(item Item) error
	Items() ([]Item, error)
	//	RemoveItem(item Item) error
}

type FSItemStore struct {
	filename string
	items    []Item
	fs       fileSystem
}

func NewFSItemStore(filename string) *FSItemStore {
	path := getPath()
	return &FSItemStore{filename: path + filename, items: make([]Item, 0), fs: fs}
}

func (is *FSItemStore) AddItem(item Item) error {
	if _, err := is.fs.Append(is.filename, item.Text()+"\n"); err != nil {
		return err
	}
	return nil
}

func (is *FSItemStore) Items() ([]Item, error) {
	data, err := is.fs.ReadFile(is.filename)
	if is.fs.IsNotExist(err) {
		return []Item{}, nil
	}
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	items := []Item{}
	for _, line := range lines {
		if len(line) > 0 {
			items = append(items, NewItem(line))
		}
	}
	return items, nil
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
		if is.items[i].Id() == item.Id() {
			is.items = append(is.items[:i-1], is.items[i+1:]...)
		}
	}
	return nil
}

func (is *InMemoryItemStore) Items() ([]Item, error) {
	return is.items, nil
}

func NewInMemoryItemStore() *InMemoryItemStore {
	return new(InMemoryItemStore)
}
