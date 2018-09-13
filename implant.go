//
//  This is a simple script which is designed to output a list of
// static resources to a golang source-file which can be then compiled
// into your program.
//
//  The program basically reads all files in a given directory and spits
// out a templated file.
//

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

// Created so that multiple inputs can be accecpted
type inputFlags []string

// Convert our input-flags to a string
func (i *inputFlags) String() string {
	return strings.Join(*i, ",")
}

// Set adds an -input-flag to our list of directories to scan.
func (i *inputFlags) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

// ConfigOptions is the globally visible structure which is designed to
// hold our configuration-options - as set by the command-line flags.
//
// It is perhaps poor practice to do things this way, but it eases coding.
var ConfigOptions struct {
	// Input contains the directories we should search for files to include
	Input inputFlags

	// Output is used to hold the name of the file we generate.
	Output string

	// Exclude contains a regular expression of files to exclude from
	// the implantation process
	Exclude string

	// Package contains the package to which the generated file should be
	// a member of.  (By default `main`.)
	Package string

	// Format controls whether we format the generated file with `go fmt`
	Format bool

	// Verbose is set to increase verbosity.
	Verbose bool

	// Version is used to determine whether we should just report our
	// version-number then terminate.
	Version bool
}

//
// Our version number.
//
var (
	version = "master/latest"
)

// Resource is the structure used to temporarily hold data about all the
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

// CheckInput is designed to test that the input path
// specified is a directory.
//
// If this function returns false then the application terminates.
func CheckInput(directory string) bool {

	stat, err := os.Stat(directory)
	if err != nil {
		fmt.Printf("Failed to stat %s - Does it exist?\n", directory)
		return false
	}

	//
	// Test that the input path is a directory.
	//
	if !stat.IsDir() {
		fmt.Printf("Error %s is not a directory!\n", directory)
		return false
	}

	return true
}

//
// PipeCommand pipes some data through a binary, and returns the output.
//
func PipeCommand(cmd string, input []byte) []byte {

	fmtCmd := exec.Command("gofmt")
	In, _ := fmtCmd.StdinPipe()
	Out, _ := fmtCmd.StdoutPipe()
	fmtCmd.Start()
	In.Write(input)
	In.Close()
	output, _ := ioutil.ReadAll(Out)
	fmtCmd.Wait()

	return output
}

// ShouldInclude is invoked by our filesystem-walker, and determines whether
// any particular directory-entry beneath the input tree should be included
// in our generated `static.go` file.
//
// We skip directories, and we might skip particular entries via a user-supplied
// regular expression too.
func ShouldInclude(path string) bool {

	//
	// Examine the file.
	//
	stat, err := os.Stat(path)
	if err != nil {
		if ConfigOptions.Verbose {
			fmt.Printf("Failed to stat file : %s\n", path)
		}
		return false
	}

	//
	// Is it a regular file?
	//
	mode := stat.Mode()
	if !mode.IsRegular() {
		//
		// If it isn't we shouldn't include it.
		//
		return false
	}

	//
	// If we have a regular expression then test against it.
	//
	if ConfigOptions.Exclude != "" {

		//
		// If the regexp matches then we must EXCLUDE
		// the file.
		//
		match, _ := regexp.MatchString(ConfigOptions.Exclude, path)
		if match {
			return false
		}
	}

	//
	// OK the file wasn't excluded, so we'll add it.
	//
	return true
}

// findFiles finds all the files in the given directory, returning an array
// of Resource-structures to describe each one.
func findFiles(directory string) ([]Resource, error) {

	// The resources we've found.
	var entries []Resource

	// The list of files we'll add
	fileList := []string{}

	// Recursively look for files to add.
	err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		if ConfigOptions.Verbose {
			fmt.Printf("Found file: %s\n", path)
		}
		if ShouldInclude(path) {
			if ConfigOptions.Verbose {
				fmt.Printf("\tIncluding file: %s\n", path)
			}
			fileList = append(fileList, path)
		} else {
			if ConfigOptions.Verbose {
				fmt.Printf("\tExcluding file: %s\n", path)
			}

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

// renderTemplate populates our template-file with the list of resources
// which have been discovered.
func renderTemplate(entries []Resource) (string, error) {

	//
	// Get our template-text.
	//
	tmpl, err := getResource("data/static.tmpl")
	if err != nil {
		return "", err
	}

	//
	// Compile it.
	//
	t := template.Must(template.New("tmpl").Parse(string(tmpl)))

	//
	// Populate the template-data
	//
	var Templatedata struct {
		Resources []Resource
		Package   string
	}
	Templatedata.Package = ConfigOptions.Package
	Templatedata.Resources = entries

	//
	// Execute the template into a temporary buffer.
	//
	buf := &bytes.Buffer{}
	err = t.Execute(buf, Templatedata)

	//
	// If there was an error return it.
	if err != nil {
		return "", err
	}

	//
	// Otherwise return the result of the template-compilation.
	//
	return buf.String(), nil
}

// Implant is our main entry-point.
func Implant() {

	//
	// We'll build up a list of resources which will be
	// inserted into our template.
	//
	var resources []Resource

	//
	// Process each directory which was specified.
	//
	for _, directory := range ConfigOptions.Input {

		//
		// If we're running verbosely show our settings.
		//
		if ConfigOptions.Verbose {
			fmt.Printf("Reading input directory %s\n", directory)
		}

		//
		// Check we've been given an input path that a) exists and b) is
		// a directory.
		//
		if !CheckInput(directory) {
			os.Exit(1)
		}

		//
		// Now find the files beneath that path.
		//
		files, err := findFiles(directory)
		if err != nil {
			fmt.Printf("Error processing files: %s\n", err.Error())
			return
		}

		//
		// Append the files to our result-set
		//
		for _, entry := range files {
			resources = append(resources, entry)
		}

		//
		// If a given directory resulted in no input then report that.
		//
		if len(files) < 1 {
			fmt.Printf("Failed to find files beneath %s\n", directory)
			break
		}
	}

	//
	// Show how many files we found.
	//
	if ConfigOptions.Verbose {
		fmt.Printf("Populating %s with the following files:\n", ConfigOptions.Output)
		for _, ent := range resources {
			fmt.Printf("\t%s\n", ent.Filename)
		}
	}

	//
	// Now render our template with their details
	//
	tmpl, err := renderTemplate(resources)
	if err != nil {
		fmt.Printf("Failed to render template %s\n", err.Error())
		return
	}

	//
	// This is the output of the template
	//
	output := []byte(tmpl)

	//
	// Optionally we pipe our generated output through `gofmt`
	//
	if ConfigOptions.Format {
		if ConfigOptions.Verbose {
			fmt.Printf("Piping our output through `gofmt` and into %s\n", ConfigOptions.Output)
		}

		output = PipeCommand("gofmt", []byte(tmpl))
	}

	//
	// Write our output, formatted or not, to our file.
	//
	ioutil.WriteFile(ConfigOptions.Output, output, 0644)
}

// main is our entry-point.
func main() {

	//
	// The command-line flags we support
	//
	flag.StringVar(&ConfigOptions.Exclude, "exclude", "", "A regular expression of files to ignore, for example '.git'.")
	flag.Var(&ConfigOptions.Input, "input", "The directory to read from.")
	flag.StringVar(&ConfigOptions.Output, "output", "static.go", "The output file to generate.")
	flag.StringVar(&ConfigOptions.Package, "package", "main", "The (go) package that the generated file is part of.")
	flag.BoolVar(&ConfigOptions.Verbose, "verbose", false, "Should we be verbose.")
	flag.BoolVar(&ConfigOptions.Format, "format", true, "Should we pipe our template through 'gofmt'?")
	flag.BoolVar(&ConfigOptions.Version, "version", false, "Should we report our version and exit?")

	//
	// Parse the flags
	//
	flag.Parse()

	//
	// If we received no input-directory/input-directories then we
	// should default to processing the contents of ./data
	//
	if len(ConfigOptions.Input) == 0 {
		ConfigOptions.Input = append(ConfigOptions.Input, "./data")
	}

	//
	// Showing our version?
	//
	if ConfigOptions.Version {
		fmt.Printf("implant %s\n", version)
		os.Exit(0)
	}

	//
	// Otherwise process the directories.
	//
	Implant()
}
