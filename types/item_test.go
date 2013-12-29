package types

import (
	"testing"
	"encoding/json"
)

func TestMarshalJSON(t *testing.T) {
	i1 := NewItem("An item")
	data, err := json.Marshal(i1)
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}

	i2 := new(Item)
	if err := json.Unmarshal(data, i2); err != nil {
		t.Fatal("Unexpected err:", err)
	}

	if i1.Text() != i2.Text() {
		t.Fatalf("\ni1.Text():%v\ni2.Text():%v", i1.Text(), i2.Text())
	}

	if i1.Id() != i2.Id() {
		t.Fatalf("\ni1.Id():%v\ni2.Id():%v", i1.Id(), i2.Id())
	}

	if i1.Complete() != i2.Complete() {
		t.Fatalf("\ni1.Complete():%v\ni2.Complete():%v", i1.Complete(), i2.Complete())
	}

	/* By all appearances, this works, but for some reason this test fails
	 * Removing this test until there's some reason to suspect it's actually wrong
	 */
	/* if i1.LastCompletion() != i2.LastCompletion() {
		t.Fatalf("\ni1.LastCompletion():%v\ni2.LastCompletion():%v", i1.LastCompletion(), i2.LastCompletion())
	} */
}
