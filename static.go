//
// This file was generated via github.com/skx/implant/
//
// Local edits will be lost.
//
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

//
// EmbeddedResource is the structure which is used to record details of
// each embedded resource in your binary.
//
// The resource contains the (original) filename, relative to the input
// directory `implant` was generated with, along with the original size
// and the compressed/encoded data.
//
type EmbeddedResource struct {
	Filename string
	Contents string
	Length   int
}

//
// RESOURCES is a map containing all embedded resources. The map key is the
// file name.
//
// It is exposed to callers via the `getResources()` function.
//
var RESOURCES = map[string]EmbeddedResource{

	"data/static.tmpl": {
		Filename: "data/static.tmpl",
		Contents: "H4sIAAAAAAAC/5SUUW/jNgzHn+VPwQUYzsYZ9h6GPXTow67rhgGH29Bu2ENRXBWbdojIkiHJyaWGv/tASU6DdvdwDwEUyiT//JFUXWd1DX/vyEFHCuEoHfSo0UqPLRxIQk9+N22rxgy123+paRiV1J7d2POjaaQCbMk7OJJSsEVQxvmK70fZ7GWPMM/VX/G4LFlGw2ishzwDANhsTx7dJp4bM4wWnav7ZxqTDXVjWtJ9vZUOf/oxWbvBpxOZmszkSW2yIkuiboctti22d+jMZBsEcuB3CM7bqfGTRTjuqNmxeXLYgjdgsTG2hRa9JOXAdBwHZbMDTMHAnqNpOJnJwpa0tKcqWxniyyeN0V6SjmlzY6knLVURGGs5YAkWlfR0QE7OH5EeJ89xWrLYeGNP8JRYP73qypH8rgSpjO7DOfivOcDRM3IcqdtwsULFtg4ssYVWehlk+9OIb2lFTDAHwL8lxWwl3QfbjdEetXeXto+oe78DANI+W9ZO3N3e//nP3c3tPbOWMMhxRUO6B6nUW7yuCij50z2eUus4VphPlrIS/8PzLX4ZTWpiI5VC68LYculPPfq1KpcXT9BNuvFkdIhwkPZC3zVnfIgFPb5GMmdinq3UPUJ1DrgsmdjMc7USWpbNFcyZEKvhCl5dl5kQK7t4uf5LlxHiFS9MPC5LmQn+zTPqdlkuyKKfrE4dTu0wHcgzx1AjFwwXFHIYpd+lvgEUkD888gaWgNYaW7B+6gC1t6cSzB6url8YPbDv489s5jKZn5VHCBtcfZi6Dm0yo7UxYJYJUdfwK/LgBbFx9oQgHXJygrjZ1b1vb9OyV9HhPsjMg5ozqoKdu+D73TVoUkGMsJGHJhXiZkIsKfnvk36m8Zx8XbhGEWqfCdHbsxJ+d6pPeLxD2aLNY2Wf8BiLy0kXKXt+mb/4ugIhWuzQQm+rG2Uc5uzPKs4p4+NVccZflMp7+w31CfE5xrnmRlT/WvKYc/RvVBk5pYkiz+1Jn3HYD0whL0r2ycQStv0ySjf46pab3eWbTpKKu9iRvngy333v3m3KMHzF/86wPEhScqte3lDH76wEp+grs+zyAh7erGqYYAb7Q8ZF8GmQe8zfflmCQp2fp5tb2xkLn0s4sFdc95cHYo5UHugRruHAgN+/ZxxnVMjv3n8BAAD//7+/WtJQBwAA",
		Length:   1872,
	},
}

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	if entry, ok := RESOURCES[path]; ok {
		var raw bytes.Buffer
		var err error

		// Decode the data.
		in, err := base64.StdEncoding.DecodeString(entry.Contents)
		if err != nil {
			return nil, err
		}

		// Gunzip the data to the client
		gr, err := gzip.NewReader(bytes.NewBuffer(in))
		if err != nil {
			return nil, err
		}
		defer gr.Close()
		data, err := ioutil.ReadAll(gr)
		if err != nil {
			return nil, err
		}
		_, err = raw.Write(data)
		if err != nil {
			return nil, err
		}

		// Return it.
		return raw.Bytes(), nil
	}
	return nil, fmt.Errorf("failed to find resource '%s'", path)
}

//
// Return the available resources in a slice.
//
func getResources() []EmbeddedResource {
	i := 0
	ret := make([]EmbeddedResource, len(RESOURCES))
	for _, v := range RESOURCES {
		ret[i] = v
		i++
	}
	return ret
}
