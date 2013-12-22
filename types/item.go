package types

type Item struct {
	Text string
	Complete bool
}

func NewItem(text string) *Item {
	return &Item{Text: text}
}
