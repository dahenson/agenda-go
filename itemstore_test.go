package main

import (
	"testing"
)

func initTestItemStore() *FSItemStore {
	is := NewFSItemStore("default.txt")
	is.fs = newFakeFS()
	return is
}

func eq(i1, i2 Item) bool {
	return i1.Text() == i2.Text()
}

func TestAddItem_WhenEmpty_ExpectItemInItems(t *testing.T) {
	is := initTestItemStore()
	newItem := NewItem("An item")
	if err := is.AddItem(newItem); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	items, err := is.Items()
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}
	if len(items) != 1 {
		for _, item := range items {
			println(item.Text())
		}
		t.Fatalf("Expected 1 item, but found %d", len(items))
	}

	if exp, act := newItem, items[0]; !eq(act, exp) {
		t.Errorf("Expected item: %s; actual item: %s", exp, act)
	}
}

func TestAddItem_WhenTwoItemsAdded_ExpectBothItemsInItems(t *testing.T) {
	is := initTestItemStore()
	addedItems := []Item{NewItem("First item"), NewItem("Second item")}

	for _, item := range addedItems {
		if err := is.AddItem(item); err != nil {
			t.Fatal("Unexpected err:", err)
		}
	}

	storedItems, err := is.Items()
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}

	if len(storedItems) != len(addedItems) {
		println("Added items:")
		for _, item := range addedItems {
			println(item.Text())
		}

		println("\nStored items:")
		for _, item := range storedItems {
			println(item.Text())
		}
		t.Fatalf("Expected %d items, but found %d", len(addedItems), len(storedItems))
	}

	for i:=0; i<len(addedItems); i++ {
		if act, exp := storedItems[i], addedItems[i]; !eq(act, exp) {
			t.Fatalf("Expected: %s; Got: %s", exp, act)
		}
	}
}
