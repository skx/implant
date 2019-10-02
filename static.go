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
	"os"
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
	ModTime  int64
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
		Contents: "H4sIAAAAAAAC/6RW32/bNhB+Fv+Kq4GhEmpIeyj6EMwPa5YNA7puSFrsIQgaWjrJB1OkQVJ2HEP/+3AkZQdxg3bbQwv6yPvuu+9+KFUlqgo+rchBSwphJx10qNFKjw1sSUJHfjUsy9r0lVs/VNRvlNSe3djzg6mlAmzIO9iRUrBEUMb5ku83sl7LDuFwKP+Kx3EUgvqNsR5yAQAwW+49ulk816bfWHSu6h5pk2yoa9OQ7qqldPjubbK2vU8nMhWZwZNKv42biUIkdlf9EpsGm2t0ZrA1AjnwKwTn7VD7wSLsVlSv2Dw4bMAbsFgb20CDXpJyYFrGQVmvABMY2COahr0ZLCxJS7svxSQmnp7URntJOobNjaWOtFRFEFvLHudgUUlPW+Tg/Ij0ZvCM05DF2hu7h/sk+v2z8uzIr+YgldFdOAf/KQY4ekTGkboJF5O62FRBVGygkV4G2n6/wXO1okxwCMr+mhizlXQXbJdGe9TePbV9QN35FQCQ9sHwh2k+UY/B8O6tGKfiXF/d/Pn5+vLqhuWX0MvNpBbpDqRS54q7MqjLT9e4T9VkrNC7zG4qwu+eb/FhY1Jda6kUWhdamtW479BPibq8uId20LUnowPCVton/BYc8TbmePdcpYPIDgcrdYdQHgHHUWSzw6GcRBvH2QUcRJZNhgt4dj0XWTbJGS+nX+ky6nrBwxSP48jmpG6wp3O44H+HA+pmHFlyzg469IH9ibzLN9KvUvkKyG/veB7ngNYaWzBj1N7u52DWcLE4SXLLbncioxZemXXIrKrgN25jCd7uWfGWdAMq7IepfiLLLPrBatCk5vyfyEYhMpbbyh2EZVC+H9oWbbSitZGMEBzhF+TGDQWMvZuRDmyZXdwQ5Y1vrtLSKOP7m5BeHlI5yloE9uz6asFMQhJP2aG1kR1nNuhH2hzjTrNaK0LtRdbZIwleXeVH3F2jbNDmMaOPuItJ5aSL7w2cNdiihc6Wl8o4zAuRcexjpLj2Sg70s1J5Z78b+EvEWLDm5d+WPOaM/K8UuY528qWY3jDae843L2Jtn8x6fBHXUNoZpgV57IwwdVOTTt35zeZMWCfxX2jwQmTfyq3tfXnFwG0+ayWpuDVCDx+X+esf3OvZHBLiyKDOS++O8Y0rb7z0U8zziFUFl2ZQDWjjI7hG8iu0YHSaFbd3HvtSZOyeEoTFCYIxPiYn0qcdefL7n3lxYhwkoZwt4ZNwR/nTILMXtdCaQTfMbWm4fIHXhMYdoHGHzsMWrSOjRdabxvMH4mIBQc5pj+VF+VnTQx6VnF799HwLTc8ngf8D75fmivfzy8X82mglWwRk+K+NgNxKUnKpTn8nOJZLglP0wii4vIDbs28P0yDm/GOIzKderjE/fzkHhTo/KsdLqDUWvsxhy17x+3X64qX0bukOFrDlXnzz5ml+Fr0YxT8BAAD//1QjRHo9CgAA",
		Length:   2621,
		ModTime:  1570058703,
	},
}

func getEmbededResources(path string) ([]byte, error) {
	entry, ok := RESOURCES[path]
	if !ok {
		// Give a try to find local resource
		return nil, nil
	}

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

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	content, err := getEmbededResources(path)

	if err != nil {
		return nil, err
	}

	stats, err := os.Stat(path)
	if err != nil {
		// Could not find neither on local system.
		if content == nil {
			// Neither in embedded system.
			return nil, fmt.Errorf("failed to find resource '%s'", path)
		}
		// return embedded resource
		return content, nil
	}
	// if found in both system return the newest version
	modtime := stats.ModTime().Unix()
	if modtime < RESOURCES[path].ModTime {
		// return embedded resource
		return content, nil
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
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
