package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("test.txt파일 생성 실패")
		return
	}

	// 실행 순서 4->5 -> 3->2->1
	defer fmt.Println("반드시 호출됩니다.") //1
	defer f.Close() //2
	defer fmt.Println("파일을 닫았습니다.") //3

	fmt.Println("파일에 'Hello World!' 를 씁니다!") // 4
	fmt.Fprintln(f, "Hello World!") // 5 파일에 씀
}