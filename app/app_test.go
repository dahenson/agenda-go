package app

import (
	"github.com/dahenson/agenda/testutils"
	. "github.com/dahenson/agenda/types"
	"testing"
)

type AppTest struct {
	ui  *testutils.FakeUi
	is  *testutils.FakeItemStore
	app *App
}

func InitFixture() *AppTest {
	fixture := new(AppTest)
	fixture.ui = testutils.NewFakeUi()
	fixture.is = testutils.NewFakeItemStore("default.txt")
	fixture.app = NewApp(fixture.is, fixture.ui)
	return fixture
}

// Given an empty itemstore
// When user enters some text
// And 'add' button is pressed
// Expect itemstore contains one item
// And that item has the same text as entered by the user
// And that the ui displays the item
// And that the ui entry text is cleared
func TestOnAddItem_ExpectOneItemAdded(t *testing.T) {
	fixture := InitFixture()
	expText := "Some text"

	// When user enters some text
	fixture.ui.SetEntryText(expText)
	// and 'add' button is pressed
	fixture.ui.PressAddButton()

	// Expect itemstore contains one item
	items, err := fixture.is.Items()
	if err != nil {
		t.Fatal("Untestutils.Expected err:", err)
	}
	if err := testutils.ExpectItemCount(1, len(items)); err != nil {
		t.Fatal(err)
	}

	// Expect that the correct item text was stored
	if err := testutils.ExpectText(expText, items[0].Text); err != nil {
		t.Fatal(err)
	}

	// Expect UI displays one item
	items = fixture.ui.Items()
	if err := testutils.ExpectItemCount(1, len(items)); err != nil {
		t.Fatal(err)
	}

	// Expect the UI displays the correct text
	if err := testutils.ExpectText(expText, items[0].Text); err != nil {
		t.Fatal(err)
	}

	// Expect that the entry text is cleared
	if err := testutils.ExpectText("", fixture.ui.GetEntryText()); err != nil {
		t.Fatal(err)
	}
}

// Given empty ItemStore
// When user enters some text
// And the 'add' button is pressed
// And the itemstore returns an error
// Then testutils.Expect the UI displays an error message
// And the UI doesn't display any items
// And testutils.Expect the text entry retains the entered text
func TestOnAddItem_WhenItemStoreReturnsError_ExpectUIDisplaysError(t *testing.T) {
	fixture := InitFixture()

	enteredText := "Some Text"

	fixture.ui.SetEntryText(enteredText)
	expErrText := "Failed to add item"
	fixture.is.ReturnErrorOnAddItem(expErrText)
	fixture.ui.PressAddButton()

	// Expect the UI displays the correct error text
	if err := testutils.ExpectText(expErrText, fixture.ui.CurrentErrorMessage()); err != nil {
		t.Fatal(err)
	}

	// Expect no items were added to the UI
	if err := testutils.ExpectItemCount(0, len(fixture.ui.Items())); err != nil {
		t.Fatal(err)
	}

	// Expect the text entry retains the entered text
	if err := testutils.ExpectText(enteredText, fixture.ui.GetEntryText()); err != nil {
		t.Fatal(err)
	}
}

// Given 3 items in itemstore
// When LoadItems() called
// Then expect that the UI displays all 3 items
func TestLoadItems_ExpectUIDisplaysAllItemsInItemStore(t *testing.T) {
	fixture := InitFixture()
	items := []*Item {
		NewItem("Item1"),
		NewItem("Item2"),
		NewItem("Item3"),
	}
	for _, item := range items {
		if err := fixture.is.AddItem(item); err != nil {
			t.Fatal("Unexpected err:", err)
		}
	}

	if err := fixture.app.LoadItems(); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	actItems := fixture.ui.Items()
	if err := testutils.ExpectItemCount(len(items), len(actItems)); err != nil {
		t.Fatal(err)
	}

	for i, item := range items {
		if err := testutils.ExpectText(item.Text, actItems[i].Text); err != nil {
			t.Fatal(err)
		}
	}
}
