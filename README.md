# implant

This is a simple utility which allows data to be embedded directly in `golang`
applications.

The expected use-case is that you have a HTTP-server, or similar, which
needs to work with files, but because you're wanting to deploy a single
binary you do not wish to include this tree of support files with your
releases - Instead you wish to embed them into your application



## Usage

Assuming that all the files you wish to embed are located within a
single directory you would run:

      $ implant -input data/ [-output static.go]

This would generate the file `static.go` which contains the contents
of each file beneath the `data/` subdirectory.



## API

The generated file contains two functions to help you access your
embedded data:

* `getResources() []string`
    * Returns a list of all the files embedded in your binary.
* `getResource( path string  ) ([]byte, error)`
    * Return the content of a single embedded file.

So you could retrieve the content of the file `data/index.html` via
this call:

      contents, err := getResource( "data/index.html" )


## Data Storage

To save space the embeded file contents are compressed with `gzip`, albeit
encoded via `encoding/hex`.


Steve
--
