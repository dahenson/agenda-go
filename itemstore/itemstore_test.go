package itemstore

import (
	"fmt"
	"github.com/dahenson/agenda/fs"
	"github.com/dahenson/agenda/testutils"
	. "github.com/dahenson/agenda/types"
	"testing"
)

type Context struct {
	itemstore ItemStore
}

func setup() *Context {
	return &Context{itemstore: NewItemStoreWithFileSystem("default.txt", fs.NewFakeFS())}
}

func eq(i1, i2 *Item) bool {
	return i1.Id() == i2.Id() && i1.Complete() == i2.Complete() && i1.Text() == i2.Text()
}

func expectItemListsEq(exp, act []*Item) error {
	// compare lengths
	if err := testutils.ExpectCount(len(exp), len(act)); err != nil {
		return err
	}
	// compare items in each list
	for i := 0; i < len(exp); i++ {
		if !eq(exp[i], act[i]) {
			return fmt.Errorf("Expected %v; Got %v", exp, act)
		}
	}
	return nil
}

// When `ItemStore.Save()` called
// Then expect no error returned
func TestSave_WhenSaved_ExpectNoErrs(t *testing.T) {
	ctx := setup()
	items := []*Item{}

	// When Save() called
	if err := ctx.itemstore.Save(items); err != nil {
		t.Fatal(err)
	}
}

// When saved with empty item list
// And `Load()` called
// Then expect empty item list returned
func TestSave_WhenEmptyListSaved_ExpectLoadReturnsEmptyList(t *testing.T) {
	ctx := setup()
	items := []*Item{}

	// When Save() called
	if err := ctx.itemstore.Save(items); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	// When Load() called
	loaded, err := ctx.itemstore.Load()
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}

	// Expect empty list loaded
	if err := expectItemListsEq(items, loaded); err != nil {
		t.Fatal(err)
	}
}

// When multiple items saved
// And `Load()` called
// Then expect loaded items match saved items
func TestSave_WhenItemsSaved_ExpectLoadedItemsMatchSaved(t *testing.T) {
	ctx := setup()
	items := []*Item{
		NewItem("First"),
		NewItem("Second"),
		NewItem("Third"),
	}

	// When Save() called
	if err := ctx.itemstore.Save(items); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	// When Load() called
	loaded, err := ctx.itemstore.Load()
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}

	// Expect loaded items match saved items
	if err := expectItemListsEq(items, loaded); err != nil {
		t.Fatal(err)
	}
}
