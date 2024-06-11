package main

import (
	"fmt"
	"sync"
)

type counter struct {
	i  int64
	wg sync.WaitGroup
	mu sync.Mutex
}

func (c *counter) increment() {
	defer c.wg.Done()
	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

func main() {
	c := counter{i: 0}

	for i := 0; i < 1000; i++ {
		c.wg.Add(1)
		go c.increment()
	}

	c.wg.Wait()

	fmt.Println("Final Counter Value:", c.i)
}
