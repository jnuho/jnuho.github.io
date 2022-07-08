package main

import (
  "fmt"
  "sync"
  "time"
)

func square(wg *sync.WaitGroup, ch chan int) {
  // 원하는 시간 간격으로 신호를 보내주는 채널을 만들 수 있음
  // 1초간격 주기로 시그널 보내주는 '채널' 생성하여 반환
  // func Tick(d Duration) <-chan Time
  //    return 타입: <-chan time.Time
  // 이채널에서 데이터를 읽어오면 일정 시간간격으로 현재 시각을 나타내는 Timer 객체를 반환
  tick := time.Tick(time.Second)  // 1초 간격 시그널

  // func After(d Duration) <-chan Time
  //    waits for the duration to elapse 
  //    and then sends the current time on the returned channel.
  // 이채널에서 데이터를 읽으면 일정시간 경과 후에 현재시각을 나타내는 Timer 객체를 반환
  terminate := time.After(10*time.Second) // 10초 이후 시그널

  for {
    select {
    case <- tick:
      fmt.Println("Tick", time.Now())
    case <- terminate:
      fmt.Println("Terminated")
      wg.Done()
      return
    case n := <-ch:
      fmt.Printf("Squared: %d\n", n*n)
      time.Sleep(time.Second)
    }
  }
}

func main() {
  var wg sync.WaitGroup
  ch := make(chan int)

  wg.Add(1)
  go square(&wg, ch)

  for i:=0; i<10; i++ {
    ch <-i
  }
  wg.Wait()
}