package finder

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"testing"
)

// TestFinder merely invokes our constructor for sanity
func TestFinder(t *testing.T) {
	f := New()
	if f == nil {
		t.Fatalf("Failed to construct the finder")
	}
}

// TestInclusion tests we don't include directories, but do include
// files we expect.
func TestInclusion(t *testing.T) {

	finder := New()
	found, err := finder.FindFiles("../_tests")

	if err != nil {
		t.Errorf("Unexpected error finding files: %s", err.Error())
	}

	//
	// OK we've found files beneath our top-level
	// _tests directory.
	//
	// We know what that should contain, so we can
	// test the sizes.
	//
	if len(found) != 3 {
		t.Fatalf("We found an unexpected number of files: %d", len(found))
	}

	//
	// Ensure we don't have the directories
	//
	if finder.ShouldInclude("/missing/file/or/directory") {
		t.Fatalf("We should not include files that don't exist!")
	}
	if finder.ShouldInclude("../_test/etc") {
		t.Fatalf("We should not include directories!")
	}
}

// TestFind ensures found-files have our expected content.
func TestFind(t *testing.T) {

	finder := New()
	found, err := finder.FindFiles("../_tests")

	if err != nil {
		t.Errorf("Unexpected error finding files: %s", err.Error())
	}

	//
	// Process the file named `etc/foo/bar/baz.txt`.
	//
	for _, entry := range found {

		if entry.Filename == "../_tests/etc/foo/bar/baz.txt" {

			if entry.Length != 15 {
				t.Errorf("File has wrong length: %d\n", entry.Length)
			}

			//
			// Expected content
			//
			c := "This is a test\n"

			//
			// Decode the contents.
			//
			decoded, err := base64.StdEncoding.DecodeString(entry.Contents)
			if err != nil {
				t.Errorf("Error decoding content: %s", err.Error())

			}

			// Gunzip the data to the client
			gr, err := gzip.NewReader(bytes.NewBuffer(decoded))
			if err != nil {
				t.Errorf("Error creating gzip-reader: %s", err.Error())
			}
			defer gr.Close()
			data, err := ioutil.ReadAll(gr)
			if err != nil {
				t.Errorf("Error gunzipping: %s", err.Error())
			}

			if string(data) != c {
				t.Errorf("File has wrong content: '%s'", entry.Contents)
			}
			return
		}
	}

	t.Errorf("We didn't find the file we expected in our walk")
}
