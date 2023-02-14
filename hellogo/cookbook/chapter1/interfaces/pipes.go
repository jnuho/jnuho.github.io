package interfaces

import (
	"io"
	"os"
	"fmt"
)

func PipeExample() error {
	// implements io.Reader and io.Writer
	// Write to the PipeWriter blocks until it has satisfied one or more Reads from PipeReader
	// the Data is copied directly from the Write to the corresponding Read
	// Write -> Read pipe
	// 	if Goroutine is blocked on a channel Write operation
	// 	it cannot proceed until another Goroutine performs corresponding read operation
	// 	with the separate Goroutine, even if either one of the Reader or Writer Goroutine is blocked,
	// 	one of them will be able to proceed Read or Write operation
	// 	this prevents deadlocks and allows two operations proceed concurrently
	r, w := io.Pipe()

	// separate go routine as it will block
	// waiting for the reader close at the end for cleanup
	go func() {
		w.Write([]byte{"test\n"})
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}


