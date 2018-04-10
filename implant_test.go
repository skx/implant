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
	if err != nil {
		t.Errorf("Failed to write our data to a file")
	}
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
		{"tgit", true}, // Excluded because ".git" matches "tgit"
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
		if err != nil {
			t.Errorf("Failed to write our data to a file")
		}

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

//
// Test we can find files.
//
func TestFileFinding(t *testing.T) {

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
	// Create a single file.
	//
	txt := []byte("hello, world!\n")
	err = ioutil.WriteFile(filepath.Join(p, "bar"), txt, 0644)

	//
	// Find our files.
	//
	out, err := findFiles(ConfigOptions.Input)
	if err != nil {
		t.Errorf("Error finding files!")
	}

	//
	// We should have one output result.
	//
	if len(out) != 1 {
		t.Errorf("We expected to find one file!")
	}

	//
	// Cleanup our temporary directory
	//
	os.RemoveAll(ConfigOptions.Input)

}

//
// Test we can output a template.
//
func TestOutputTemplate(t *testing.T) {

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
	ConfigOptions.Package = "main"

	//
	// Create a single file.
	//
	txt := []byte("hello, world!\n")
	err = ioutil.WriteFile(filepath.Join(p, "input"), txt, 0644)

	//
	// Find our files.
	//
	out, err := findFiles(ConfigOptions.Input)
	if err != nil {
		t.Errorf("Error finding files!")
	}

	//
	// Render our template
	//
	out2, err := renderTemplate(out)
	if err != nil {
		t.Errorf("Error rendering template")
	}

	//
	// Ensure that our output looks valid.
	//
	if len(out2) < 1 {
		t.Errorf("Rendered template was empty")
	}

	if !strings.Contains(out2, "package main") {
		t.Errorf("Rendered template was not in the main-package")
	}

	//
	// Cleanup our temporary directory
	//
	os.RemoveAll(ConfigOptions.Input)

}

//
// Test we can sanity-check our input path.
//
func TestInputDirectory(t *testing.T) {

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

	//
	// Test a directory
	//
	if !CheckInput() {
		t.Errorf("A valid directory wasn't accepted.")
	}

	//
	// Test a missing thing
	//
	ConfigOptions.Input = filepath.Join(p, "missing.ent")
	if CheckInput() {
		t.Errorf("A missing file was accepted.")
	}

	//
	// Test a file, rather than a directory
	//
	txt := []byte("hello, world!\n")
	err = ioutil.WriteFile(filepath.Join(p, "bar"), txt, 0644)
	if err != nil {
		t.Errorf("Failed to write our data to a file")
	}

	ConfigOptions.Input = filepath.Join(p, "bar")
	if CheckInput() {
		t.Errorf("A missing file was accepted.")
	}

	//
	// Cleanup our temporary directory
	//
	os.RemoveAll(ConfigOptions.Input)

}
