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
	tmp.Contents = "1f8b08000000000004ffb4554d8fe3360c3ddbbf820d50ac0d04d6bd400edd695a1458ec2e325dec61b1e828366d13ab4806254f2663f8bf17943f269dc9a1970a48a03014f9f81e2929952a057fb5e4a1268370d61e1ab4c83a60058fa4a1a1d0f6c7a27427e57f3c293a7546db20c7e4e407576a035851f0702663e088609c0f45aa54dae9f2876e1086a1f83c6dc7314de9d4390e90a500009be325a0df4cfbd29d3a46ef55f34cdd6c435bba8a6ca35a7c9a4df529cc3b728a5c1fc86cd23c958c4ac1fe74c4aac2ea80def55c229087d022f8c07d197a4638b754b6401e7a8f1504078ca5e30a2a0c9a8c07574b1cd4650b3807035ea359b8b89ee14856f325961909c41797d2d9a0c94e6933c7d490d5268f045b7dc22d301a1de81121b8888d6cd70789531163191c5fe08126a21f5e4972a6d06e411b679bb88fe7971ce0e919258eb655fc6361142b1589c40a2a1d74841d2e1dbe656ba20986a8c8ef642262b1926da2edced98036f86bdb07b44d6801806c48c74589c3fefed397c3ddfe5e24d0e0a52204cdac2f0b47641b7016016de00bd48eff4dbb94b2301f312b057f0609874f9d9bd52bb531c83e36ab08fdd06058c4f759fe00756fcb40cec6088f9aaf807dfbfeba5d16f09f5dd71b1d1044ec0583971e937040964296c390464e246838756fd89cfe1d06d6b641285654e39826e1d4152bbf3bd80cc3fa731ca7fe96256e2be593dbf2f3b5dbacc20e8661de8f638427eb458c1de8ae435b652fb6ada4813c3a0f03da6a1caf64c4d0b38dd44a6747ed5d0d7a6525f21a59b9623e834e87766e12801cb26fdf65d6b780cc8e853a812592ffbd9df5ff650791a92b64208e8952f281af08b5ebe7d68ed7d502a888c16429b56ea9866c8abc320bbbdd842b864d12118ef51904992fdef7758dbc9e7fbdc41999e5e3384d9324510a7ec3d25518d999262b4912b2b14aa9a7c5a76272b90f324259ecf455c23cbad712127eda8125038398129e48b7646228b18d4bca3f7afb4cdd9a72b9444a43688378361c0f497eb9498b8f783ea0ae90b3a9ce8f789e4acdc8e693eab7d6c41fafc866cede60bb7556d628582aac91a1e1e2ce388f592c58885a114e177821007f35266b38bafc674e922491fe6106699e73f19529602619fecfc224ad52308f0685420cb366accfc57b79d5b27c2b02aefc4cc3387dcfbe51dffa148abd0c459d6d6a4d06e3a35493bd7a76defdecdf6db6b177f39ba3a91f35197d342fef90bf39973ecb6fdc7af37d3fa33aecef3f7d39dcedefd331fd270000ffff4909b64a21080000"
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
