package main

import (
	"fmt"
	"sync"
	"context"
)

var wg sync.WaitGroup

func square(ctx context.Context) {
	// 컨텍스트에서 값을 읽기
	// Value는 빈 인터페이스 타입이므로 (int)로 변환하여 n에 할당
	if v:= ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square:%d\n", n*n)
	}

	wg.Done()
}

func main() {
	wg.Add(1)

	// "number"를 키로 값을 9로 설정한 컨텍스트를 만듦
	// square의 인수로 넘겨서 값을 사용할 수 있도록 함
	ctx := context.WithValue(context.Background(), "number", 9)

	go square(ctx)

	wg.Wait()
}