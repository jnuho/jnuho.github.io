package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	// 채널에 데이터가 들어 올때까지 '계속' 기다림
	// 데드락 방지: square() 호출 밖에서 close(채널)로 채널이 닫히면
	// for문을 종료하여 프로그램 정상 종료하도록 함
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}

}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	// 10번만 데이터를 넣음
	// square 내에서 채널 데이터 계속 기다림
	for i := 0; i < 5; i++ {
		ch <- i * 2
	}

	// 작업완료를 기다리지만,
	// square() 내에서 wg.Done()이 실행 되지 않고 deadlock발생
	// 하지만 채널을 닫아서 데드락 방지 가능
	// 채널에서 데이터를 모두 빼낸 상태이고, 채널이 닫혔으면
	// for range 문을 빠져나가게 됨
	quit <- true

	wg.Wait()
}
