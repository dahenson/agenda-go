package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type FileSystem interface {
	ReadFile(name string) ([]byte, error)
	Write(name string, data []byte) (int, error)
	IsNotExist(err error) bool
}

type osFS struct {}

// GetPath() uses the XDG_DATA_HOME environmental variable to create agenda's
// working directory. If the variable is not set, it uses the default of
// "$HOME/.local/share".
func GetPath() string {
	xdg := os.Getenv("XDG_DATA_HOME")
	if xdg == "" {
		xdg = "$HOME/.local/share"
	}
	xdg = os.ExpandEnv(xdg) + "/agenda/"
	return xdg
}

// ReadFile() reads from the file identified by name
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

var fs FileSystem = osFS{}

func OsFs() FileSystem { return fs }

const (
	NOT_EXIST_PREFIX = "open "
	NOT_EXIST_SUFFIX = ": no such file or directory"
)
type FakeFS struct {
	files map[string][]byte
}
func NewFakeFS() *FakeFS {
	return &FakeFS{files: make(map[string][]byte)}
}

func doesntExistErr(filename string) error {
	return fmt.Errorf(NOT_EXIST_PREFIX+"%s"+NOT_EXIST_SUFFIX, filename)
}
func (fs *FakeFS) ReadFile(name string) ([]byte, error) {
	data, found := fs.files[name]
	if !found {
		return nil, doesntExistErr(name)
	}
	return data, nil
}
func (fs *FakeFS) Write(name string, data []byte) (int, error) {
	fs.files[name] = data
	return len(data), nil
}
func (fs *FakeFS) IsNotExist(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.HasPrefix(msg, NOT_EXIST_PREFIX) && strings.HasSuffix(msg, NOT_EXIST_SUFFIX)
}
