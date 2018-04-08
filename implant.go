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
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

//
// This is the template which is used to generate the output file.
//
var TEMPLATE = `
//
// This file was generated via github.com/skx/implant/
//
// Local edits will be lost.
//
package main

import (
    "bytes"
    "compress/gzip"
    "encoding/hex"
    "io"
    "io/ioutil"
    "errors"
)

//
// The definition of our resources
//
type EmbeddedResource struct {
    Filename string
    Contents string
}

//
// The list of our resources
//
var RESOURCES []EmbeddedResource

//
// Populate our resources
//
func init() {

    var tmp EmbeddedResource

    {{ range . }}
	tmp.Filename = "{{ .Filename }}"
        tmp.Contents = "{{ .Contents }}"
        RESOURCES = append( RESOURCES, tmp )
    {{end}}
}

// Write gunzipped data to a Writer
func gunzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	defer gr.Close()
	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}

//
// Return the contents of a resource.
//
func getResource( path string  ) ([]byte, error) {
    for _, entry := range( RESOURCES ) {
        if ( entry.Filename == path ) {
			//
			// We found the file contents.
			// Return the decoded version.
			//
			var raw bytes.Buffer
			in, err := hex.DecodeString(entry.Contents)
			if err != nil {
				return nil, err
			}
			err = gunzipWrite(&raw, in)
			if err != nil {
				return nil, err
			}
			return raw.Bytes(), nil
        }
    }
    return nil, errors.New( "Failed to find resource")
}
`

//
// This is the structure of files we've found, which we'll use to
// populate the output file.
//
type Resource struct {
	Filename string
	Contents string
}

// Write gzipped data to a Writer
func gzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gw, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	defer gw.Close()
	gw.Write(data)
	return err
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
		//
		// We only care about plain-files
		//
		// So we skip directories, symlinks, etc.
		//
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		switch mode := stat.Mode(); {

		case mode.IsRegular():

			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

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

		var gzipped bytes.Buffer
		err = gzipWrite(&gzipped, data)
		if err != nil {
			return nil, err
		}

		//
		// Add the filename + data
		//
		tmp.Filename = file
		tmp.Contents = hex.EncodeToString(gzipped.Bytes())
		entries = append(entries, tmp)
	}
	return entries, nil
}

//
// Given the array of files render our output file.
//
func renderTemplate(entries []Resource) (string, error) {

	//
	// Create our template object
	//
	t := template.Must(template.New("tmpl").Parse(TEMPLATE))

	//
	// Execute into a temporary buffer.
	//
	buf := &bytes.Buffer{}
	err := t.Execute(buf, entries)

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

func main() {

	//
	// The command-line flags we support
	//
	input := flag.String("input", "data/", "The directory to read from.")
	output := flag.String("output", "static.go", "The output file to generate.")
	verbose := flag.Bool("verbose", false, "Should we be verbose.")

	//
	// Parse the flags
	//
	flag.Parse()

	if *verbose {
		fmt.Printf("Reading input directory %s\n", *input)
		fmt.Printf("Creating output file %s\n", *output)
	}

	//
	// Find the files.
	//
	files, err := findFiles(*input)
	if err != nil {
		panic(err)
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
	fmtCmd := exec.Command("gofmt")
	In, _ := fmtCmd.StdinPipe()
	Out, _ := fmtCmd.StdoutPipe()
	fmtCmd.Start()
	In.Write([]byte(tmpl))
	In.Close()
	Bytes, _ := ioutil.ReadAll(Out)
	fmtCmd.Wait()

	//
	// At this point we have "pretty" generated file.
	//
	ioutil.WriteFile(*output, Bytes, 0644)

}
