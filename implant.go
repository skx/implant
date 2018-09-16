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
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"text/template"

	"github.com/skx/implant/finder"
)

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

// RenderTemplate populates our template-file with the list of resources
// which have been discovered.
func RenderTemplate(entries []finder.Resource) (string, error) {

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
		Resources []finder.Resource
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
	// Return the buffer contents and/or the error.
	//
	return buf.String(), err
}

// Choose lets us perform a test against each member of an array.
func Choose(ss []finder.Resource, test func(finder.Resource) bool) (ret []finder.Resource) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

// TestRegexp tests whether the specified resource should be excluded,
// based on the regular expression a user might have optional specified.
//
// If the function returns true the resource is included in our generated
// output, when false it is not.
func TestRegexp(x finder.Resource) bool {

	if ConfigOptions.Exclude == "" {
		return true
	}

	match, _ := regexp.MatchString(ConfigOptions.Exclude, x.Filename)
	return match == false
}

// Implant is our main entry-point.
func Implant() {

	//
	// We'll build up a list of all resources which were found.
	//
	var all []finder.Resource

	//
	// Process each directory which was specified.
	//
	for _, directory := range ConfigOptions.Input {

		if ConfigOptions.Verbose {
			fmt.Printf("Processing directory %s\n", directory)
		}

		//
		// Find the files beneath the input-directory.
		//
		helper := finder.New()
		entries, err := helper.FindFiles(directory)

		if err != nil {
			fmt.Printf("Failed to find files beneath %s - %s\n", directory, err.Error())
			continue
		}

		//
		// Append each found-resource to the list.
		//
		for _, entry := range entries {
			all = append(all, entry)
		}
	}

	//
	// Now exclude the resources we don't care about, because the
	// entries match the regular expression the user specified.
	//
	resources := Choose(all, TestRegexp)

	//
	// Now render our template with the found-resources.
	//
	tmpl, err := RenderTemplate(resources)
	if err != nil {
		fmt.Printf("Failed to render template: %s\n", err.Error())
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
