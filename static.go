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
	"errors"
	"io/ioutil"
)

//
// The definition of our resources
//
type EmbeddedResource struct {
	Filename string
	Contents string
	Length   int
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

	tmp.Filename = "data/static.tmpl"
	tmp.Contents = "1f8b08000000000004ff8c54c18edb380c3ddb5fc1cdc90606d67d811cb6d9742f455b245bf450140bc5a61da2b26448f464a686ff7d41c9f6143339544002997e22df7ba4ac54ae14fc7ba5002d19849b0ed0a145af191b78240d1df175bc54b5eb55f8f1a4a81f8cb62cc7e4e407576b03d81007b891317041302e70952b950fbafea13b8469aa3ea7ed3ce739f583f30c450e00b0bb3c33865ddad7ae1f3c86a0ba9f342c31b4b56bc876ea8a4f4b889c22373299e519bd773eecf23297a2510e42832d596272165c0b6ef4e031b8d1d71804c5cf03c2b1bf60d360735ade40603fd60c53a4f39e0c5addc728d92ec60ece325a0e825c631fd0767c0500b29ccfbf723014f86ef547ede1743c7ffa723a1ccff0edfb6b226b92cf6e188d667ccbbf1d6d0d22b02861ca233749cafdf046557a3b4d5edb0ea15ac58679ce33ee876ad3b987dd346d8ff39cec9525b04d7a82ad8faf618b1b7b98a6653fcf919eac17cd7bd0c380b6295e620f5206ca089e26b4cd3cbfd879421ebd05be22d46b0f5c0b7aeb6a1cb8e84a87bc6a2c60d07c5d9a055042f1edbb0cdc03c49911eb8456eb3cfcf70068d93fc39f7b884efdc20c049829253ff88ad0bad136914bbc332ba12a2693a5d4b6a5168a94797316f6fbc42ba6cd32699cd7371066a17a37b62dfa3ccfb24c29f81b6bd760acd568d69584c946fec2f48a4f55829c5906b2881ab6e69411de8a5af8630f960c4c12ca7cb2d39289a92436af25ff19ed4f1ab692c02eee6b436859909d8f87a4be5cd4ea23de4ea81bf44552f0116f494441b68c141a6cd143e7ab8371018b14d3acb73ce93e5792e62f638ace47c86f33cfb24cfae73d48f36ed5574f8c855448e3746fa5c6f8cd98a5196facb97756d6e6d7329ac4b1398bb35edfaa77f2692bca07b1794b932e43fa5fb06b179c0fe25d01bbf79a0c36e27c4bb6d9667c57debd10fa5193d117831b30dcbd0da128ef7c6b96afddc2e5743c7ffa723a1ccff99cff1f0000ffff1024a2a51c060000"
	tmp.Length = 1564
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

			// Decode the data.
			in, err := hex.DecodeString(entry.Contents)
			if err != nil {
				return nil, err
			}

			// Gunzip the data to the client
			gr, err := gzip.NewReader(bytes.NewBuffer(in))
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
	return nil, errors.New("Failed to find resource")
}

//
// Return the available resources.
//
func getResources() []EmbeddedResource {
	return RESOURCES
}
