package testutils

import (
	"github.com/dahenson/agenda/fs"
	"github.com/dahenson/agenda/itemstore"
	. "github.com/dahenson/agenda/types"
	"fmt"
)

type FakeItemStore struct {
	itemstore.ItemStore
	addItemErrorText string
}

func NewFakeItemStore(filename string) *FakeItemStore {
	return &FakeItemStore{
		ItemStore: itemstore.NewItemStoreWithFileSystem(filename, fs.NewFakeFS()),
	}
}

func (is *FakeItemStore) ReturnErrorOnAddItem(errText string) {
	is.addItemErrorText = errText
}

func (is *FakeItemStore) AddItem(item *Item) error {
	if is.addItemErrorText == "" {
		return is.ItemStore.AddItem(item)
	}
	return fmt.Errorf(is.addItemErrorText)
}
