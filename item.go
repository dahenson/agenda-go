package main

import ()

type Item interface {
	Text() string
	Complete() bool
}

type item struct {
	text string
	complete bool
}

func (i *item) Text() string {
	return i.text
}

func (i *item) Complete() bool {
	return i.complete
}

func NewItem(text string) Item {
	return &item{text: text}
}
