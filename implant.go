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
	"text/template"
)

//
// ConfigOptions is the globally visible structure which is designed to
// hold our configuration-options - as set by the command-line flags.
//
// It is perhaps poor practice to do things this way, but it eases coding.
//
var ConfigOptions struct {
	Input   string
	Output  string
	Exclude string
	Package string
	Format  bool
	Verbose bool
	Version bool
}

//
// Our version number.
//
var (
	version = "master/latest"
)

//
// Resource is the structure used to temporarily hold data about all the
// static-resources we've discovered and will write out to the generated
// `static.go` file.
//
type Resource struct {
	Filename string
	Contents string
	Length   int
}

//
// CheckInputLocation is designed to test that the input path
// specified is a directory.
//
// If this function returns false then the application terminates.
//
func CheckInput() bool {

	stat, err := os.Stat(ConfigOptions.Input)
	if err != nil {
		fmt.Printf("Failed to stat %s - Did you forget to specify a directory to read?\n", ConfigOptions.Input)
		return false
	}

	//
	// Test that the input path is a directory.
	//
	if !stat.IsDir() {
		fmt.Printf("Error %s is not a directory!\n", ConfigOptions.Input)
		return false
	}

	return true
}

//
// ShouldInclude is invoked by our filesystem-walker, and determines whether
// any particular directory entry beneath the input tree should be included
// in our generated `static.go` file.
//
// We skip directories, and we might skip particular entries via a user-supplied
// regular expression too.
//
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

//
// Find all the files in the given directory, returning an array
// of structures.
//
func findFiles(input string) ([]Resource, error) {

	// The resources we've found.
	var entries []Resource

	// The list of files we'll add
	fileList := []string{}

	// Recursively look for files to add.
	err := filepath.Walk(input, func(path string, f os.FileInfo, err error) error {
		if ConfigOptions.Verbose {
			fmt.Printf("Found file: %s\n", path)
		}
		if ShouldInclude(path) {
			if ConfigOptions.Verbose {
				fmt.Printf("Including file: %s\n", path)
			}
			fileList = append(fileList, path)
		} else {
			if ConfigOptions.Verbose {
				fmt.Printf("Excluding file: %s\n", path)
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

//
// Given the array of resource-entries render our template.
//
func renderTemplate(entries []Resource) (string, error) {

	//
	// Create our template object
	//
	tmpl, err := getResource("data/static.tmpl")
	if err != nil {
		return "", err
	}

	t := template.Must(template.New("tmpl").Parse(string(tmpl)))

	//
	// The data we add to our template.
	//
	var Templatedata struct {
		Resources []Resource
		Package   string
	}

	//
	// Populate the template-data
	//
	Templatedata.Package = ConfigOptions.Package
	Templatedata.Resources = entries

	//
	// Execute into a temporary buffer.
	//
	buf := &bytes.Buffer{}
	err = t.Execute(buf, Templatedata)

	//
	// If there were errors, then show them.
	if err != nil {
		return "", err
	}

	//
	// Otherwise write the result.
	//
	return buf.String(), nil
}

//
// This is our entry-point.
//
func main() {

	//
	// The command-line flags we support
	//
	flag.StringVar(&ConfigOptions.Exclude, "exclude", "", "A regular expression of files to ignore, for example '.git'.")
	flag.StringVar(&ConfigOptions.Input, "input", "data/", "The directory to read from.")
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
	// Showing our version?
	//
	if ConfigOptions.Version {
		fmt.Printf("implant %s\n", version)
		os.Exit(0)
	}

	//
	// If we're running verbosely show our settings.
	//
	if ConfigOptions.Verbose {
		fmt.Printf("Reading input directory %s\n", ConfigOptions.Input)
	}

	//
	// Check we've been given an input path that a) exists and b) is
	// a directory.
	//
	if !CheckInput() {
		os.Exit(1)
	}

	//
	// Now find the files beneath that path.
	//
	files, err := findFiles(ConfigOptions.Input)
	if err != nil {
		fmt.Printf("Error processing files: %s\n", err.Error())
		os.Exit(1)
	}

	//
	// If there were no files found we should abort.
	//
	if len(files) < 1 {
		fmt.Printf("Failed to find files beneath %s\n", ConfigOptions.Input)
		os.Exit(1)
	}

	//
	// Show how many files we found.
	//
	if ConfigOptions.Verbose {
		fmt.Printf("Populating static.go with the following files.\n")
		for _, ent := range files {
			fmt.Printf("\t%s\n", ent.Filename)
		}
	}

	//
	// Render our template with their details
	//
	tmpl, err := renderTemplate(files)
	if err != nil {
		panic(err)
	}

	//
	// Now we pipe our generated template through `gofmt`
	//
	if ConfigOptions.Format {
		if ConfigOptions.Verbose {
			fmt.Printf("Piping our output through `gofmt` and into %s\n", ConfigOptions.Output)
		}

		fmtCmd := exec.Command("gofmt")
		In, _ := fmtCmd.StdinPipe()
		Out, _ := fmtCmd.StdoutPipe()
		fmtCmd.Start()
		In.Write([]byte(tmpl))
		In.Close()
		Bytes, _ := ioutil.ReadAll(Out)
		fmtCmd.Wait()
		ioutil.WriteFile(ConfigOptions.Output, Bytes, 0644)
	} else {
		if ConfigOptions.Verbose {
			fmt.Printf("Writing output to %s\n", ConfigOptions.Output)
		}
		ioutil.WriteFile(ConfigOptions.Output, []byte(tmpl), 0644)
	}

}
