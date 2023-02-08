

package main

import (
	"bytes"
	"fmt"
	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter1/bytestrings"
)

func main() {
	in :=  bytes.NewReader([]byte ("example"))
	out := &bytes.Buffer{}

	fmt.Print("1. stdout on Copy = ")

	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Print("2. stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}

}