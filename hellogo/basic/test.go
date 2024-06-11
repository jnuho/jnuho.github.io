package main

import (
	"fmt"
	"runtime"
	"sync"
)

const initialValue = -500

type counter struct {
	i    int64
	mu   sync.Mutex
	once sync.Once
}

func (c *counter) increment() {
	c.once.Do(func() {
		c.i = initialValue
	})

	c.mu.Lock()
	c.i++
	c.mu.Unlock()
}

func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}          // 카운터 생성
	done := make(chan struct{}) // 완료 신호 수신용 채널

	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()
			done <- struct{}{}
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	c.display()
}
