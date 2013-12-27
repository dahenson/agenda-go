package widgets

import (
	"github.com/weberc2/gotk3/gtk"
	. "github.com/dahenson/agenda/ui/uicallbacks"
	"log"
)

type ToggleButton struct {
	*gtk.CellRendererToggle
	*ListStore
}

func NewToggleButton(btn *gtk.CellRendererToggle, ls *ListStore) *ToggleButton {
	return &ToggleButton{CellRendererToggle: btn, ListStore: ls}
}

func (btn *ToggleButton) SetCallback(callback ToggleCallback) {
	btn.CellRendererToggle.Connect("toggled", func(_ interface{}, path string) {
		iter, err := btn.GetIterFromString(path)
		if err != nil {
			log.Fatalf("Failed to get iter from path: '%s': %v", path, err)
		}

		for _, r := range btn.ListStore.rows {
			if *r.iter == *iter {
				callback(r.id, btn.ListStore.getComplete(iter))
				return
			}
		}
		log.Fatal("Couldn't find an item matching path:", path)
	})
}
