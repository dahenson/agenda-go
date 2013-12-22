package widgets

import (
	"github.com/conformal/gotk3/gtk"
	. "github.com/dahenson/agenda/types"
	"log"
)

const (
	COL_COMPLETE = 1
	COL_TEXT = 0
)

type ListStore struct {
	ls *gtk.ListStore
	itemmap map[string]*gtk.TreeIter
}

func NewListStore(ls *gtk.ListStore) *ListStore {
	return &ListStore{ls: ls, itemmap: make(map[string]*gtk.TreeIter)}
}

func (ls *ListStore) AddItem(item *Item) {
	iter := new(gtk.TreeIter)
	ls.itemmap[item.Text] = iter // TODO use item.Id when implemented
	ls.ls.Append(iter)
	cols := []int{COL_COMPLETE, COL_TEXT}
	vals := []interface{}{item.Complete, item.Text}
	if err := ls.ls.Set(iter, cols, vals); err != nil {
		log.Fatal(err)
	}
}
