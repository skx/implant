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

Once you have executed the above command you'll have a generated file
which contains the following function to retrieve your contents at
run-time:

     func getResource( path string  ) ([]byte, error) {

So you could retrieve the content of the file `data/index.html` via
this call:

      contents, err := getResource( "data/index.html" )


## Data Storage

To save space the generated file embeds the data compressed via gzip,
albeit encoded via `encoding/hex`.


Steve
--
