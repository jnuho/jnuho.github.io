package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
	str1 := "Hello World"
	str2 := str1


	// string -> unsafe.Pointer -> *reflect.StringHeader
	stringHeader1 := (*reflect.StringHeader) (unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader) (unsafe.Pointer(&str2))
	
	// 두 값이 같음
	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}