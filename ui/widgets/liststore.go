package widgets

import (
	"fmt"
	. "github.com/dahenson/agenda/types"
	"github.com/weberc2/gotk3/gtk"
	"log"
	"time"
)

type row struct {
	id                string
	iter              *gtk.TreeIter
	lastTimeCompleted time.Time
}

type style int

const (
	normal        style = 0
	oblique       style = 1
	italic        style = 2
	COL_SENSITIVE       = 3
	COL_STYLE           = 2
	COL_COMPLETE        = 1
	COL_TEXT            = 0
)

type ListStore struct {
	*gtk.ListStore
	rows []*row
}

func NewListStore(ls *gtk.ListStore) *ListStore {
	return &ListStore{ListStore: ls, rows: []*row{}}
}

func determineStyle(complete bool) style {
	if complete {
		return italic
	}
	return normal
}

func (ls *ListStore) AddItem(item *Item) {
	r := new(row)
	var iter gtk.TreeIter
	ls.Append(&iter)
	cols := []int{COL_COMPLETE, COL_TEXT, COL_STYLE, COL_SENSITIVE}
	vals := []interface{}{item.Complete(), item.Text(), determineStyle(item.Complete()), !item.Complete()}
	if err := ls.Set(&iter, cols, vals); err != nil {
		log.Fatal(err)
	}
	r.id = item.Id()
	r.iter = &iter
	r.lastTimeCompleted = item.LastTimeCompleted()
	ls.rows = append(ls.rows, r)
}

func (ls *ListStore) getVal(iter *gtk.TreeIter, col int) interface{} {
	val, err := ls.GetValue(iter, col)
	if err != nil {
		log.Fatal("Failed to get COL_COMPLETE value from TreeView:", err)
	}
	goval, err := val.GoValue()
	if err != nil {
		log.Fatal("Failed to convert GLib.Value to Go type:", err)
	}
	return goval
}

func (ls *ListStore) getComplete(iter *gtk.TreeIter) bool {
	goval := ls.getVal(iter, COL_COMPLETE)
	complete, ok := goval.(bool)
	if !ok {
		log.Fatal("Failed to type-cast interface{} to bool")
	}
	return complete
}

func (ls *ListStore) getText(iter *gtk.TreeIter) string {
	goval := ls.getVal(iter, COL_TEXT)
	text, ok := goval.(string)
	if !ok {
		log.Fatal("Failed to type-cast interface{} to string")
	}
	return text
}

func (ls *ListStore) get(r *row) *Item {
	return NewItemFromData(r.id, ls.getText(r.iter), ls.getComplete(r.iter), r.lastTimeCompleted)
}

func (ls *ListStore) setComplete(iter *gtk.TreeIter, complete bool) error {
	cols := []int{COL_COMPLETE, COL_STYLE, COL_SENSITIVE}
	vals := []interface{}{complete, determineStyle(complete), !complete}
	return ls.Set(iter, cols, vals)
}

func (ls *ListStore) Items() []*Item {
	items := make([]*Item, len(ls.rows))
	for i, r := range ls.rows {
		items[i] = ls.get(r)
	}
	return items
}

func (ls *ListStore) Len() int {
	return len(ls.rows)
}

func (ls *ListStore) getRowIndex(id string) (int, error) {
	for i, r := range ls.rows {
		if r.id == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Couldn't find id: %s", id)
}

func (ls *ListStore) getRow(id string) (*row, error) {
	i, err := ls.getRowIndex(id)
	if err != nil {
		return nil, err
	}
	return ls.rows[i], nil
}

func (ls *ListStore) SetItemComplete(id string, complete bool) error {
	r, err := ls.getRow(id)
	if err != nil {
		return err
	}
	if err := ls.setComplete(r.iter, complete); err != nil {
		return err
	}
	if complete {
		r.lastTimeCompleted = time.Now()
	}
	return nil
}

func (ls *ListStore) RemoveItem(id string) error {
	i, err := ls.getRowIndex(id)
	if err != nil {
		return err
	}
	r := ls.rows[i]
	ls.Remove(r.iter)
	// Remove the row from rows
	ls.rows = append(ls.rows[:i], ls.rows[i+1:]...)
	return nil
}
