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
	tmp.Contents = "1f8b08000000000004ff8c54c16ee336103d4b5f31f549020cf15ec0876e9aeda5d82e9c167b088282164732b1142990a3385941ff5e0c29d1e9c68710b0410d8733efbd99a110a510f0f75907e8b441b8c8003d5af49250c1b396d06b3a4fa7a6758308df5f841e46232df135bef9a76ba501549a025cb4317042302e50530a518eb2fd2e7b84796ebea6edb294a51e46e709aa120060777a250cbbb46fdd307a0c41f43ff4b8dad0b64e69db8b33beac26ed84761369b37ea3f7ce875d59979c34d24150d869ab493b0bae033779f018dce45b0cec45af23c2fd7042a5501dd71308e4a796608e703e6b83560ed1aa6d1f6d77ce125a0aecc9b6e56d4aa303dd4cf62c3d1cef1ffefae77877ff008f4f3fe7dd827c75e36424e17bb8dd645b603e550d7319a170501ac67724d2e93c7b697b8466e31696a52c68189b4ceb00bb79ce9fcb92d4e4c56e996972db3edfba5d291d408e235a555d6d7b8e0275843acf68d5b25cd53a224dde029d11da4d51d781cc358aed1349f7481b850a4649e7557a801aaac7276e9f3dc40e6065187de73cfcbb07b4e45fe1d7034421de2003762c84e01f7c43e8dc6455c412276003d4c460bc84c85bdd41952267e1e07048b862d8a2e0ba78790146169a4f53d7a12fcba2288480dfb1750a632e2549366cd636e267a4677c6992cb03717b559143d6be8eee1db3855f0e60b581994d854f725a6d6228b62d5bca3f26fb438f3925908bfbd668b4c49ebd8f97383f8f5df3052f47940a7d95187cc14b2251695b47080a3bf4d0fbe6ceb88055b24992394e9ace86c3fc664cd5fbe8f261e4455170fdbc072edea5f9e63561c519523bdd5aa9303e0bb316e39d34b7eef2ca7aadada929166755d6cb4bf3891faaaadeb3cc39cc1277e97ff5ddaae07c60ed2ad87d96daa062e53b6d55eef15d7d7320f8c909711a9ea536f26430df0837c72254353c3e85d831ebd3157b10c36428e4a3f42e7c683858105e5b88eb78af9675baf20ca4bafc4f068f613214caa5fc2f0000ffff20504ad663060000"
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
