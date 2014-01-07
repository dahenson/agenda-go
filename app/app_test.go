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
	Entry        *FakeEntry
	ItemStore    *MockItemStore
	AddButton    *FakeAddButton
	ToggleButton *FakeToggleButton
	*App
}

func (ctx *Context) GivenNItemsAdded(n int) {
	for i:=0; i<n; i++ {
		ctx.Entry.SetText(fmt.Sprintf("item %d", i))
		ctx.AddButton.Click()
	}
}

func setup() *Context {
	ctx := new(Context)
	ctx.ItemStore = NewMockItemStore()
	ctx.Entry = NewFakeEntry()
	ctx.AddButton = NewFakeAddButton()
	ctx.ToggleButton = NewFakeToggleButton()
	ui := ui.NewUi(ctx.Entry, NewFakeListStore(), ctx.AddButton, ctx.ToggleButton, func() {})
	ctx.App = NewApp(ctx.ItemStore, ui, 10)
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
	if err := Expect(text, act[0].Text()); err != nil {
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
	if err := Expect(enteredText, ctx.ui.Items()[0].Text()); err != nil {
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
	ctx.GivenNItemsAdded(1)

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
	if err := Expect(enteredText, lastItemInLastCall.Text()); err != nil {
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
	ctx.GivenNItemsAdded(1)

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
	if err := Expect(enteredText, lastUiItem.Text()); err != nil {
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
	ctx.GivenNItemsAdded(1)
	item := ctx.ui.Items()[0]

	// When item toggled
	// Sending "false" for the state because it's the responsibility of the button's
	// callback to update the button state
	ctx.ToggleButton.Click(item.Id(), false)

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
	if err := Expect(true, lastCall[0].Complete()); err != nil {
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
	ctx.GivenNItemsAdded(1)
	item := ctx.ui.Items()[0]

	// And given ItemStore.Save() will return an error
	ctx.ItemStore.saveErr = fmt.Errorf("ItemStore: error saving")

	// When item toggled
	// Sending "false" for the state because it's the responsibility of the button's
	// callback to update the button state
	ctx.ToggleButton.Click(item.Id(), false)

	// Then expect the Ui shows the item is still incomplete
	if err := Expect(false, item.Complete()); err != nil {
		t.Fatal("Expecting item is complete:", err)
	}
}

// Given the max complete items threshold is 10 items
// And 11 items added
// And 10 items are complete
// When the last item is completed
// Then expect the first of the completed items is deleted from the UI
func TestToggleItem_GivenMaxCompletedItems_WhenNewItemAdded_ExpectOneItemDeleted(t *testing.T) {
	// Given the max complete items threshold is 10 items added (default)
	ctx := setup()

	// And 11 items added
	ctx.GivenNItemsAdded(11)
	items := ctx.ui.Items()

	incompleteItemIndex := 0
	completedItems := append(items[:incompleteItemIndex], items[incompleteItemIndex:]...)

	// And 10 items are complete
	for _, item := range completedItems {
		ctx.ui.SetItemComplete(item.Id(), true)
	}

	firstCompletedItem := findFirstCompletedItem(items)
	if firstCompletedItem == nil {
		t.Fatal("Unexpected err: couldn't find any completed items")
	}

	// When the last item is completed (passing `false` b/c the item is _currently_ not checked
	ctx.ToggleButton.Click(items[incompleteItemIndex].Id(), false)

	// Then expect the first of the completed items is deleted from the UI
	uiItems := ctx.ui.Items()
	if err := ExpectCount(10, ctx.ui.Len()); err != nil {
		t.Fatal(err)
	}

	for _, uiItem := range uiItems {
		if err := ExpectItemNeq(uiItem, firstCompletedItem); err != nil {
			t.Fatal(err)
		}
	}
}

// Given the max complete items threshold is 10 items
// And 11 items added
// And 5 items are complete (fewer than max)
// When an incomplete item is completed
// Then expect no items removed from UI
func TestToggleComplete_GivenFewCompleteItems_ExpectNoItemsRemovedFromUi(t *testing.T) {
	// Given the max complete items threshold is 10 items
	ctx := setup()

	// And 11 items added
	ctx.GivenNItemsAdded(11)

	// And 5 items are complete
	nItemsCompleted := 5
	items := ctx.ui.Items()
	for _, item := range items[:nItemsCompleted] {
		item.SetComplete(true)
	}

	// When an incomplete item is completed (false -> _current_ state--not state to be set to)
	ctx.ToggleButton.Click(items[nItemsCompleted].Id(), false)

	// Then expect no items removed from UI
	if err := ExpectCount(len(items), ctx.ui.Len()); err != nil {
		t.Fatal(err)
	}

	currentUiItems := ctx.ui.Items()
	for i:=0; i<len(items); i++ {
		if err := ExpectItemEq(currentUiItems[i], items[i]); err != nil {
			t.Fatal(err)
		}
	}
}

func expectItemsCorrectlyOrdered(items []*Item) error {
	foundCompletedItem := false
	var prev *Item
	for _, item := range items {
		if foundCompletedItem {
			// once we've found a completed item, make sure there are no subsequent
			// incomplete items
			if !item.Complete() {
				msg := "Found one or more completed items between incomplete items"
				return fmt.Errorf(msg)
			}
			// once we've found a completed item, make sure all subsequent items
			// are ordered descending by completion date
			if item.CompletedBefore(prev) {
				msg := "Completed items not ordered desc by last completion time"
				return fmt.Errorf(msg)
			}
		}
		if item.Complete() {
			foundCompletedItem = true
			prev = item
		}
	}
	return nil
}

// Given 2 incomplete items
// When first item marked complete
// Then expect items correctly ordered
func TestToggleComplete_GivenTwoItems_ExpectCompletedItemAtBottom(t *testing.T) {
	ctx := setup()

	// Given 2 incomplete items
	ctx.GivenNItemsAdded(2)

	items := ctx.ui.Items()
	first := items[0]

	// When first item marked complete
	ctx.App.OnToggleItem(first.Id(), false)

	// Then expect items correctly ordered
	if err := expectItemsCorrectlyOrdered(ctx.ui.Items()); err != nil {
		t.Fatal(err)
	}
}
