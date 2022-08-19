package main
import (
  "fmt"
  "sync"
  "time"
)

func square(wg *sync.WaitGroup, ch chan int) {
  // 데이터를 빼온다
  n := <- ch

  time.Sleep(time.Second)
  fmt.Printf("Square: %d\n", n*n)

  wg.Done()
}

func main() {
  var wg sync.WaitGroup
  ch := make(chan int)

  wg.Add(1)
  go square(&wg, ch)
  ch <- 9
  wg.Wait()
}
