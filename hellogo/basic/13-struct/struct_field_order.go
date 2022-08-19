package main

import (
  "fmt"
  "unsafe"
)


type User struct {
  A int32 // 4바이트
  B float64 // 8바이트
}

type User2 struct {
  A int8    // 1바이트
  B int     // 8바이트
  C int8    // 1바이트
  D int     // 8바이트
  E int8    // 1바이트
}

type User3 struct {
  A int8    // 1바이트
  C int8    // 1바이트
  E int8    // 1바이트
  B int     // 8바이트
  D int     // 8바이트
}

func main() {
  user := User {34, 97.5}
  // 12바이트가 아닌 16바이트 출력 : 메모리정렬
  // 변수시작주소를 변구크기의 배수로 맞춤
  fmt.Println("user size : ", unsafe.Sizeof(user))

  user2 := User2{1,2,3,4,5}
  fmt.Println("user2 size : ", unsafe.Sizeof(user2))

  user3 := User3{1,2,3,4,5}
  fmt.Println("user3 size : ", unsafe.Sizeof(user3))

}