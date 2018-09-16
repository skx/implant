//
// Simple testing of our package.
//
//
package main

import (
	"strings"
	"testing"
)

//
// Test invoking `gofmt` in a filter.
//
func TestFilter(t *testing.T) {

	//
	// The "program" we're filtering - note the excessive whitespace.
	//
	in := []byte(" package     main\n")

	//
	// Pipe + get the output
	//
	out := PipeCommand("gofmt", in)
	str := string(out)

	//
	// Look for it to be corrected.
	//
	if !strings.Contains(str, "package main\n") {
		t.Errorf("Our filtering didn't work?")
	}
}
