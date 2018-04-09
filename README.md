[![Travis CI](https://img.shields.io/travis/skx/implant/master.svg?style=flat-square)](https://travis-ci.org/skx/implant)
[![Go Report Card](https://goreportcard.com/badge/github.com/skx/implant)](https://goreportcard.com/report/github.com/skx/implant)
[![license](https://img.shields.io/github/license/skx/implant.svg)](https://github.com/skx/implant/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/implant.svg)](https://github.com/skx/implant/releases/latest)

# implant

`implant` is a simple utility which allows data to be embedded directly
in `golang` applications (`implant` is a synonym of `embed`).

The expected use-case is that you have a HTTP-server, or similar `golang`
application which you wish to distribute as a single binary but which
needs some template files, HTML files, or other media.

Rather than distributing your binary along with a collection of files you
instead embed the file contents in your application, and extract/use them
at run-time.

implant allows you to do this, by generating a file `static.go` which
contains the contents of a directory hierarchy.  At run-time you can both
list available files, and extract the contents of specific ones.



## Usage

Assuming that all the files you wish to embed are located within a
single directory you would run:

      $ implant -input data/ [-output static.go]

This would generate the file `static.go` which contains the contents
of each file located beneath the `data/` directory.



## Runtime API

The generated file contains two functions to help you access your embedded data:

* `getResources() []string`
    * Returns a list of all the files embedded in your binary.
* `getResource( path string  ) ([]byte, error)`
    * Return the content of a single embedded file.

So you could retrieve the content of the file `data/index.html` via this call:

      contents, err := getResource( "data/index.html" )


## Data Storage

To save space the embeded file contents are compressed with `gzip`, however
this results in binary-data, so the compressed bytes are encoded via
`encoding/hex`.  The use of the hex-representation does mean that some of
the compression is effectively wasted, but it is a reasonable trade-off.


## Users

Three of my small applications use this library:

* https://github.com/skx/dns-api-go
* https://github.com/skx/puppet-summary
* https://github.com/skx/purppura

You'll see in each case I've committed the generated `static.go` file to the repository, which means users who don't need to change any of the resources aren't forced to install this tool.

I suggest a similar approach in your own deployments.


Steve
--
