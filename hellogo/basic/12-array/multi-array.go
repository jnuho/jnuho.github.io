package main

import (
  "fmt"
)

func main() {
  a := [2][5] int {
    {1,2,3,4,5},
    {6,7,8,9,10},
    // 행바뀜 있을때는 마지막에 , 필요
    // 행바뀜없는 inline 일때는 필요 없음 e.g. { {...},{.., 9,10} }
  }

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, "")
		}
		fmt.Println()
	}
}