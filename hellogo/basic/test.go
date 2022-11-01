package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}

func main() {
	wg.Add(1)

	// 5초 후 컨텍스트의 Done()채널에 시그널을 보내 작업종료 요청
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go PrintEverySecond(ctx)

	wg.Wait()
}
