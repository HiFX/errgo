// Copyright 2013, 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package errgo_test

import (
	"fmt"

	"github.com/hifx/errgo"
)

func ExampleTrace() {
	var err1 error = fmt.Errorf("something wicked this way comes")
	var err2 error = nil

	// Tracing a non nil error will return an error
	fmt.Println(errgo.Trace(err1))
	// Tracing nil will return nil
	fmt.Println(errgo.Trace(err2))

	// Output: something wicked this way comes
	// <nil>
}
