//
// Simple testing of the HTTP-server
//
//
package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

//
// Test that files/directories are included/excluded as expected.
//
func TestSimpleExclusions(t *testing.T) {

	//
	// Create a temporary directory
	//
	p, err := ioutil.TempDir(os.TempDir(), "prefix")
	if err != nil {
		t.Errorf("Error setting up test.")
	}

	//
	// Setup our options.
	//
	ConfigOptions.Input = p
	ConfigOptions.Verbose = true

	//
	// Create a directory, and test it shouldn't be included.
	//
	os.Mkdir(filepath.Join(p, "foo"), 0777)
	if ShouldInclude(filepath.Join(p, "foo")) {
		t.Errorf("We shouldn't include a directory")
	}

	//
	// This is a simple error-case.
	//
	ShouldInclude(filepath.Join(p, "missing.file"))

	//
	// Create a file and test it should be included
	//
	txt := []byte("hello, world!\n")
	err = ioutil.WriteFile(filepath.Join(p, "bar"), txt, 0644)
	if !ShouldInclude(filepath.Join(p, "bar")) {
		t.Errorf("We should include a file")
	}

	// Cleanup our temporary directory
	//
	os.RemoveAll(ConfigOptions.Input)

}

//
// Test that files are excluded via regular expressions.
//
func TestRegexpExclusions(t *testing.T) {

	//
	// Create a temporary directory
	//
	p, err := ioutil.TempDir(os.TempDir(), "prefix")
	if err != nil {
		t.Errorf("Error setting up test.")
	}

	//
	// Setup our options.
	//
	ConfigOptions.Input = p
	ConfigOptions.Verbose = true
	ConfigOptions.Exclude = "/.git"

	//
	// We'll test that each of these files should be missing.
	//
	type TestCase struct {
		Filename string
		Exclude  bool
	}

	//
	// Now our tests
	//
	tests := []TestCase{
		{"test", false},
		{"tgit", true},
		{"git", false},
		{".git", true},
		{".gitignore", true}}

	for _, entry := range tests {

		//
		// Create a file and test it should be included
		//
		txt := []byte("hello, world!\n")
		path := filepath.Join(p, entry.Filename)
		err = ioutil.WriteFile(path, txt, 0644)

		out := ShouldInclude(path)

		if out != !entry.Exclude {
			t.Errorf("Regexp exclusion failed for %s, got %v expected %v", entry.Filename, out, entry.Exclude)
		}
	}

	//
	// Cleanup our temporary directory
	//
	os.RemoveAll(ConfigOptions.Input)

}
