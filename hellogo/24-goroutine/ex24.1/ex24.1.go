package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가','나','다','라','마','바', '사'}
	for _,v := range hanguls {
		time.Sleep(300*time.Millisecond)
		fmt.Printf("%c", v)
	}
}

func PrintNumbers() {
	for i:=0; i<5; i++ {
		time.Sleep(400*time.Millisecond)
		fmt.Printf("%d", i+1)
	}
}

func main() {

	// main() 고루틴 이외에
	// go 키워드로 고루틴 2개 추가 생성함
	// 각기 다른 새로운 고루틴에서 실행되기 때문에 동시에 실행 됨
	// 총 3개의 고루틴 실행됨
	// 만약 CPU 코어개수가 3개 이상이면 동시 실행,
	// 3개이상이 아니라면 동시에 실행 되는 것 처럼 보이지만 CPU1개가 두개를 동시에 작업 하게됨
	go PrintHangul()
	go PrintNumbers()

	// 위의 2개 고루틴이 끝날때 까지 기다림
	// 이 라인이 없거나 3초이내 Sleep이면, 2개 고루틴이 끝나기 전에 종료되어 2개 고루틴의 일부 결과가 미출력
	time.Sleep(3 * time.Second)
}