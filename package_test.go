// Copyright 2013, 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package errgo_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	gc "gopkg.in/check.v1"

	"github.com/hifx/errgo"
)

func Test(t *testing.T) {
	gc.TestingT(t)
}

func checkDetails(c *gc.C, err error, details string) {
	c.Assert(err, gc.NotNil)
	expectedDetails := replaceLocations(details)
	c.Assert(errgo.Details(err), gc.Equals, expectedDetails)
}

func checkErr(c *gc.C, err, cause error, msg string, details string) {
	c.Assert(err, gc.NotNil)
	c.Assert(err.Error(), gc.Equals, msg)
	c.Assert(errgo.Cause(err), gc.Equals, cause)
	expectedDetails := replaceLocations(details)
	c.Assert(errgo.Details(err), gc.Equals, expectedDetails)
}

func replaceLocations(line string) string {
	result := ""
	for {
		i := strings.Index(line, "$")
		if i == -1 {
			break
		}
		result += line[0:i]
		line = line[i+1:]
		i = strings.Index(line, "$")
		if i == -1 {
			panic("no second $")
		}
		result += location(line[0:i]).String()
		line = line[i+1:]
	}
	result += line
	return result
}

func location(tag string) Location {
	loc, ok := tagToLocation[tag]
	if !ok {
		panic(fmt.Sprintf("tag %q not found", tag))
	}
	return loc
}

type Location struct {
	file     string
	line     int
	function string
}

func (loc Location) String() string {
	return fmt.Sprintf("%s:%d%s", loc.file, loc.line, loc.function)
}

var tagToLocation = make(map[string]Location)

func setLocationsForErrorTags(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	filename = "github.com/hifx/errgo/" + filename
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if j := strings.Index(line, "//err "); j >= 0 {
			tag := line[j+len("//err "):]
			tags := strings.Split(tag, " ")
			tag = tags[0]
			if _, found := tagToLocation[tag]; found {
				panic(fmt.Sprintf("tag %q already processed previously", tag))
			}
			l := Location{file: filename, line: i + 1}
			if len(tags) > 1 {
				l.function = " " + tags[1]
			}
			tagToLocation[tag] = l
		}
	}
}

func init() {
	setLocationsForErrorTags("error_test.go")
	setLocationsForErrorTags("functions_test.go")
}
