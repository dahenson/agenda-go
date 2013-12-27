package testutils

import . "github.com/dahenson/agenda/ui/uicallbacks"

type FakeAddButton struct {
        callback AddCallback
}

func NewFakeAddButton() *FakeAddButton {
        return &FakeAddButton{callback: func() {}}
}

func (btn *FakeAddButton) SetCallback(callback AddCallback) {
        btn.callback = callback
}

func (btn *FakeAddButton) Click() {
        btn.callback()
}

type FakeToggleButton struct {
        callback ToggleCallback
}

func NewFakeToggleButton() *FakeToggleButton {
        return &FakeToggleButton{callback: func(_ string, _ bool) {}}
}

func (btn *FakeToggleButton) SetCallback(callback ToggleCallback) {
        btn.callback = callback
}

func (btn *FakeToggleButton) Click(id string, state bool) {
        btn.callback(id, state)
}

