package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	fs fileSystem = osFS{}
	filename = "default.json"
)

type fileSystem interface {
	ReadFile(name string) ([]byte, error)
	Append(name, text string) (int, error)
	IsNotExist(err error) bool
}

type osFS struct{}

// getPath() uses the XDG_DATA_HOME environmental variable to create agenda's
// working directory. If the variable is not set, it uses the default of
// "$HOME/.local/share", and creates the directory if it does not exist.
func getPath() (string, error) {
	xdg := os.Getenv("XDG_DATA_HOME")
	if xdg == "" {
		xdg = "$HOME/.local/share"
	}
	xdg = os.ExpandEnv(xdg) + "/agenda/"
	if err := os.MkdirAll(xdg, 0744); err != nil { // Create directory
		return "", err
	}
	return xdg, nil
}

// ReadFile() reads from the file identified by name
func (osFS) ReadFile(name string) ([]byte, error) {
	data, err := ioutil.ReadFile(name)
	return data, err
}

// Append() appends an item to the file
func (osFS) Append(name, text string) (int, error) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return file.WriteString(text)
}

func (osFS) IsNotExist(err error) bool {
	return os.IsNotExist(err)
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

func (fs *fakeFS) Append(name, text string) (int, error) {
	if _, found := fs.files[name]; found {
		fs.files[name] = append(fs.files[name], []byte(text)...)
	} else {
		fs.files[name] = []byte(text)
	}
	return len(text), nil
}

func (fs *fakeFS) IsNotExist(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.HasPrefix(msg, NOT_EXIST_PREFIX) && strings.HasSuffix(msg, NOT_EXIST_SUFFIX)
}
