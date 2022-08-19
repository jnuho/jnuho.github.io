package main

import (
  "fmt"
)

func main() {
  var a int = 500
  var p *int

  p = &a

  fmt.Printf("p의 값: %p\n", p) // 메모리 주솟값
  fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p) // p가 가리키는 메모리의 값 *p == a

  *p = 100 // p가 가리키는 메모리의 값 변경
  fmt.Printf("===p가 가리키는 메모리의 값 변경->100===\n")
  fmt.Printf("a의 값: %d\n", a) // a값 변화 확인 *p == a


  var p2 *int = &a
  fmt.Printf("p2도 a를 가리키는 *int 포인터 입니다: \n\tp2= %p, *p2=%d, &a=%p\n", p2,*p2, &a)
}