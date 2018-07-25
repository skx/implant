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
	tmp.Contents = "1f8b08000000000004ffb455416fdc380f3ddbbf82df9c6c6060dd0bcce16b365d2c50b4c5648b1e8a62a3b1699ba84632283993a9e1ffbea0643bd92487bdac8019c834453ebe47ca4ae54ac19f3d7968c9205cb4870e2db20ed8c00369e828f4e3a9aadd59f99f8f8ace83d136c83139f9d1d5da0036143c5cc818382118e743952b950fbafea93b8469aabea4ed3ce7399d07c7018a1c006077ba06f4bbb4afdd7960f45e75bf68586c686bd790ed548f8f8b899c22370632cb33323bf6bbbccc25a952707b3e61d3607344ef46ae11c843e8117ce0b10e23235c7aaa7b200fa3c7068203c6da71030d064dc6836b250eeaba075c82016fd12c5cddc87022abf91a2b8d1ce2934bed6cd06453dac2317564b52923c7569f710f8c46077a40082e62233b8c41e234c45807c757b8a7c4f5fd0b552e14fa3d68e36c17f7f1fc9a033cfd4289a36d135faca462a32297d840a3838eb0c375c0d76c259a608aa27c2013118b956c176d37ce06b4c13fb77d44db851e00c8867c5e9538dede7dfe7abcb9bd13093478a9084133ebebca11d90e9c45401bf80aade37fd22ea5accc47cc4ac11f41c2e1e3e016f56a6d0cb28ffd2a42df771856f17d51de433bda3a90b331c283e667c0beff78d92e2bf82f6e188d0e0822f68ac14b8f4938204ba12861ca232712349c87576ca6b7d3c4da7608d5866a9ef32c9c876ae3f700bb69da1ee73935b72c71db284f6eebe34bb74585034cd3b29fe7084fd6931807d0c380b6299e6c7b490365749e26b4cd3c3f9311c3c836522b9d1db5772de88d95c86b64e519f3050c3af44b930094507cff21e3be8738b1429dc012c9ffda2ffabf3b4064ea193210c74c29f9c13784d68d4b6bc71b6b0554c560b294dab6d44291226fccc2e19070c5b05926c2b1be8020f3d5fbb16d91b7f32f973823b3fc1ce77996654ac16f58bb06233b69b2b22c231bab8403f4f858258fbb201354c446df142ca3772b11e17f07b064601253c689734b264612dbbc66fc7db4bf68d832ae77486d086d10cf8ee32138805ca5d527bc1c5137c845aaf2135e52a105d93269fed64aecf1066c61ec15b4b7ceca9a054a832d32745cdd18e7b188f50a4d09e0bb03a4bbbc1280ff37a6e838bafc6b4ab22c93ee612996f5a5fac614b0900cff6561925629580683422586453241f15e3e6b45b917fd367ed228a6ffc57795d7b117550ad87dd064307e935ab24f5f9d5df9e638ea074d469fccd3b7c7bf398bbe28dfb8e9963b7ec172bcbdfbfcf578737b97cff9df010000ffffff022b6f18080000"
	tmp.Length = 2072
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
