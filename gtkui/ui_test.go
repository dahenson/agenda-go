package gtkui

import (
	"testing"
	"github.com/dahenson/agenda/testutils"
	. "github.com/dahenson/agenda/types"
	"github.com/conformal/gotk3/glib"
	"log"
)

type FakeEntry struct {
	text string
	onActivate func()
}

func NewFakeEntry() *FakeEntry {
	return &FakeEntry{text: "", onActivate: func() {}}
}

func (e *FakeEntry) SetText(text string) { e.text = text }
func (e *FakeEntry) GetText() (string, error) { return e.text, nil }
func (e *FakeEntry) Activate() bool { e.onActivate(); return true }
func (e *FakeEntry) Connect(signal string, f interface{}, data...interface{}) (glib.SignalHandle, error) {
	callback, ok := f.(func())
	if !ok {
		log.Fatal("Failed to convert `f` to `func()`")
	}
	if signal == "activate" {
		e.onActivate = callback
	}
	return 0, nil
}

type FakeWindow struct {}
func (FakeWindow) ShowAll() {}

type UiFixture struct {
	ls *testutils.FakeListStore
	ui *Ui
}

func initFixture() *UiFixture {
	fixture := new(UiFixture)
	fixture.ls = testutils.NewFakeListStore()
	fixture.ui = NewUi(fixture.ls, NewFakeEntry())
	fixture.win = FakeWindow{}
	return fixture
}

func TestAddItem_ExpectItemAdded(t *testing.T) {
	fixture := initFixture()
	exp := NewItem("Some Text")
	fixture.ui.AddItem(exp)

	// expect 1 item added
	items := fixture.ls.Items()
	if err := testutils.ExpectItemCount(1, len(items)); err != nil {
		t.Fatal(err)
	}

	// expect correct item added
	if err := testutils.ExpectText(exp.Text, items[0].Text); err != nil {
		t.Fatal(err)
	}
}

// Given callback set
// When entry activated
// Then expect the callback gets called exactly one time
func TestSetAddItemCallback_ExpectCallbackCalledOnAddButtonPressed(t *testing.T) {
	timesCalled := 0
	callback := func() {
		timesCalled += 1
	}

	fixture := initFixture()
	fixture.ui.SetAddItemCallback(callback)

	// on entry activated
	fixture.ui.entry.Activate()

	if timesCalled != 1 {
		t.Fatalf("Expected callback called %d times; it was actually called %d times", 1, timesCalled)
	}
}

// Given some text in item entry
// When ClearEntryText called
// Then expect entry text cleared
func TestClearEntryText_ExpectEntryTextCleared(t *testing.T) {
	fixture := initFixture()
	fixture.ui.entry.SetText("Some text")

	fixture.ui.ClearEntryText()

	if err := testutils.ExpectText("", fixture.ui.GetEntryText()); err != nil {
		t.Fatal(err)
	}
}
