package finder

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// TestFinder merely invokes our constructor for sanity
func TestFinder(t *testing.T) {
	f := New()
	if f == nil {
		t.Fatalf("Failed to construct the finder")
	}
}

// TestInclusion tests we don't include directories
func TestInclusion(t *testing.T) {

	finder := New()

	if finder.ShouldInclude("/tmp") {
		t.Fatalf("We should not include directories!")
	}
	if finder.ShouldInclude("/missing/file/or/directory") {
		t.Fatalf("We should not include files that don't exist!")
	}
	if !finder.ShouldInclude("/etc/hosts") {
		t.Fatalf("We should include normal files!")
	}
}

// TestFind tests we find some files
func TestFind(t *testing.T) {

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
	err = ioutil.WriteFile(filepath.Join(p, "input"), txt, 0644)
	if err != nil {
		t.Errorf("Error writing file!")
	}

	//
	// Create our finder
	//
	finder := New()

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
}
