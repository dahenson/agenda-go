package types

// Old and busted import "github.com/mattetti/uuid"
/* New hotness */ import "github.com/germ/go-bits/puuid"

type Item struct {
	Text string
	Complete bool
	Id string
}

func NewItem(text string) *Item {
	return &Item{Text: text, Complete: false, Id:puuid.Generate()}
}
