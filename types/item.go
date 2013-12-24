package types

import "github.com/mattetti/uuid"

type Item struct {
	Text string
	Complete bool
	Id string
}

func NewItem(text string) *Item {
	return &Item{Text: text, Complete: false, Id:uuid.GenUUID()}
}
