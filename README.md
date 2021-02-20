[![Go Report Card](https://goreportcard.com/badge/github.com/skx/implant)](https://goreportcard.com/report/github.com/skx/implant)
[![license](https://img.shields.io/github/license/skx/implant.svg)](https://github.com/skx/implant/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/implant.svg)](https://github.com/skx/implant/releases/latest)
[![gocover store](http://gocover.io/_badge/github.com/skx/implant)](http://gocover.io/github.com/skx/implant)

# implant

`implant` is a simple utility which allows data to be embedded directly
in [golang](https://golang.org/) applications (`implant` is a synonym of `embed`).

The expected use-case is that you have a HTTP-server, or similar `golang`
application which you wish to distribute as a single binary but which
needs some template files, HTML files, or other media.

Rather than distributing your binary along with a collection of files you
instead embed the file contents in your application, and extract/use them
at run-time.

implant allows you to do this, by generating a file `static.go` which
contains the contents of a particular directory hierarchy.  At run-time you
can list available files, and extract the contents of specific ones.


## Obsolete

go v1.16 makes this tool obsolete, via the addition of the `go embed` syntax
for embedding resources inside generated binaries.

Further details can be found in the release-notes:

* https://golang.org/doc/go1.16

As a result of this update this repository has been marked read-only on 20/02/2021 and no further development work will be carried out.


## Installation

There are two ways to install this project from source, which depend on the version of the [go](https://golang.org/) version you're using.

### Source Installation go <=  1.11

If you're using `go` before 1.11 then the following command should fetch/update the project and install it upon your system:

     $ go get -u github.com/skx/implant

### Source installation go  >= 1.12

If you're using a more recent version of `go` (which is _highly_ recommended), you need to clone to a directory which is not present upon your `GOPATH`:

    git clone https://github.com/skx/implant
    cd implant
    go install


If you prefer you can fetch a binary release from our [releases page](https://github.com/skx/implant/releases).



## Usage

Assuming that all the files you wish to embed are located within a
single directory you would run:

      $ implant -input data/ [-output static.go]

This would generate the file `static.go` which contains the contents
of each file located beneath the `data/` directory.

Further options are available to change the package-name, increase
verbosity and disable the automatic formatting of the code via `gofmt`.
Please consult the output of `implant -help` for details.


## Runtime API

The generated file contains two functions to help you access your embedded data:

* `getResource( path string  ) ([]byte, error)`
    * Return the content of a single embedded file.
* `getResources() []EmbeddedResource`
    * Returns a list of all the embedded resources in the your binary.
    * The returned structure contains the original name of the file, the size, and the contents in encoded form.
    * It is assumed you'll use `getResource` to get the original data, rather than trying to decode & decompress manually.

Sample usage of retrieving the content of the file `data/index.html` would look like:

      contents, err := getResource( "data/index.html" )


## Data Storage

To save space the embedded file contents are compressed with `gzip`, however
this results in binary-data so the compressed bytes are encoded via
`encoding/hex`.  The use of the hex-representation does mean that some of
the compression is effectively wasted, but it is a reasonable trade-off.


## Users

The application uses itself to embed the template file which is written
to `static.go`.  In addition to that several of my small applications use
this library, for example:

* https://github.com/skx/dns-api-go
* https://github.com/skx/puppet-summary
* https://github.com/skx/purppura

You'll see in each case I've committed the generated `static.go` file to the repository, which means users who don't need to change any of the resources aren't forced to install the tool.

I suggest a similar approach in your own deployments.


## Testing

The repository contains a number of test-cases, they can can be executed via:

    $ go test ./...
    ok  	github.com/skx/implant	(cached)
    ok  	github.com/skx/implant/finder	(cached)

To receive coverage details:

    $ go test -coverprofile=tmp.t ./... && go tool cover -html=tmp.t && rm ./tmp.t

## Github Setup

This repository is configured to run tests upon every commit, and when
pull-requests are created/updated.  The testing is carried out via
[.github/run-tests.sh](.github/run-tests.sh) which is used by the
[github-action-tester](https://github.com/skx/github-action-tester) action.

Releases are automated in a similar fashion via [.github/build](.github/build),
and the [github-action-publish-binaries](https://github.com/skx/github-action-publish-binaries) action.

Steve
--
