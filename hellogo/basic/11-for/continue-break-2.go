package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		// 입력 된 값이 변수에 저장되기 전에 버퍼에 저장
		// 숫자는 number에 저장되고 
		// 숫자 + 
		fmt.Print("숫자를 입력하세요: [-1로 종료]")
		var number int
		_, err := fmt.Scanln(&number)

		if err != nil {
			fmt.Println("숫자가 아닙니다!")

			//키보드 버퍼 비움
			stdin.ReadString('\n')
			continue
		}

		fmt.Printf("입력 하신 숫자는 %d 입니다.", number)

		if number%2 == 0 {
			fmt.Printf("입력 하신 숫자는 짝수, %d, 입니다.\n", number)
		} else if number%2 == 1 {
			fmt.Printf("입력 하신 숫자는 홀수, %d, 입니다.\n", number)
		} else if number == -1 {
			fmt.Printf("for문 종료\n")
			break
		}
	}
}
