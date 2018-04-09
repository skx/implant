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
	tmp.Contents = "1f8b08000000000004ff8c55416fdb38133d4bbf623e9f24c010ef1fe0c3369b2e1628dac2d9a287a2d8d0d2481a9426057214c711f4df17434a4a36c96109d8a046c3c7c7f78623a572a5e0af9e02b464102e3a408716bd666ce0813474c4fd78aa6a7756e1d7a3a2f360b46559262b3fb95a1bc08638c0858c8113827181ab5ca97cd0f52fdd214c53f5354de739cfe93c38cf50e40000bbd39531ecd2bc76e7c16308aa7ba26189a1ad5d43b6533d3e2e21728adcc8649667f4def9b0cbcb5c36550a6ecf276c1a6c8e18dce86b040ac03d42603fd63c7a844b4f750f14600cd8003bf0583bdf4083acc90470ade0a0ae7bc0050cfc8666e1ea460f27b2da5fe349a386f89c523bcb9a6cdab6709e3ab2da945163abcfb8078f46333d20b08bdcc80e230b4e431e6b76fe0af794b4be7fe5ca85b8df8336ce76711ed7af7b40a027141c6d9bf86215151b15b5c4061acd3ad2e6eb806fd54a32c1144df94826329628d92ec66e9c65b41c5ec63ea1edb80700b29ccfab13c7dbbb2fdf8e37b777628186202742d0deebebaa11d90e9c4540cbfe0aadf3ff965d8eb22a1f392b057fb2c0e1e3e016f76a6d0cfa10eb558cbeef9057f34351de433bda9ac9d988f0a0fd0b623f7ebe2e9795fc57378c46338298bd72085263020764898b12a63c6a22a07c1edea899de4e93d7b643a83656f39c677c1eaa4ddf03eca6697b9ce754dc32246d933ca5ad8fafd316170e304dcb7c9e233d19cf661c400f03daa6788eed651b2863f234a16de6f9858dc8a3b7515aa9ece8bd6b416faa445da32a2f942f60d0dc2f45025042f1e3a75cf73dc41b2bd2092db1fceffde2ffff0f10957ac10c2431534a7ef01da175e352dab163ad84aa082643a96d4a2d14097953160e87c42bc2669918e7f5058459a83e8c6d8b3ecfb32c530a7ec7da3518cf9dee4c966564237f61dae3639552ee582e47116b7833a78ce9ad9c16fe77004b062609653ec969c9442889cdeb967f8cf689866dcbb53dd486d0b264763e2e92fda54d569ff17244dda02fd2093ee3251da2205b460a0db6e8a1f3d58d71018b14d3ac379cd44d2b81f9cd98a2f331e53f33cfb24cfcf31ec4bc4bf5dd1363213ba4727a6f2463fc26cc62c61b69de5b2b63d36b294de24a782cca7a7da93ec887a528f722f306932e43fa5f7257179c0fa25d01bb8f9a0cc6af424bf6b9efefca772f847ed064f4c93c77fff0ee6d0845f94eaf59baecc2e5787bf7e5dbf1e6f62e9ff37f020000ffff06ea26a29a070000"
	tmp.Length = 1946
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
