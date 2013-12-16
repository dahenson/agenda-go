package main

import (
	"github.com/mattetti/uuid"
)

type Item interface {
	Text() string
	Id() string
}

type item struct {
	text string
	id   string
}

// Text returns the item's text
func (i *item) Text() string {
	return i.text
}

// Id returns the item's Id
func (i *item) Id() string {
	return i.id
}

// NewItem creates a new item with some text and default values
func NewItem(text string) Item {
	return &item{text: text, id: uuid.GenUUID()}
}
