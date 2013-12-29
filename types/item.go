package types

import (
	// Old and busted "github.com/mattetti/uuid"
	/* New hotness */ "github.com/germ/go-bits/puuid"
	"time"
	"encoding/json"
)

// a workaround to make marshalable the unexported members of Item
type itemData struct {
	Text     string
	LastCompletion time.Time
	Complete bool
	Id       string
}

type Item struct {
	itemData
}

func NewItem(text string) *Item {
	return &Item{itemData{Text: text, Complete: false, Id: puuid.Generate()}}
}

func NewItemFromData(id, text string, complete bool, lastCompletion time.Time) *Item {
	return &Item{itemData{
		Id: id,
		Text: text,
		Complete: complete,
		LastCompletion: lastCompletion,
	}}
}

// Resets the last completion time to `time.Now()` if `complete` is `true`
func (i *Item) SetComplete(complete bool) {
	if complete {
		i.itemData.LastCompletion = time.Now()
	}
	i.itemData.Complete = complete
}

func (i *Item) Complete() bool {
	return i.itemData.Complete
}

func (i *Item) Id() string {
	return i.itemData.Id
}

func (i *Item) CompletedBefore(other *Item) bool {
	return i.itemData.LastCompletion.Before(other.itemData.LastCompletion)
}

func (i *Item) LastCompletion() time.Time {
	return i.itemData.LastCompletion
}

func (i *Item) Text() string {
	return i.itemData.Text
}

func (i *Item) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.itemData)
}

func (i *Item) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &i.itemData)
}
