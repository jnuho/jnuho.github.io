package main

import (
	"fmt"
)

var arr []int = make([]int, 2)

func fibonacci(n int) int {
	if n == 0 {
				arr[0] = arr[0] + 1
				fmt.Printf("0")
        return 0
    } else if n == 1 {
				arr[1] = arr[1] + 1
				fmt.Printf("1")
        return 1
    } else {
        return fibonacci(n-1) + fibonacci(n-2)
    }
}

/** 1은 2번 출력되고, 0은 1번 출력된다.
N이 주어졌을 때, fibonacci(N)을 호출했을 때, 0과 1이 각각 몇 번 출력되는지 구하는 프로그램을 작성하시오.
입력
첫째 줄에 테스트 케이스의 개수 T가 주어진다.
각 테스트 케이스는 한 줄로 이루어져 있고, N이 주어진다. N은 40보다 작거나 같은 자연수 또는 0이다.
- 입력
T
N1
N2
...
NT

- 출력
r1 r2
...
r1 r2
*/

// ./test T
func main() {
	T,N := 0, 0
	arr := make([]int, 0)
	cnt := 0

	fmt.Scanln(&T)

	for i in range T {
		fmt.Scanln(&N)
		arr = append(arr, N)
	}

	for _, n in range arr {
		cnt = 0
		fibonacci(n)
		fmt.Println(cnt)
	}

}
