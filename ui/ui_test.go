package ui

import (
	"testing"
	. "github.com/dahenson/agenda/types"
	"github.com/dahenson/agenda/testutils"
)

type Context struct {
	entry Entry
	liststore ListStore
	addButton *testutils.FakeAddButton
	toggleButton *testutils.FakeToggleButton
	ui Ui
}

func setup() *Context {
	ctx := new(Context)
	ctx.entry = testutils.NewFakeEntry()
	ctx.liststore = testutils.NewFakeListStore()
	ctx.addButton = testutils.NewFakeAddButton()
	ctx.toggleButton = testutils.NewFakeToggleButton()
	ctx.ui = NewUi(ctx.entry, ctx.liststore, ctx.addButton, ctx.toggleButton, func(){})
	return ctx
}

// Given no text in entry
// Expect GetEntryText() returns ""
func TestGetEntryText_GivenBlankEntry_ExpectEmptyStringReturned(t *testing.T) {
	ctx := setup()

	if err := testutils.Expect("", ctx.ui.GetEntryText()); err != nil {
		t.Fatal(err)
	}
}

// Given some text in entry
// Expect GetEntryText() returns entered text
func TestGetEntryText_GivenSomeTextInEntry_ExpectEnteredTextReturned(t *testing.T) {
	ctx := setup()

	// Given some text in entry
	text := "Some text"
	ctx.entry.SetText(text)

	// Expect entered text returned
	if err := testutils.Expect(text, ctx.ui.GetEntryText()); err != nil {
		t.Fatal(err)
	}
}

// Given no text in entry
// When ClearEntryText()
// Expect GetEntryText() returns ""
func TestClearEntryText_GivenBlankEntry_ExpectEmptyStringReturned(t *testing.T) {
	ctx := setup()

	// When ClearEntryText()
	ctx.ui.ClearEntryText()

	// Expect empty string returned
	if err := testutils.Expect("", ctx.ui.GetEntryText()); err != nil {
		t.Fatal(err)
	}
}

// Given some text in entry
// When ClearEntryText()
// Expect GetEntryText() empty string returned
func TestClearEntryText_GivenSomeTextInEntry_ExpectEmptyStringReturned(t *testing.T) {
	ctx := setup()

	// Given some text in entry
	text := "some text"
	ctx.entry.SetText(text)

	// When ClearEntryText()
	ctx.ui.ClearEntryText()

	// Expect empty string returned
	if err := testutils.Expect("", ctx.ui.GetEntryText()); err != nil {
		t.Fatal(err)
	}
}

// Given no items added
// Expect Items() returns an empty list
func TestItems_GivenNoItemsAdded_ExpectEmptyListReturned(t *testing.T) {
	ctx := setup()

	if err := testutils.ExpectCount(0, len(ctx.ui.Items())); err != nil {
		t.Fatal(err)
	}
}

// Given no items added
// When item added
// Expect Items() returns a 1-length list
// And expect added item is the first item in that list
func TestAddItem_GivenNoItemsAdded_ExpectItemAdded(t *testing.T) {
	ctx := setup()

	item := NewItem("Added item")

	// When item added
	ctx.ui.AddItem(item)

	// Expect Items() returns a 1-length list
	items := ctx.ui.Items()

	if err := testutils.ExpectCount(1, len(items)); err != nil {
		t.Fatal(err)
	}

	if err := testutils.Expect(item, items[0]); err != nil {
		t.Fatal(err)
	}
}

// Given add callback is set
// When add button clicked
// Expect callback called 1 time
func TestSetAddItemCallback_ExpectCallbackCalled(t *testing.T) {
	ctx := setup()

	callCount := 0

	// Given callback set
	ctx.ui.SetAddCallback(func() {
		callCount++
	})

	// When add button clicked
	ctx.addButton.Click()

	// Expect callback called one time
	if err := testutils.Expect(1, callCount); err != nil {
		t.Fatal(err)
	}
}

// Given toggle callback is set
// And one item added
// When added item is marked 'complete'
// Then expect callback called 1 time
func TestSetToggleItemCallback_GivenOneItemAdded_ExpectCallbackCalled(t *testing.T) {
	ctx := setup()

	callCount := 0

	// Given callback set
	ctx.ui.SetToggleCallback(func(_ string, _ bool) {
		callCount++
	})

	// Given one item added
	item := NewItem("An item")
	ctx.ui.AddItem(item)

	// When item toggled
	ctx.toggleButton.Click(item.Id, true)

	// When Expect callback called one time
	if err := testutils.Expect(1, callCount); err != nil {
		t.Fatal(err)
	}
}
