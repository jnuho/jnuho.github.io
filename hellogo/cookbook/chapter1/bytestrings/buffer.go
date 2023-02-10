
package bytestrings

import (
	"bytes"
	"io"
	"io/util"
)

func Buffer(rawString string) *bytes.Buffer {

	// string -> raw bytes
	rawBytes := []byte(rawString)

	// Create Buffer from the raw bytes
	b := new(bytes.Buffer)
	b.Write(rawBytes) // 1 b.Write([]byte)
	b = bytes.NewBuffer(rawBytes) // 2 alternatively

	// Create Buffer from the raw string, avoiding the initial byte array altogether
	b = bytes.NewBufferString(rawString) // 2 alternatively

	return b
}

// Take io.Reader and consume it all, returning a string
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}


