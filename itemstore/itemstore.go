package itemstore

import (
	"encoding/json"
	. "github.com/dahenson/agenda/types"
	"github.com/dahenson/agenda/fs"
)

// ItemStore exists to persist a memory model
type ItemStore interface {
	Save(items []*Item) error
	Load() ([]*Item, error)
}

type itemStore struct {
	filename string
	fs fs.FileSystem
}

func NewItemStore(filename string) ItemStore {
	return &itemStore{filename: filename, fs: fs.OsFs()}
}

func NewItemStoreWithFileSystem(filename string, filesys fs.FileSystem) ItemStore {
	return &itemStore{filename: filename, fs: filesys}
}

func (is *itemStore) Save(items []*Item) error {
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}

	if _, err := is.fs.Write(is.filename, data); err != nil {
		return err
	}

	return nil
}

func (is *itemStore) Load() ([]*Item, error) {
	data, err := is.fs.ReadFile(is.filename)
	if err != nil {
		return nil, err
	}
	var items []*Item
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil
}
