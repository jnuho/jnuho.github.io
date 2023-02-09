package interfaces

import (
	"fmt"
	"io"
	"os"
)


// Copies data from in to out directly
// , writes to stdout using buffer
func Copy(in io.ReadSeeker, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout)

	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	io.Seek(0, 0)

	if _, err := io.Copy(w, in, buf); err != nil {
		return err
	}

	fmt.Println()

	return nil
}