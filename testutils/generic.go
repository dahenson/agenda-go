package testutils

import (
	"fmt"
	. "github.com/dahenson/agenda/types"
)

func Expect(exp, act interface{}) error {
	if exp != act {
		return fmt.Errorf("Expected \"%v\"; Got \"%v\"", exp, act)
	}
	return nil
}

func eq(i1, i2 *Item) bool {
	return i1.Id == i2.Id && i1.Text == i2.Text && i1.Complete == i2.Complete
}

func ExpectItem(exp, act *Item) error {
	if eq(exp, act) {
		return fmt.Errorf("Expected \"%v\"; Got \"%v\"", exp, act)
	}
	return nil
}

func ExpectCount(exp, act int) error {
	if exp != act {
		return fmt.Errorf("Expected %d elements; Got %d", exp, act)
	}
	return nil
}
