agenda
======

A simple todo (task) list built in Go and GTK (by way of the Gotk3 bindings: http://github.com/conformal/gotk3)

## Installation
agenda depends on GTK+3.6.8 or greater.

1. Install Go (http://golang.org/doc/install)
2. Get the project: `$ go get github.com/dahenson/agenda` or `$ git clone` it
3. Install it
  1. If you have GTK >= 3.10 installed: `$ go install github.com/dahenson/agenda`
  2. If you have GTK >= 3.8 < 3.10 installed: `$ go install -tags=gtk_3_8 github.com/dahenson/agenda`
  3. If you have GTK >= 3.6.8 < 3.8 installed: `$ go install -tags=gtk_3_8_6 github.com/dahenson/agenda`
4. Run it: `$ agenda`
