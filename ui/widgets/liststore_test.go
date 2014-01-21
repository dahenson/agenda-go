package widgets

import (
	"fmt"
	"github.com/conformal/gotk3/gtk"
	. "github.com/dahenson/agenda/testutils"
	. "github.com/dahenson/agenda/types"
	"testing"
)

func setup() *ListStore {
	b := initBuilder()
	return loadListStore(b)
}

// Given empty ListStore
// When one item is added
// Then expect ListStore contains the item
func TestAddItem(t *testing.T) {
	// Given empty ListStore
	ctx := setup()

	// When one item is added
	item := NewItem("item")
	ctx.AddItem(item)

	// Then expect liststore contains the item
	items := ctx.Items()
	if err := ExpectCount(1, len(items)); err != nil {
		t.Fatal(err)
	}

	if err := Expect(item.Id(), items[0].Id()); err != nil {
		t.Fatal(err)
	}
}

// Given one completed item
// When a new item is added
// Then expect it is added above the completed item
func TestAddItem_ExpectItemAddedAboveCompletedItems(t *testing.T) {
	// Given one completed item
	ctx := setup()
	item := NewItem("item")
	item.SetComplete(true)
	ctx.AddItem(item)

	// When a new item is added
	newItem := NewItem("new item")
	ctx.AddItem(newItem)

	// Then expect it is added above the completed item
	items := ctx.Items()

	if err := ExpectCount(2, len(items)); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	if err := Expect(newItem.Id(), items[0].Id()); err != nil {
		t.Fatal("Expected new item added above completed item")
	}

	if err := Expect(item.Id(), items[1].Id()); err != nil {
		t.Fatal("Expected completed item is below new item")
	}
}

// Given one item added
// When one item removed
// Expect Items() returns empty list
func TestRemove_GivenOneItem_ExpectItemsReturnsEmptyList(t *testing.T) {
	ctx := setup()

	// Given one item
	item := NewItem("An item")
	ctx.AddItem(item)

	// When one item removed
	ctx.RemoveItem(item.Id())

	// Expect Items() returns empty list
	if err := ExpectCount(0, len(ctx.Items())); err != nil {
		t.Fatal(err)
	}
}

// Given two items added
// When first item removed
// Then expect 1 item remains in ListStore
// And expect removed item not in Items()
func TestRemove_GivenTwoItems_WhenFirstRemoved_ExpectItemsReturns1Item(t *testing.T) {
	ctx := setup()

	// Given two items
	first := NewItem("first")
	second := NewItem("second")
	ctx.AddItem(first)
	ctx.AddItem(second)

	// When first removed
	ctx.RemoveItem(first.Id())

	// Expect 1 item remains in ListStore
	if err := ExpectCount(1, ctx.Len()); err != nil {
		t.Fatal(err)
	}

	// And expect removed item not in Items()
	for _, item := range ctx.Items() {
		if err := ExpectItemNeq(item, first); err != nil {
			t.Fatal(err)
		}
	}
}

// Given two items added
// When second item removed
// Expect Items() returns 1-length list
// And expect removed item not in Items()
func TestRemove_GivenTwoItems_WhenSecondRemoved_ExpectItemsReturns1Item(t *testing.T) {
	getStr := func(ls *gtk.ListStore, iter *gtk.TreeIter) string {
		val, err := ls.GetValue(iter, COL_TEXT)
		if err != nil {
			t.Fatal("unexpected err:", err)
		}

		s, err := val.GetString()
		if err != nil {
			t.Fatal("unexpected err:", err)
		}

		return s
	}

	foreach := func(ls *gtk.ListStore, f func(iter *gtk.TreeIter)) {
		if iter, notEmpty := ls.GetIterFirst(); notEmpty {
			for f(iter); ls.IterNext(iter); f(iter) {
			}
		}
	}

	recordItemText := func(ls *gtk.ListStore) []string {
		state := []string{}
		foreach(ls, func(iter *gtk.TreeIter) {
			state = append(state, getStr(ls, iter))
		})
		return state
	}

	ctx := setup()

	// Given two items
	first := NewItem("first")
	second := NewItem("second")
	t.Log("Before any items added:", recordItemText(ctx.ListStore))
	ctx.AddItem(first)
	t.Log("After 1 item added:", recordItemText(ctx.ListStore))
	ctx.AddItem(second)
	beforeRemove := recordItemText(ctx.ListStore)
	t.Log("After 2 items added:", beforeRemove)

	// When second removed
	if err := ctx.RemoveItem(second.Id()); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	afterRemove := recordItemText(ctx.ListStore)

	// Expect 1 item remains in ListStore
	if err := ExpectCount(1, len(afterRemove)); err != nil {
		t.Log("Before remove:", beforeRemove)
		t.Log("After remove:", afterRemove)
		t.Fatal(err)
	}

	// And expect removed item not in Items()
	for _, item := range ctx.Items() {
		if err := ExpectItemNeq(item, second); err != nil {
			t.Fatal(err)
		}
	}
}

func TestSetItemComplete(t *testing.T) {
	ctx := setup()

	// Given 2 items
	first := NewItem("first")
	ctx.AddItem(first)
	ctx.AddItem(NewItem("second"))

	// When first item set complete
	if err := ctx.SetItemComplete(first.Id(), true); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	iter, err := ctx.findId(first.Id())
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}
	act := ctx.getComplete(iter)

	if exp := true; act != exp {
		t.Fatalf("Expected first item complete: %v; Got: %v", exp, act)
	}
}

// Given one active item
// When item completed
// Then expect no error returned
func TestSetItemComplete_ExpectNoErr(t *testing.T) {
	// Given one active item
	ctx := setup()
	item := NewItem("item")
	ctx.AddItem(item)

	// When item completed
	err := ctx.SetItemComplete(item.Id(), true)

	if err != nil {
		t.Fatal("Unexpected err:", err)
	}
}

// Given two items
// When the first item is marked complete
// Then expect the first item moves to the bottom
func TestSetItemComplete_ExpectItemMovesBelowActiveItems(t *testing.T) {
	// Given two items
	ctx := setup()
	first := NewItem("first")
	second := NewItem("second")
	ctx.AddItem(first)
	ctx.AddItem(second)

	// When the first item is marked complete
	ctx.SetItemComplete(first.Id(), true)

	// Then expect the first item moves to the bottom
	items := ctx.Items()
	if err := ExpectCount(2, len(items)); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	if err := Expect(second.Id(), items[0].Id()); err != nil {
		t.Fatal("Expected second item moved to top:", err)
	}

	if err := Expect(first.Id(), items[1].Id()); err != nil {
		t.Fatal("Expected first item moved to bottom:", err)
	}
}

// Given 2 active items
// And given 2 completed items
// When first active item is marked complete
// Then expect it is the first completed item in item list
func TestSetItemComplete_ExpectItemBecomesFirstCompletedItem(t *testing.T) {
	// Given 2 active items
	ctx := setup()
	items := []*Item{
		NewItem("0"),
		NewItem("1"),
		NewItem("2"),
		NewItem("3"),
	}
	// Given 2 completed items
	items[2].SetComplete(true)
	items[3].SetComplete(true)
	for _, item := range items {
		ctx.AddItem(item)
	}

	// When first active item is marked complete
	if err := ctx.SetItemComplete(items[0].Id(), true); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	// Then expect it is the first completed item in item list
	for _, item := range ctx.Items() {
		if item.Complete() {
			if err := Expect(items[0].Id(), item.Id()); err != nil {
				for _, item := range items {
					t.Log(item)
				}
				t.Fatal("Expected the newly-completed item moved to top of completed items list:")
			}
			return
		}
	}

	t.Fatal("The newly-completed item is missing from the Items() list")
}

// Given 2 active items
// And given 2 completed items
// When second active item is marked complete
// Then expect it is the first completed item in item list
func TestSetItemComplete_WhenSecondActiveItemCompleted_ExpectItemBecomesFirstCompletedItem(t *testing.T) {
	// Given 2 active items
	ctx := setup()
	items := []*Item{
		NewItem("0"),
		NewItem("1"),
		NewItem("2"),
		NewItem("3"),
	}
	// Given 2 completed items
	items[2].SetComplete(true)
	items[3].SetComplete(true)
	for _, item := range items {
		ctx.AddItem(item)
	}

	// When second active item is marked complete
	if err := ctx.SetItemComplete(items[1].Id(), true); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	// Then expect it is the first completed item in item list
	for _, item := range ctx.Items() {
		if item.Complete() {
			if err := Expect(items[1].Id(), item.Id()); err != nil {
				for _, item := range items {
					t.Log(item)
				}
				t.Fatal("Expected the newly-completed item moved to top of completed items list:")
			}
			return
		}
	}

	t.Fatal("The newly-completed item is missing from the Items() list")
}

// Given 2 completed items
// When last item is activated
// Then expect the items switch positions
func TestSetItemComplete_ExpectNewlyActivatedItemMovesAboveCompletedItems(t *testing.T) {
	// Given 2 completed items
	ctx := setup()
	for i := 0; i < 2; i++ {
		item := NewItem(fmt.Sprintf("%d", i))
		item.SetComplete(true)
		ctx.AddItem(item)
	}
	itemsBefore := ctx.Items()

	// When last item is activated
	if err := ctx.SetItemComplete(itemsBefore[len(ctx.Items())-1].Id(), false); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	// Then expect the items switch positions
	itemsAfter := ctx.Items()

	if err := ExpectCount(len(itemsBefore), len(itemsAfter)); err != nil {
		t.Fatal(err)
	}

	if err := Expect(itemsBefore[0].Id(), itemsAfter[1].Id()); err != nil {
		t.Fatal("Expected first item moved to the end:", err)
	}

	if err := Expect(itemsBefore[1].Id(), itemsAfter[0].Id()); err != nil {
		t.Fatal("Expected last item moved to the front:", err)
	}
}


// Given 1 active item
// And 2 completed items
// When last completed item is activated
// Then expect the newly-activated item moves below the originally active item
// And above the completed item
func TestSetItemComplete_ExpectNewlyActivatedItemMovesBelowActiveItems(t *testing.T) {
	// Given 1 active item and 2 completed items
	ctx := setup()
	items := make([]*Item, 3)
	for i:=0; i<len(items); i++ {
		items[i] = NewItem(fmt.Sprintf("%d", i))
		ctx.AddItem(items[i])
		if i>=1 {
			if err := ctx.SetItemComplete(items[i].Id(), true); err != nil {
				t.Fatal("Unexpected err:", err)
			}
		}
	}
	beforeItems := ctx.Items()
	lastItem := beforeItems[len(beforeItems)-1]

	// When last completed item is activated
	if err := ctx.SetItemComplete(lastItem.Id(), false); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	// Then expect the newly-activated item moves below the originally active item
	// And above the completed item
	exp := []*Item{beforeItems[0], lastItem, beforeItems[1]}

	for i, item := range ctx.Items() {
		if err := Expect(exp[i].Id(), item.Id()); err != nil {
			println("\nExpected:")
			for _, item := range exp {
				fmt.Println(item)
			}

			println("\nGot:")
			for _, item := range ctx.Items() {
				fmt.Println(item)
			}
			t.Fatal(err)
		}
	}
}
