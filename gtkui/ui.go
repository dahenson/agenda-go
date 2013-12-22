package gtkui

import (
	"fmt"
	. "github.com/dahenson/agenda/types"
	"github.com/dahenson/agenda/itemstore"
	"github.com/conformal/gotk3/gtk"
)

type Ui struct {
	liststore *gtk.ListStore
}

func NewUi(ls *gtk.ListStore) *Ui {
	return &Ui{liststore: ui}
}

func (ui *Ui) AddItem(item *Item) {
	var iter gtk.TreeIter
	ui.liststore.Append(&iter)
	ui.liststore.Set(&iter, []int{0, 1}, []interface{}{item.Complete, item.Text})
}
