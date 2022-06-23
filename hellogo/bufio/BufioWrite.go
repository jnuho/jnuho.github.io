package main

import (
	"fmt"
	"bufio"
)

type Writer int


// func (b *Writer) Write(p []byte) (nn int, err error)
// https://pkg.go.dev/bufio#Writer.Write
// Write writes the contents of p into the buffer. It returns the number of bytes written. If nn < len(p), it also returns an error explaining why the write is short.
func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("Writing: %s\n", p)
	return len(p), nil
}


func main() {

	// func NewWriterSize(w io.Writer, size int) *Writer
	// https://pkg.go.dev/bufio#NewWriterSize
	// NewWriterSize returns a new Writer whose buffer has at least the specified size.
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 4)

	// Case 1: Writing to buffer until full
	// rune is a character
	// In runes we use single-quotes ''
	// rune is also an alias for int32
	bw.Write([]byte{'1'})
	bw.Write([]byte{'2'})
	bw.Write([]byte{'3'})
	bw.Write([]byte{'4'})	// write - buffer  is full
	// buffer = 0 after write!
	fmt.Printf("Available buffer size : %d\n", bw.Available())

	// Case 2
	bw.Write([]byte{'5'})
	// buffer = 3!
	fmt.Printf("Available buffer size : %d\n", bw.Available())
	err := bw.Flush() // forcefully Write remaining
	if err != nil {
		panic(err)
	}

	// buffer = 4! after flush
	fmt.Printf("Available buffer size : %d\n", bw.Available())

	// Case 3
	bw.Write([]byte("12345"))


}