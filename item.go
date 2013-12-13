package main

import ()

type Item interface {
	Text() string
}

type item struct {
	text string
}

func (i *item) Text() string {
	return i.text
}

func NewItem(text string) Item {
	return &item{text: text}
}
