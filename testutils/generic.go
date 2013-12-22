package testutils

import "fmt"

func ExpectItemCount(exp, act int) error {
        if exp != act {
                return fmt.Errorf("Expected %d items, got %d", exp, act)
        }
        return nil
}

func ExpectText(exp, act string) error {
        if exp != act {
                return fmt.Errorf("Expected '%s'; Got '%s'", exp, act)
        }
        return nil
}
