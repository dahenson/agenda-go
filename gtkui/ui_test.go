package gtkui

import (
	"testing"
	"github.com/dahenson/agenda/testutils"
	. "github.com/dahenson/agenda/types"
	"github.com/conformal/gotk3/gtk"
)


type FakeWindow struct {}
func (FakeWindow) ShowAll() {}

type UiFixture struct {
	ls *testutils.FakeListStore
	ui *Ui
	win FakeWindow
}

func initFixture() *UiFixture {
	fixture := new(UiFixture)
	fixture.ls = testutils.NewFakeListStore()
	fixture.win = FakeWindow{}
	entry, _ := gtk.EntryNew()
	fixture.ui = NewUi(fixture.ls, entry, fixture.win)
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
