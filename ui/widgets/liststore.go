package widgets

import (
	"fmt"
	. "github.com/dahenson/agenda/types"
	"github.com/conformal/gotk3/gtk"
	"log"
	"time"
)

type style int

const (
	normal             style = 0
	oblique            style = 1
	italic             style = 2
	COL_LAST_COMPLETED       = 5
	COL_ID                   = 4
	COL_SENSITIVE            = 3
	COL_STYLE                = 2
	COL_COMPLETE             = 1
	COL_TEXT                 = 0
)

type ListStore struct {
	*gtk.ListStore
}

func NewListStore(ls *gtk.ListStore) *ListStore {
	return &ListStore{ListStore: ls}
}

func determineStyle(complete bool) style {
	if complete {
		return italic
	}
	return normal
}

func (ls *ListStore) AddItem(item *Item) {
	var iter gtk.TreeIter
	ls.Append(&iter)
	lastCompleted := item.LastTimeCompleted().Format(time.ANSIC)
	cols := []int{COL_COMPLETE, COL_TEXT, COL_STYLE, COL_SENSITIVE, COL_ID, COL_LAST_COMPLETED}
	vals := []interface{}{item.Complete(), item.Text(), determineStyle(item.Complete()), !item.Complete(), item.Id(), lastCompleted}
	if err := ls.Set(&iter, cols, vals); err != nil {
		log.Fatal("Failed to add item:", err)
	}
}

func (ls *ListStore) getVal(iter *gtk.TreeIter, col int) interface{} {
	val, err := ls.GetValue(iter, col)
	if err != nil {
		log.Fatal("Failed to get value from TreeView:", err)
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

func (ls *ListStore) getId(iter *gtk.TreeIter) string {
	goval := ls.getVal(iter, COL_ID)
	id, ok := goval.(string)
	if !ok {
		log.Fatal("Failed to type-cast interface{} to string")
	}
	return id
}

func (ls *ListStore) getLastComplete(iter *gtk.TreeIter) string {
	goval := ls.getVal(iter, COL_LAST_COMPLETED)
	last, ok := goval.(string)
	if !ok {
		log.Fatal("Failed to type-cast interface{} to string")
	}
	return last
}

func (ls *ListStore) setComplete(iter *gtk.TreeIter, complete bool) error {
	cols := []int{COL_COMPLETE, COL_STYLE, COL_SENSITIVE}
	vals := []interface{}{complete, determineStyle(complete), !complete}
	return ls.Set(iter, cols, vals)
}

// this will crash if you pass an invalid iter
func (ls *ListStore) get(iter *gtk.TreeIter) *Item {
	id := ls.getId(iter)
	text := ls.getText(iter)
	complete := ls.getComplete(iter)
	last := ls.getLastComplete(iter)
	t, err := time.Parse(time.ANSIC, last)
	if err != nil {
		log.Println("Failed to parse time:", err)
	}
	return NewItemFromData(id, text, complete, t)
}

func (ls *ListStore) foreach(f func(*gtk.TreeIter)) {
	// if the list is empty, do nothing
	if iter, listIsntEmpty := ls.ListStore.GetIterFirst(); listIsntEmpty {
		for f(iter); ls.ListStore.IterNext(iter); f(iter) {}
	}
}

func (ls *ListStore) Items() []*Item {
	items := []*Item{}
	ls.foreach(func(iter *gtk.TreeIter) {
		items = append(items, ls.get(iter))
	})
	return items
}

func (ls *ListStore) Len() int {
	count := 0
	ls.foreach(func(_ *gtk.TreeIter) {
		count++
	})
	return count
}

func (ls *ListStore) findId(id string) (*gtk.TreeIter, error) {
	var i gtk.TreeIter
	found := false
	ls.foreach(func(iter *gtk.TreeIter) {
		if ls.getId(iter) == id {
			i = *iter
			found = true
			return
		}
	})
	if !found{
		return nil, fmt.Errorf("Couldn't find id: %s", id)
	}
	return &i, nil
}

func (ls *ListStore) SetItemComplete(id string, complete bool) error {
	iter, err := ls.findId(id)
	if err != nil {
		return err
	}
	if err := ls.setComplete(iter, complete); err != nil {
		log.Println("Failed to set COL_COMPLETE:", err)
	}
	return nil
}

func (ls *ListStore) RemoveItem(id string) error {
	iter, err := ls.findId(id)
	if err != nil {
		return err
	}
	ls.Remove(iter)
	return nil
}
