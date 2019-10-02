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
		Contents: "1f8b08000000000002ff9494516fe3360cc79fe54fc10518cec619f69e3be461d775c380c36d6837ec2128ae8a4ddb4464c990e4a4a9e1ef3e50b2d3a0dd3ddc430087a2f8277f2455964959c2df1d396848219ca48316355ae9b186234968c977e3bea84c5fbac37349fda0a4f67c8d6f7e3695548035790727520af608ca385ff0f920ab836c11a6a9f82b7ece7392503f18eb214d000036fbb347b789df95e9078bce95ed0b0d8b0d75656ad26dd9e1f3626a7abf7c9129c98c9ed426c99225a3bb7e8f758df53d3a33da0a811cf80ec1793b567eb408a78eaa8ecda3c31abc018b95b135d4e8252907a6e13828ab0e700906f6124dc3d98c16f6a4a53d17c90a105f5d2aa3bd241d655363a9252d5516006bd9630e1695f47444166727d2c3e8394e4d162b6fec199e16d04f6f5a7222dfe52095d16df80ef7570d70f4821c47ea3a1cac44b12e0348aca1965e86b4fd79c0f7b422269802e0df968cd94aba0db65ba33d6aefae6d9f51b7be0300d23e99d74edcdf3dfcf9cffdeddd03b396d0cb614543ba05a9d47bbcae0828d9f580e7a5751c2b0c27a7b212ffc3f3293e0f666962259542ebc2cc72e94f2dfab52a97664fd08cbaf26474887094f62abf2d2bee62418f6f914c8998262b758b505c02ce732236d354ac84e67973035322c46ab88137c77922c4ca2e1eaeff96c308f186b7257ece739e08fe4d13ea7a9eafc8a21fad5e3abcb4c334202f1c438d5c305c51486190be5bfa069041ba7be4f5cb01ad3536e3fca901d4de9e733007b8d9be32daf1ddc79fd9cc65323f2b4f10d6b7f834360ddac58cd6c68049224459c2afc88317928db32704e9a0c9021d3e17d1e321e49506f90b9b8cbd9be0fcc31634a9a02e6c04a04985408910f3a2f6fba85f68b8a8ad1b562942ed13215a7b91e657a6f882a77b9435da3496f2054fb19a9474b6a8a7d7fad9b73310a2c6062db4b6b855c661caf7398b8b647cad0a56fc45a9b4b5df519f105f639c2d932ffeb5e431e5e8df9965e4b48c1079eec7e2c6613f318534cbf94e22e6b0ded7519ade1777dcdd26dd3492545cbe86f4d51bf9e147f761938769cbfe7768e55192927bf5fa683a7e58253845df185e9766b07bb79b616419ec4f0917c15fbd3c60fade3307853abd8c33b7b63116bee670e45b71bf5f5f842952d9d1236ce1c8803f7e641c1754c80fdd7f010000ffff404485053e070000",
		Length:   1854,
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
		in, err := hex.DecodeString(entry.Contents)
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
