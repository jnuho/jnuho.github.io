package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i:=a; i<=b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계= %d\n",a,b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)	// 총 작업개수 설정

	for i :=0; i<10; i++ {
		go SumAtoB(1, 1000000000)
	}

	wg.Wait()
	fmt.Println("모든 계산이 완료 되었습니다.")
}