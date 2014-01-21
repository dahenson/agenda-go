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
	var newIter gtk.TreeIter

	found := false
	if cursor, listIsntEmpty := ls.GetIterFirst(); listIsntEmpty {
		for {
			if ls.getComplete(cursor) {
				ls.InsertBefore(&newIter, cursor)
				found = true
				break
			}

			if !ls.IterNext(cursor) {
				break
			}
		}
	}

	if !found {
		ls.Append(&newIter)
	}

	lastCompleted := item.LastTimeCompleted().Format(time.ANSIC)
	cols := []int{COL_COMPLETE, COL_TEXT, COL_STYLE, COL_SENSITIVE, COL_ID, COL_LAST_COMPLETED}
	vals := []interface{}{item.Complete(), item.Text(), determineStyle(item.Complete()), !item.Complete(), item.Id(), lastCompleted}
	if err := ls.Set(&newIter, cols, vals); err != nil {
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

func (ls *ListStore) foreach(f func(*gtk.TreeIter), _break *bool) {
	// if the list is empty, do nothing
	if iter, listIsntEmpty := ls.ListStore.GetIterFirst(); listIsntEmpty {
		for f(iter); ls.ListStore.IterNext(iter) && !(*_break); f(iter) {}
	}
}

func (ls *ListStore) Items() []*Item {
	items := []*Item{}
	_break := false
	ls.foreach(func(iter *gtk.TreeIter) {
		items = append(items, ls.get(iter))
	}, &_break)
	return items
}

func (ls *ListStore) Len() int {
	count := 0
	_break := false
	ls.foreach(func(_ *gtk.TreeIter) {
		count++
	}, &_break)
	return count
}

var IdNotFoundErr = fmt.Errorf("Couldn't find specified id")

func (ls *ListStore) findId(id string) (*gtk.TreeIter, error) {
	var i gtk.TreeIter
	found := false
	_break := false
	ls.foreach(func(iter *gtk.TreeIter) {
		if ls.getId(iter) == id {
			i = *iter
			found = true
			return
		}
	}, &_break)
	if !found{
		return nil, IdNotFoundErr
	}
	return &i, nil
}

func (ls *ListStore) SetItemComplete(id string, complete bool) error {
	var firstCompletedItem *gtk.TreeIter = nil
	var specifiedIter *gtk.TreeIter = nil
	var lastIter *gtk.TreeIter = nil

	_break := false
	var err error = nil
	ls.foreach(func(iterPtr *gtk.TreeIter) {
		iter := *iterPtr
		if firstCompletedItem == nil && ls.getComplete(iterPtr) {
			firstCompletedItem = &iter
			// if both are found, break
			if specifiedIter != nil {
				_break = true
			}
		}

		if ls.getId(iterPtr) == id {
			specifiedIter = &iter
			if err = ls.setComplete(iterPtr, complete); err != nil {
				_break = true
			}
			// if both are found, break
			if firstCompletedItem != nil {
				_break = true
			}
		}

		lastIter = &iter
	}, &_break)

	if err != nil {
		return err
	}

	if specifiedIter == nil {
		return IdNotFoundErr
	}

	// if the newly-completed item is the first completed item
	if firstCompletedItem == nil {
		ls.MoveAfter(specifiedIter, lastIter)
		return nil
	}

	// move newly-completed item before the first completed item
	ls.MoveBefore(specifiedIter, firstCompletedItem)
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
