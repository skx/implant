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
	"flag"
	"fmt"
	"os"
	"strings"
)

//
// Our version number.
//
var (
	version = "master/latest"
)

// inputFlags is a helper so that multiple input-directories may be accecpted.
type inputFlags []string

// String convert our input-flags to a string.
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
