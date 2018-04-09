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
	tmp.Contents = "1f8b08000000000004ff8c54c16ee4360c3ddb5fc1cec90602eb5e600edd34db4bb12d262df6100485c6a23dc4ca9221d199640dff7b41c9e3499b39ac8019c814453ebe474aa95229f8eb44113ab208671da1478741331a78210d3df1693a36ad1f54fcf6aa6818ad762cd7e4e6efbed516d010473893b57044b03e72532a558ebafda67b8441932b4b1a461f18aa120060777c638cbbbc6ffd30068c51f5df69dc651bbad61b72bd3ae1eb6a22afc84f4c76fdc6107c88bbb22e25592a03c160478e98bc03df819f02048c7e0a2d46f1e2b711e16138a231680eeb09440e53cb30a7d49fc9a2d343b292eb93edde3b46c7513cc5b6bc4f6929f2cd642f3ac0e1e1f18fbf0ff70f8ff0f4fcffbc97207ffa71b29af123dc6e722d483d550d7399a048501ec60f45e4d37986a05d8fd0c0b294050f63b3d5b387dd3cc3f57b593291b2c4712b7275dcbedf3b5eebd9831e4774a6badaee240ed409e73ca333cb72a5ea803c05077c42682f74fa0ef42650ea9954718f7cd1a68251f369e51da086eae9597ae70e92fc428be0ef7c807fee001d8737f8799f4978870cc4b1504a7ef015a1f39333094b6afb0ba0260593a5d4b6a50eaa1cf9cadd7e9f71a5b04521a2047d0641169b4f53d76128cba22894825fb1f506532ea35937622697f00bd213be36d9e591a5b7aa54c3467e9ddc3ba9167eda83230bb3988a90e974645328b12d9794bf4dee3b8d5b4a609ff6ad25742c9e7d489724bfcc5cf305cf07d40643952bf882e75c4445ae4e100c7618a00fcdbdf511ab6cd3acb73879341b09f38bb5551f92cb0f232f8a42f40b0144bc73f33510632519723bdd5a5998b011b38af1819a5b77656d7cadad499cc459990dfadc7c7a638c557d27346f6196b4cbffabef45051fa27057c1eeb3268b4698efc899adc777f5cd8190f726a66978d164f5d1e27623de1c8b58d5f0f41c53c7acef56ea418c93e5b81de547e187864308917509711defd5b24ed736035997ffd010304e9663b994ff060000ffff174e899c58060000"
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
// Return the names of available resources.
//
func getResources() []string {
	var results []string

	for _, entry := range RESOURCES {
		results = append(results, entry.Filename)
	}
	return results
}
