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
	tmp.Contents = "1f8b08000000000004ffb455c18edc360c3ddb5fc1cec90606d6bdc01c9aeda6281024c16c831c82a0abb1699b8846322879672786ffbda0647bb7bb7be8a5026620d314f9f81e292b952b057ff5e4a1258370d11e3ab4c83a60030fa4a1a3d08fa7aa7667e57f3c2a3a0f46db20c7e4e407576b03d850f0702163e084609c0f55ae543ee8fa87ee10a6a9fa9cb6f39ce7741e1c0728720080dde91ad0efd2be76e781d17bd5fda461b1a1ad5d43b6533d3e2e26728adc18c82ccfc8ecd8eff23297a44ac1edf9844d83cd11bd1bb946200fa147f081c73a8c8c70e9a9ee813c8c1e1b080e186bc70d341834190fae9538a8eb1e700906bc45b3707523c389ace66bac3472884f2eb5b341934d690bc7d491d5a68c1c5b7dc63d301a1de80121b8888dec300689d310631d1c5fe19e12d7f72f54b950e8f7a08db35ddcc7f36b0ef0f413258eb64d7cb1928a8d8a5c62038d0e3ac20ed7015fb3956882298af29e4c442c56b25db4dd381bd006ffdcf6016d177a00201bf27955e2787bf7e9cbf1e6f64e24d0e0a52204cdacaf2b47643b7016016de02bb48eff4dbb94b2321f312b057f0609878f835bd4abb531c83ef6ab087ddf6158c5f745790fed68eb40cec6080f9a9f01fbf6fd65bbace03fbb61343a2088d82b062f3d26e1802c85a284298f9c48d0701e5eb199de4e136bdb21541baa79ceb3701eaa8ddf03eca6697b9ce7d4dcb2c46da33cb9ad8f2fdd16150e304dcb7e9e233c594f621c400f03daa678b2ed250d94d1799ad036f3fc4c460c23db48ad7476d4deb5a0375622af919567cc1730e8d02f4d025042f1edbb8cfb1ee2c40a75024b24ff7bbfe8ffeb012253cf908138664ac90fbe22b46e5c5a3bde582ba02a0693a5d4b6a5168a147963160e87842b86cd32118ef5050499afde8d6d8bbc9d7fb9c41999e5e738cfb32c530a7ec7da3518d94993956519d958a5d4d3e363955cee828c50113b7d93b08ceead84845f0e60c9c024a68c13e9964c0c25b6794df9c7687fd2b0a55c2f91da10da209e1dc743925f2ed3ea235e8ea81be422d5f9112fa9d4826c99547f6b25fe7843b670f60adb5b6765cd82a5c116193aae6e8cf358c48285a80d61bacd2b01f89b3145c7d1e53f73926599f40f3348f35caaaf4c010bc9f07f162669958265342854625834637da9dec987ad28f722e0c64f1ac6f4bff8aefa3af6a24a01bbf79a0cc6af524bf6e9bbb32bdf1c48fda0c9e89379fafaf837a7d117e51b77dd72cb2f588eb7779fbe1c6f6eeff239ff270000ffff1a7cc9011a080000"
	tmp.Length = 2074
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
	return nil, errors.New("Failed to find resource")
}

//
// Return the available resources.
//
func getResources() []EmbeddedResource {
	return RESOURCES
}
