package widgets

import (
	"github.com/conformal/gotk3/gtk"
	. "github.com/dahenson/agenda/ui/uicallbacks"
)

type Entry struct {
	*gtk.Entry
}

func NewEntry(e *gtk.Entry) *Entry {
	return &Entry{e}
}

func (e *Entry) GetText() string {
	text, _ := e.Entry.GetText()
	return text
}

func (e *Entry) SetCallback(callback AddCallback) {
	e.Connect("activate", callback)
	e.Connect("icon-release", callback)
}
