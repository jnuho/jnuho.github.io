package readwriter

import "fmt"

// 1.
// 인터페이스를 포함하는 인터페이스

type Reader interface {
	Read(n int, err error)
	Close() error
}

type Writer interface {
	Write(n int, err error)
	CLose() error
}


type ReadWriter interface {
	Reader
	Writer
}
