package interfaces

import (
	"fmt"
	"io"
	"os"
)

func Copy(in *io.ReedSeeker, out io.Writer) error {
	// write to out and 
	w := io.MultiWriter(out, os.Stdout)

	// standard copy
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0, 0)

	// standard copy
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	fmt.Println()

	return nil
}