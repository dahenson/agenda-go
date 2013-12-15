package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

var fs fileSystem = osFS{}

type fileSystem interface {
	ReadFile(name string) ([]byte, error)
	Write(name string, data []byte) (int, error)
	IsNotExist(err error) bool
}

type osFS struct {}
func (osFS) ReadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}
func (osFS) IsNotExist(err error) bool {
	if err == nil {
		return false
	}
	return os.IsNotExist(err)
}
func (osFS) Write(name string, data []byte) (int, error) {
	file, err := os.Create(name)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(data)
}

const (
	NOT_EXIST_PREFIX = "open "
	NOT_EXIST_SUFFIX = ": no such file or directory"
)
type fakeFS struct {
	files map[string][]byte
}
func newFakeFS() *fakeFS {
	return &fakeFS{files: make(map[string][]byte)}
}
func doesntExistErr(filename string) error {
	return fmt.Errorf(NOT_EXIST_PREFIX+"%s"+NOT_EXIST_SUFFIX, filename)
}
func (fs *fakeFS) ReadFile(name string) ([]byte, error) {
	data, found := fs.files[name]
	if !found {
		return nil, doesntExistErr(name)
	}
	return data, nil
}

func (fs *fakeFS) Write(name string, data []byte) (int, error) {
	fs.files[name] = data
	return len(data), nil
}

func (fs *fakeFS) IsNotExist(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.HasPrefix(msg, NOT_EXIST_PREFIX) && strings.HasSuffix(msg, NOT_EXIST_SUFFIX)
}
