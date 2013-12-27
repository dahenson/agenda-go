package app

import (
	"fmt"
	. "github.com/dahenson/agenda/testutils"
	. "github.com/dahenson/agenda/types"
	"github.com/dahenson/agenda/ui"
	"testing"
)

type MockItemStore struct {
	saveCalls [][]*Item
	saveErr   error
	loadItems []*Item
	loadErr   error
}

func (is *MockItemStore) TimesSaveCalled() int { return len(is.saveCalls) }

func NewMockItemStore() *MockItemStore {
	return &MockItemStore{
		saveCalls: [][]*Item{},
		saveErr:   nil,
		loadItems: []*Item{},
		loadErr:   nil,
	}
}

func (is *MockItemStore) Save(items []*Item) error {
	is.saveCalls = append(is.saveCalls, items)
	return is.saveErr
}

func (is *MockItemStore) Load() ([]*Item, error) {
	return is.loadItems, is.loadErr
}

type Context struct {
	Entry     *FakeEntry
	ItemStore *MockItemStore
	AddButton *FakeAddButton
	ToggleButton *FakeToggleButton
	*App
}

func setup() *Context {
	ctx := new(Context)
	ctx.ItemStore = NewMockItemStore()
	ctx.Entry = NewFakeEntry()
	ctx.AddButton = NewFakeAddButton()
	ctx.ToggleButton = NewFakeToggleButton()
	ui := ui.NewUi(ctx.Entry, NewFakeListStore(), ctx.AddButton, ctx.ToggleButton, func(){})
	ctx.App = NewApp(ctx.ItemStore, ui)
	return ctx
}

// Given empty itemstore
// And some text in entry
// When Add button clicked
// Then expect `ItemStore.Save()` called one time
// And expect `ItemStore.Save()` receives the new item as its parameter
func TestAddItem_GivenEmptyItemStore_ExpectItemSavedAndAddedToUi(t *testing.T) {
	// Given empty itemstore
	ctx := setup()

	// And some text in entry
	text := "some text"
	ctx.Entry.SetText(text)

	// When Add button clicked
	ctx.AddButton.Click()

	// Then expect ItemStore.Save() called one time
	if err := ExpectCount(1, ctx.ItemStore.TimesSaveCalled()); err != nil {
		t.Fatal(err)
	}

	// And expect ItemStore.Save() receives the new item as its parameter
	act := ctx.ItemStore.saveCalls[0]
	if err := ExpectCount(1, len(act)); err != nil {
		t.Fatal(err)
	}
	if err := Expect(text, act[0].Text); err != nil {
		t.Fatal(err)
	}
}

// Given empty itemstore
// And some text in entry
// When Add button clicked
// Then expect UI contains one item
// And that item's text matches entered text
func TestAddItem_ExpectItemAddedToUi(t *testing.T) {
	// Given empty itemstore
	ctx := setup()

	// And some text in entry
	enteredText := "Some text"
	ctx.Entry.SetText(enteredText)

	// When OnAddItem() caled
	ctx.App.OnAddItem()

	// Then expect UI contains one item
	if err := ExpectCount(1, len(ctx.ui.Items())); err != nil {
		t.Fatal(err)
	}

	// And that item's text matches entered text
	if err := Expect(enteredText, ctx.ui.Items()[0].Text); err != nil {
		t.Fatal(err)
	}
}

// Given empty itemstore
// And some text in entry
// When Add button clicked
// Then expect entry cleared
func TestAddItem_ExpectEntryCleared(t *testing.T) {
	// Given empty itemstore
	ctx := setup()

	// And some text in entry
	ctx.Entry.SetText("Some text")

	// When Add button clicked
	ctx.AddButton.Click()

	// Then expect entry cleared
	if err := Expect("", ctx.Entry.GetText()); err != nil {
		t.Fatal("Expecting entry has been cleared:", err)
	}
}

// Given empty itemstore
// And some text in entry
// And ItemStore.Save() will return an error
// When Add button clicked
// Then expect UI contains no items
func TestAddItem_WhenSaveReturnsErr_ExpectItemNotAdded(t *testing.T) {
	// Given empty itemstore
	ctx := setup()
	// And Some text in entry
	ctx.Entry.SetText("Some text")
	// And ItemStore.Save() will return an error
	ctx.ItemStore.saveErr = fmt.Errorf("An error")

	// When Add button clicked
	ctx.AddButton.Click()

	// Expect UI contains no items
	if err := ExpectCount(0, len(ctx.ui.Items())); err != nil {
		t.Fatal("Expecting UI contains no items:\n", err)
	}
}

// Given empty itemstore
// And some text in entry
// And ItemStore.Save() will return an error
// When Add button clicked
// The expect entry text unchanged
func TestAddItem_WhenSaveReturnsErr_ExpectEntryTextUnchanged(t *testing.T) {
	// Given empty itemstore
	ctx := setup()

	// And some text in entry
	enteredText := "Some Text"
	ctx.Entry.SetText(enteredText)

	// And ItemStore.Save() will return an error
	ctx.ItemStore.saveErr = fmt.Errorf("Some err")

	// When Add button clicked
	ctx.AddButton.Click()

	// The expect entry text unchanged
	if err := Expect(enteredText, ctx.Entry.GetText()); err != nil {
		t.Fatal("Expecting entry's text unchanged:", err)
	}
}

// Given one item added
// And some text in entry
// When add button clicked
// Then expect ItemStore.Save() called twice
// And expect 2 items in last call to ItemStore.Save()
// And expect the last item in last call to ItemStore.Save() has text matching the entered text
func TestAddItem_GivenOneItemAdded_ExpectNewItemSaved(t *testing.T) {
	ctx := setup()

	// Given one item added
	ctx.Entry.SetText("First item")
	ctx.AddButton.Click()

	// And some text in entry
	enteredText := "Second item"
	ctx.Entry.SetText(enteredText)

	// When add button clicked
	ctx.AddButton.Click()

	// Then expect ItemStore.Save() called twice
	if err := ExpectCount(2, len(ctx.ItemStore.saveCalls)); err != nil {
		t.Fatal(err)
	}

	// And expect 2 items in last call to ItemStore.Save()
	lastCall := ctx.ItemStore.saveCalls[1]
	if err := ExpectCount(2, len(lastCall)); err != nil {
		t.Fatal(err)
	}

	// And expect the last item in last call to ItemStore.Save() has text matching the entered text
	lastItemInLastCall := lastCall[1]
	if err := Expect(enteredText, lastItemInLastCall.Text); err != nil {
		t.Fatal(err)
	}
}

// Given one item added
// And some text in entry
// When add button clicked
// Then expect 2 items in UI
// And expect the last item in UI has text matching the entered text
func TestAddItem_GivenOneItemAdded_ExpectNewItemAddedToUi(t *testing.T) {
	ctx := setup()

	// Given one item added
	ctx.Entry.SetText("First item")
	ctx.AddButton.Click()

	// And some text in entry
	enteredText := "Second item"
	ctx.Entry.SetText(enteredText)

	// When add button clicked
	ctx.AddButton.Click()

	// Then expect 2 items in UI
	items := ctx.ui.Items()
	if err := ExpectCount(2, len(items)); err != nil {
		t.Fatal(err)
	}

	// And expect the last item in UI has text matching the entered text
	lastUiItem := items[1]
	if err := Expect(enteredText, lastUiItem.Text); err != nil {
		t.Fatal(err)
	}
}

// Given empty itemstore
// When application loaded
// Expect UI displays no items
func TestLoad_GivenEmptyItemStore_ExpectNoItems(t *testing.T) {
	// Given empty itemstore
	ctx := setup()

	// When application loaded
	ctx.Load()

	// expect the UI displays no items
	if err := ExpectCount(0, len(ctx.ui.Items())); err != nil {
		t.Fatal()
	}
}

// Given one item added
// When item toggled
// Then expect ItemStore.Save() called twice
// And expect the last call to contain all the items in the UI
func TestToggleItem(t *testing.T) {
	ctx := setup()

	// Given one item added
	ctx.Entry.SetText("An item")
	ctx.AddButton.Click()
	item := ctx.ui.Items()[0]

	// When item toggled
	// Sending "false" for the state because it's the responsibility of the button's
	// callback to update the button state
	ctx.ToggleButton.Click(item.Id, false)

	// Then expect `ItemStore.Save()` called twice
	calls := ctx.ItemStore.saveCalls
	if err := ExpectCount(2, len(calls)); err != nil {
		t.Fatal(err)
	}

	// And expect the last call to contain a 1-length list
	lastCall := calls[1]
	if err := ExpectCount(1, len(lastCall)); err != nil {
		t.Fatal(err)
	}

	// And expect the item in that list to have a `Complete` value of `true`
	if err := Expect(true, lastCall[0].Complete); err != nil {
		t.Fatal("Expecting complete value:", err)
	}
}

// Given one item added
// And given ItemStore.Save() will return an error
// When item toggled
// Then expect the Ui shows the item is still incomplete
func TestToggleItem_WhenSaveReturnsErr_ExpectNoChangetoUi(t *testing.T) {
	ctx := setup()

	// Given one item added
	ctx.Entry.SetText("An item")
	ctx.AddButton.Click()
	item := ctx.ui.Items()[0]

	// And given ItemStore.Save() will return an error
	ctx.ItemStore.saveErr = fmt.Errorf("ItemStore: error saving")

	// When item toggled
	// Sending "false" for the state because it's the responsibility of the button's
	// callback to update the button state
	ctx.ToggleButton.Click(item.Id, false)

	// Then expect the Ui shows the item is still incomplete
	if err := Expect(false, item.Complete); err != nil {
		t.Fatal("Expecting item is complete:", err)
	}
}
