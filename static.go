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
		Contents: "H4sIAAAAAAAC/5SUUW/bNhDHn8VPcTMwVEINaQ9FHzLkYc2yYUDXDUmHPRhBQ0sn+WCKFEjKriPouw9HUo6RrA99MEAfdf+7+93xqkpUFXzekYOWFMJROuhQo5UeGziQhI78btyWtekrt/9aUT8oqT27sedHU0sF2JB3cCSlYIugjPMl3w+y3ssOYZrKv+NxnoWgfjDWQy4AAFbbk0e3iufa9INF56ruiYZkQ12bhnRXbaXD9++Ste19OpGpyIye1EoUIiV122+xabC5Q2dGWyOQA79DcN6OtR8twnFH9Y7No8MGvAGLtbENNOglKQemZR2U9Q4wiYE9q2k4mdHClrS0p1IsDPH5k9poL0nHsLmx1JGWqgiMtexxDRaV9HRADs4fkR5GzzoNWay9sSd4TKwfX3TlSH63BqmM7sI5+C8xwNETso7UTbhYoGJTBZbYQCO9DGn704CvaUVMMAXAv6WM2Uq6C7Yboz1q7y5tH1F3fgcApH0w/Gmaz9RjMLx/J+alOXe393/9c3dze8/4JfRyWGiR7kAq9Zq4KwNd/nSPp9RN1gojy9ktTfjD8y1+HUzqay2VQuvCJDONxw79UqjLi0doR117MjooHKS9yO+aI25ijQ8vKU0imyYrdYdQngXnWWSraSoXaPO8uoJJZNliuIIX12uRZQvOeLn8S5eR6xW/oXicZzYnusGezuGCf9OEupnnC+ToR6vTNKTWmRbkGXAonknABZ4cBul3qccABeSbB36ta0BrjS24MGoBtbenNZg9XF0/w9uw78PPbOb6GayVRwivvfwwti3aZEZro6AQWVZV8CvykIZk45xmGekQkwPELVDe++Y2LYYyOtyHNPOQzZlhwc5t8P3hGjSpkExmIw9NKuiKLJtT8N9H/UTDOfjyOGtFqL3Iss6eM+EdVX7C4x3KBm0eK/uEx1hcTrpI0fPL+MW3M8iyBlu00NnyRhmHOftzFueQcdGVHPEXpfLOfkd9WfYl6lxzI8p/LXnMWf07s4yc0kSR5/akz1j2A1PIizX7iGwOi+BSpe19ecvNbvNVK0nFR9qSvlivb350b1brMHzF/86wPEhScque963jnSzBKfrGLLu8gM2rNxwmmMH+JLgIPvVyj/nrL9egUOfn6ebWtsbClzUc2CvugefNMUUqG3qAazgw4LdvGccZFXoxi/8CAAD//0G30x18BwAA",
		Length:   1916,
		ModTime:  1570055768,
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
