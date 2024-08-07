package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)


var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)

	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}


func main() {
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	cnt := 1

	for {
		fmt.Printf("숫자값을 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("[ERROR!] 숫자만 입력하세요.")
		} else {
			if n > r {
				fmt.Println("입력하신 숫자가 더 큽니다")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 적습니다")
			} else {
				fmt.Printf("입력하신 숫자는 %d 입니다.\n", n)
				break
			}
			cnt ++
		}
	}
}