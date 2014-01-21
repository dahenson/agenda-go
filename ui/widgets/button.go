package widgets

import (
	"github.com/conformal/gotk3/gtk"
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

		id := btn.ListStore.getId(iter)
		complete := btn.ListStore.getComplete(iter)
		callback(id, complete)
		return
	})
}
