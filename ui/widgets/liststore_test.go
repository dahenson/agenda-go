package widgets

import(
	"testing"
	. "github.com/dahenson/agenda/types"
	. "github.com/dahenson/agenda/testutils"
	"github.com/conformal/gotk3/gtk"
)

func setup() *ListStore {
	b := initBuilder()
	return loadListStore(b)
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
			for f(iter); ls.IterNext(iter); f(iter) {}
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
