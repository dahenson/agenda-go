package widgets

import (
	. "github.com/dahenson/agenda/types"
	"github.com/dahenson/agenda/ui"
	"github.com/weberc2/gotk3/gtk"
	"log"
)

const (
	COL_COMPLETE = 1
	COL_TEXT     = 0
)

type gtkPath string

type ListStore struct {
	ls       *gtk.ListStore
	renderer *gtk.CellRendererToggle
	idstore  map[gtkPath]string
	pathstore map[string]gtkPath
}

func NewListStore(ls *gtk.ListStore, renderer *gtk.CellRendererToggle) *ListStore {
	return &ListStore{
		ls: ls,
		idstore: map[gtkPath]string{},
		pathstore: map[string]gtkPath{},
		renderer: renderer,
	}
}

func (ls *ListStore) AddItem(item *Item) {
	iter := new(gtk.TreeIter)
	ls.ls.Append(iter)
	path, err := ls.ls.GetPath(iter)
	if err != nil {
		log.Fatal(err)
	}
	ls.idstore[gtkPath(path.String())] = item.Id
	ls.pathstore[item.Id] = gtkPath(path.String())
	cols := []int{COL_COMPLETE, COL_TEXT}
	vals := []interface{}{item.Complete, item.Text}
	if err := ls.ls.Set(iter, cols, vals); err != nil {
		log.Fatal(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (ls *ListStore) SetToggled(id string, toggled bool) {
	path, found := ls.pathstore[id]
	if !found {
		log.Fatal("Couldn't find id in pathstore")
	}

	iter, err := ls.ls.GetIterFromString(string(path))
	check(err)

	if err := ls.ls.Set(iter, []int{COL_COMPLETE}, []interface{}{toggled}); err != nil {
		log.Fatal(err)
	}
}

func (ls *ListStore) SetOnToggled(callback ui.ToggleItemCallback) {
	ls.renderer.Connect("toggled", func(_ interface{}, pathStr string) {
		iter, err := ls.ls.GetIterFromString(pathStr)
		check(err)

		glibVal, err := ls.ls.GetValue(iter, COL_COMPLETE)
		check(err)

		goIface, err := glibVal.GoValue()
		check(err)

		toggled, ok := goIface.(bool)
		if !ok {
			log.Fatal("Failed to convert liststore value to bool")
		}

		id, found := ls.idstore[gtkPath(pathStr)]
		if !found {
			log.Fatal("Couldn't find path in idstore")
		}

		callback(id, toggled)
	})
}
