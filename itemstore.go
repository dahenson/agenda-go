package main

import (
	"encoding/json"
)

type ItemStore interface {
	AddItem(item Item) error
	Items() ([]Item, error)
}

type itemStore struct {
	filename string
	fs fileSystem
	items []Item
}

func NewItemStore(filename string) ItemStore {
	return &itemStore{filename: filename, fs: fs, items: []Item{}}
}

func (is *itemStore) Flush(items []Item) error {
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}

	if _, err := is.fs.Write(is.filename, data); err != nil {
		return err
	}
	return nil
}

func (is *itemStore) AddItem(item Item) error {
	// in case writing to the file fails, we don't want to update our in-memory model
	// thus getting out of sync with the file's state, so we handle any filesystem errors
	// before updating our internal model
	temp := append(is.items, item)
	if err := is.Flush(temp); err != nil {
		return err
	}
	is.items = temp
	return nil
}

func (is *itemStore) Load() error {
	data, err := is.fs.ReadFile(is.filename)
	if is.fs.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &is.items)
}

// returns the contents of the itemStore's file or an error if there are filesystem problems
func (is *itemStore) Items() ([]Item, error) {
	if err := is.Load(); err != nil {
		return nil, err
	}
	return is.items, nil
}
