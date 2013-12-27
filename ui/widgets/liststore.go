package widgets

import (
	. "github.com/dahenson/agenda/types"
	"github.com/weberc2/gotk3/gtk"
	"log"
)

type row struct {
	id   string
	iter *gtk.TreeIter
}

type style int

const (
	normal style = 0
	oblique style = 1
	italic style = 2
	COL_STYLE    = 2
	COL_COMPLETE = 1
	COL_TEXT     = 0
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
	cols := []int{COL_COMPLETE, COL_TEXT, COL_STYLE}
	vals := []interface{}{item.Complete, item.Text, determineStyle(item.Complete)}
	if err := ls.Set(&iter, cols, vals); err != nil {
		log.Fatal(err)
	}
	r.id = item.Id
	r.iter = &iter
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
	return &Item{
		Id:       r.id,
		Text:     ls.getText(r.iter),
		Complete: ls.getComplete(r.iter),
	}
}

func (ls *ListStore) setComplete(iter *gtk.TreeIter, complete bool) error {
	cols := []int{COL_COMPLETE, COL_STYLE}
	vals := []interface{}{complete, determineStyle(complete)}
	return ls.Set(iter, cols, vals)
}

func (ls *ListStore) Items() []*Item {
	items := make([]*Item, len(ls.rows))
	for i, r := range ls.rows {
		items[i] = ls.get(r)
	}
	return items
}

func (ls *ListStore) SetItemComplete(id string, complete bool) error {
	for _, r := range ls.rows {
		if r.id == id {
			return ls.setComplete(r.iter, complete)
		}
	}
	return nil
}
