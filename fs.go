package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

var fs fileSystem = osFS{}

type fileSystem interface {
	ReadFile(name string) ([]byte, error)
	Append(name, text string) (int, error)
}


type osFS struct {}
func (osFS) ReadFile(name string) ([]byte, error) { return ioutil.ReadFile(name) }
func (osFS) Append(name, text string) (int, error) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return file.WriteString(text)
}

type fakeFS struct {
	files map[string][]byte
}
func newFakeFS() *fakeFS {
	return &fakeFS{files: make(map[string][]byte)}
}
func (fs *fakeFS) ReadFile(name string) ([]byte, error) {
	data, found := fs.files[name]
	if !found {
		return nil, fmt.Errorf("No such file: %s", name)
	}
	return data, nil
}
func (fs *fakeFS) Append(name, text string) (int, error) {
	if _, found := fs.files[name]; found {
		fs.files[name] = append(fs.files[name], []byte(text)...)
	} else {
		fs.files[name] = []byte(text)
	}
	return len(text), nil
}
