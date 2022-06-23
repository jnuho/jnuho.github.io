package main

import "fmt"

func main() {

	// 문자 하나를 표혀하는데 rune 타입 사용
	// UTF-8은 1~3바이트
	// 알파벳 문자열크기 =1, 한글 문자열 크기 = 3
	var c rune = '한'

	fmt.Printf("%T\n", c) // 타입 출력 : int32
	fmt.Println(c)				// int32(rune) 형식 값 출력
	fmt.Printf("%c\n", c) // 문자출력

	// 12 = (1*3) +  (3*3)
	a := "abc가나다"
	fmt.Println(len(a))


	// string <-> []rune 타입 변환
	// 		[]rune(str)
	// 		string([]rune{...})
	str := "Hello World!"
	// []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	runes := []rune(str)

	fmt.Println(str)
	fmt.Println(runes)
	fmt.Println(string(runes))
	fmt.Println(len(runes))

	for _,v := range runes {
		fmt.Printf("타입: %T, 값: %d, 문자열: %c\n", v,v,v)
	}

	// strng <-> []byte 변환
	// 모든 문자열은 1바이트 배열로 변환 가능
}