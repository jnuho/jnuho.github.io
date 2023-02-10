


package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// make use of Buffer() function
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"
	b := Buffer(rawString)

	// bytes.Buffer -> string
	fmt.Println(b.String())

	// bytes.Buffer -> bytes
	fmt.Println(b.Bytes())

	s, err := toString(b)
	if err != nil {
		return err
	}

	fmt.Println(s)

	// Take bytes to create kbytes reader
	// which implmenets io.Reader, io.ReadSeeker, ...
	reader := bytes.NewReader([]byte(rawString))

	// plug into a scanner that allows buffered reading tokenization
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// iterater over all of the scan events
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	return nil
}




