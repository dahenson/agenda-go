package testutils

type FakeEntry struct {
        text string
}
func NewFakeEntry() *FakeEntry { return &FakeEntry{text: ""} }
func (e *FakeEntry) SetText(text string) { e.text = text }
func (e *FakeEntry) GetText() string { return e.text}

