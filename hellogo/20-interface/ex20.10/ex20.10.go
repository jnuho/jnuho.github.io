package main

type Reader interface {
	Read()
}
type Closer interface {
	Close()
}

type File struct {

}

func (f *File) Read() {
}

func ReadFile(reader Reader) {
	// c, ok := reader.(Closer)
	// if ok {
	// 	c.close()
	// }
	if c, ok  := reader.(Closer); ok {
		c.Close()
	}

}

func main(){
	file := &File{}
	ReadFile(file)
}
