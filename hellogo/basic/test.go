package main

import (
	"runtime"
	"sync"
)

const initialValue = -500

type counter struct {
	i    int64
	mu   sync.Mutex
	once sync.Once
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}          // 카운터 생성
	done := make(chan struct{}) // 완료 신호 수신용 채널
}
