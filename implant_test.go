//
// Simple testing of our package.
//
//
package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/skx/implant/finder"
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

func TestChoose(t *testing.T) {

	//
	// Test an empty array
	//
	all := []finder.Resource{}
	out := Choose(all, TestRegexp)
	if len(out) != 0 {
		t.Errorf("Failed to filter an empty array!")
	}

	//
	// Now test a real regexp against these inputs:
	//
	in := []finder.Resource{
		finder.Resource{Filename: "def", Contents: "Steve", Length: 5},
		finder.Resource{Filename: "abc", Contents: "Kemp", Length: 4}}

	//
	// Setup a regexp which will exclude no files.
	//
	ConfigOptions.Exclude = ""
	out = Choose(in, TestRegexp)
	if len(out) != 2 {
		t.Errorf("Expected all entries to be present, found only:%d", len(out))
	}

	//
	// Setup a regexp which will exclude all files.
	//
	ConfigOptions.Exclude = "..."
	out = Choose(in, TestRegexp)
	if len(out) != 0 {
		t.Errorf("Expected all entries to be filtered, found:%d", len(out))
	}

	//
	// Setup a regexp which will exclude only half
	//
	ConfigOptions.Exclude = "a"
	out = Choose(in, TestRegexp)
	if len(out) != 1 {
		t.Errorf("Expected one entry to be filtered, found:%d", len(out))
	}

}

func TestRenderTemplate(t *testing.T) {

	//
	// Create a temporary directory
	//
	p, err := ioutil.TempDir(os.TempDir(), "prefix")
	if err != nil {
		t.Errorf("Error setting up test.")
	}
	defer os.RemoveAll(p)

	//
	// Create a single file.
	//
	txt := []byte("hello, world!\n")
	err = ioutil.WriteFile(filepath.Join(p, "moi.kissa"), txt, 0644)
	if err != nil {
		t.Errorf("Error writing file!")
	}

	//
	// Create our finder
	//
	finder := finder.New()

	//
	// Find our files
	//
	resources, err := finder.FindFiles(p)
	if err != nil {
		t.Fatalf("We shouldn't have an error")
	}

	if len(resources) != 1 {
		t.Fatalf("We expected to find 1 file, but found %d", len(resources))
	}

	//
	// Now render a template.
	//
	str, err := RenderTemplate(resources)
	if err != nil {
		t.Fatalf("Found error rendering our template")
	}
	if !strings.Contains(str, "moi.kissa") {
		t.Fatalf("Our rendered template didn't contain our filename")
	}
}
