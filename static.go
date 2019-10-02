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
// RESOURCES is a simple array containing one entry for each embedded
// resource.
//
// It is exposed to callers via the `getResources()` function.
//
var RESOURCES []EmbeddedResource

//
// Populate our resources
//
func init() {

	var tmp EmbeddedResource

	tmp.Filename = "data/static.tmpl"
	tmp.Contents = "1f8b08000000000002ffb455516fdb460c7e967e056760a80418d2fb003fac5d360c28ba22d9d087a268ce1225113ddf0977541cc7d07f1f787792dd3479d8c30c2438f3c88f1f3f92e7baceeb1afe1ec843471ae1a83cf468d029c6161e48414f3c4cfbaab187da7f7bace9306a6558c224f2bd6d94066c893d1c496bd82368ebb992fb5135df548f703e571fe3719ef39c0ea3750c450e00b0d99f18fd269e1b7b181d7a5ff74f34261b9ac6b664fa7ac0c764ea0e9c4e646bb21393dee4659e18dd1cf6d8b6d8dea2b7936b10c8030f089eddd4f0e4108e03358398278f2db005878d752db4c88ab407db090eaa66004c60e0563403273b39d89351ee54e58b80787169ac614526a62daca39e8cd26510d8a8036ec1a1564c0f28c9c589cc38b1e0b4e4b061eb4e709f84be7fd69223f1b005a5ade9c339c42f39c0d3130a8e326db85814c5b60e42620bad621568f369c41fd58a32c13908fc7b622c56327db0bdb386d1b0bfb6bd47d3f3000064389f974edcdedcfdf5cfedbb9b3bd15a81978a109473eab46844a6076b10d0b03b4167ddf7b20bca22eb22f59f2c70f838dad4bd46698dce8761959aef7be4a51c5f94f7d04da661b226203c287745ecf397e7022ce43fda71d28a11a4d90b072f97020764888b12ce79a85f40f930c28f60727b3e3b657a846a6535cf79c687b15af5ddc1e67c5ebfce739c6ff988db2a79745bbe3e774b5dd8c9bec5f33caf0e979a77a0c6114d5b5c6cdbc0be4c6cd1b4f37cd546e4c999344e8988ed407ddf99a0ca95f2058c8a8734240025149fbfc8ae6f019db3ae4c13262dffba4dfdff650741a92b66208e595dcb1f7c42e8ec94463b3c570ba16aadb3aed723755044e42ba177915780cd32699c534708af50f576ea3a746bfcf38f38a373b1805cc2eb1a7e4359abc0286e56966564429552cf808f5574b90b4a1491cfd2c232b877c1fba71d18d29157e6a2e886748012dbbca4fc63324f34ae299747a4d18486c5a7776b7e7949ab0f78bc45d5a22b629d1ff0184b2dc894e5abf546fd2eccca57b8bd163f8b778b1d3ae85df54e5b8f45285858af0ce3035e09c15fb52e7af7df34c9b2ec6bc492e139569f1c311692e1ff2c2c7522ad0671687b0a17166f45e7a2dc0a567e09bbfcbf4ed51db8ba9199ea8a4da748c767ad2373f5b3f3e667ff66b30db35bbeb89aea4191567b7df91df22feea52fca175ebdb48d89d5ba7df99cff1b0000ffff4909b64a21080000"
	tmp.Length = 2081
	RESOURCES = append(RESOURCES, tmp)

}

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	for _, entry := range RESOURCES {
		//
		// We found the file contents.
		//
		if entry.Filename == path {
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
	}
	return nil, fmt.Errorf("failed to find resource '%s'", path)
}

//
// Return the available resources.
//
func getResources() []EmbeddedResource {
	return RESOURCES
}
