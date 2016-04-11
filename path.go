// Copyright 2013, 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package errgo

import (
	"os"
	"runtime"
	"strings"
)

// prefixSize is used internally to trim the user specific path from the
// front of the returned filenames from the runtime call stack.
var prefixSize int

// goPath is the deduced path based on the location of this file as compiled.
var goPath string

func init() {
	_, file, _, ok := runtime.Caller(0)
	if file == "?" {
		return
	}
	if ok {
		// We know that the end of the file should be:
		// github.com/hifx/errgo/path.go
		size := len(file)
		suffix := len("github.com/hifx/errgo/path.go")
		goPath = file[:size-suffix]
		prefixSize = len(goPath)
	}
}

func trimGoPath(filename string) string {
	if strings.HasPrefix(filename, goPath) {
		return filename[prefixSize:]
	}
	return filename
}

func trimPackage(function string) string {
	slashIndex := strings.LastIndex(function, string(os.PathSeparator))
	if slashIndex < 0 {
		slashIndex = 0
	}
	dotIndex := strings.Index(function[slashIndex:], ".")
	if dotIndex == -1 {
		return function
	}
	return function[slashIndex+dotIndex+1:]
}
