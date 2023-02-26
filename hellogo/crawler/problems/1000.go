package main

import (
	"fmt"
)
/**
https://www.acmicpc.net/problem/1000

두 정수 A와 B를 입력받은 다음, A+B를 출력하는 프로그램을 작성하시오.첫째 줄에 A와 B가 주어진다. (0 < A, B < 10)첫째 줄에 A+B를 출력한다.
[INPUT 1]
1 2
[OUTPUT 1]
3
[INPUT 2]

[OUTPUT 2]

[INPUT 3]

[OUTPUT 3]

[INPUT 4]

[OUTPUT 4]

*/
func main() {
	var A,B int
	fmt.Scan(&A, &B)
	fmt.Println(A+B)
}