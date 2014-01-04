package widgets

import(
	"testing"
	. "github.com/dahenson/agenda/types"
	. "github.com/dahenson/agenda/testutils"
)

type Context struct {
	*ListStore
}

func setup() *Context {
	return &Context{ListStore: loadListStore()}
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
	ctx := setup()

	// Given two items
	first := NewItem("first")
	second := NewItem("second")
	ctx.AddItem(first)
	ctx.AddItem(second)

	// When second removed
	ctx.RemoveItem(second.Id())

	// Expect 1 item remains in ListStore
	if err := ExpectCount(1, ctx.Len()); err != nil {
		t.Fatal(err)
	}

	// And expect removed item not in Items()
	for _, item := range ctx.Items() {
		if err := ExpectItemNeq(item, second); err != nil {
			t.Fatal(err)
		}
	}
}
