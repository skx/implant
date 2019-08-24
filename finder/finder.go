// Package finder is a package for finding resources.
package finder

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Resource is the structure used to hold state about each of the
// static-resources we've discovered and will write out to the generated
// `static.go` file.
type Resource struct {

	// Filename holds the filename of the resource we've discovered.
	Filename string

	// Contents holds the string contents of the file.
	Contents string

	// Length contains the length of the input file.
	Length int
}

// Finder holds our object-state.
type Finder struct {
}

// New is the constructor for our object
func New() *Finder {
	m := new(Finder)
	return m
}

// ShouldInclude is invoked by our filesystem-walker, and determines whether
// any particular directory-entry beneath the input tree should be included
// in our generated `static.go` file.
//
// We skip directories, and otherwise add all files.
func (f *Finder) ShouldInclude(path string) bool {

	//
	// Examine the file.
	//
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	//
	// Is it a regular file?
	//
	mode := stat.Mode()

	//
	// We only add regular files.
	//
	return mode.IsRegular()
}

// FindFiles finds all the files in the given directory, returning an array
// of Resource-structures to describe each one.
func (f *Finder) FindFiles(directory string) ([]Resource, error) {

	// The resources we've found.
	var entries []Resource

	// The list of files we'll add
	fileList := []string{}

	// Recursively look for files to add.
	err := filepath.Walk(directory, func(path string, file os.FileInfo, err error) error {
		if f.ShouldInclude(path) {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// For each file we'll now add it to our list.
	for _, file := range fileList {

		//
		// The entry for this file
		//
		var tmp Resource

		//
		// Read the file contents.
		//
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		//
		// gzip the data.
		//
		var gzipped bytes.Buffer
		gw, err := gzip.NewWriterLevel(&gzipped, gzip.BestSpeed)
		if err != nil {
			return nil, err
		}
		_, err = gw.Write(data)
		if err != nil {
			return nil, err
		}
		gw.Close()

		//
		// Add the filename + data, which is now encoded
		// such that it is printable in our template.
		//
		tmp.Filename = file
		tmp.Contents = hex.EncodeToString(gzipped.Bytes())
		tmp.Length = len(data)
		entries = append(entries, tmp)
	}
	return entries, nil
}
