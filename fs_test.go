package main

import (
	"os"
	"testing"
)

var (
	home string
	xdg_default = "/.local/share/agenda/"
)

func testInit() error {
	home = os.ExpandEnv("$HOME")
	if xdg := os.Getenv("XDG_DATA_HOME"); xdg != "" {
		err := os.Setenv("XDG_DATA_HOME", "") // Need to "unset" XDG_DATA_HOME
		if err != nil {
			return err
		}
	}
	return nil
}

func TestGetDir_WhenEnvNotSet_ExpectDefault(t *testing.T) {
	if err := testInit(); err != nil {
		t.Fatalf("Could not initialize test: ", err)
	}
	expected := home + "/.local/share/agenda/"
	path, err := getPath()
	if err != nil {
		t.Fatalf("Unexpected error: ", err)
	}
	if path != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, path)
	}
}

func TestGetDir_WhenEnvSet_ExpectEnv(t *testing.T) {
	if err := testInit(); err != nil {
		t.Fatalf("Could not initialize test: ", err)
	}
	if err := os.Setenv("XDG_DATA_HOME", home + "/.local/share/test"); err != nil {
		t.Fatalf("Could not initialize test: ", err)
	}
	expected := home + "/.local/share/test/agenda/"
	path, err := getPath()
	if err != nil {
		t.Fatalf("Unexpected error: ", err)
	}
	if path != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, path)
	}
}
