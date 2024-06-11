
### 고루틴과 동시성 프로그래밍

- [고루틴](#고루틴)
  - [sync.WaitGroup](#syncwaitgroup)
  - [sync.Once](#synconce)
- [고루틴의 동작방법](#고루틴의-동작방법)
  - [동시성 프로그래밍 주의점](#동시성-프로그래밍-주의점)
  - [뮤텍스를 이용한 동시성 문제 해결](#뮤텍스를-이용한-동시성-문제-해결)
  - [mutex to ensure atomic access to a shared variable](#mutex-to-ensure-atomic-access-to-a-shared-variable)
    - [Using mutex and WaitGroup](#using-mutex-and-waitgroup)
    - [Using mutex and done channel](#using-mutex-and-done-channel)
    - [More done channel example](#more-done-channel-example)
  - [뮤텍스의 문제점](#뮤텍스의-문제점)
  - [또 다른 자원 관리 기법](#또-다른-자원-관리-기법)
- [채널](#채널)
  - [go channel with range and close](#go-channel-with-range-and-close)
  - [채널 크기](#채널-크기)
  - [채널에서 데이터 대기](#채널에서-데이터-대기)
  - [SELECT 문](#SELECT-문)
  - [일정간격으로 실행](#일정간격으로-실행)
  - [채널로 생산자 소비자 패턴 구현](#채널로-생산자-소비자-패턴-구현)
  - [unbuffered vs. buffered channel](#buffered-vs-unbuffered-channel)
- [컨텍스트](#컨텍스트)
  - [특정 값을 설정한 컨텍스트](#특정-값을-설정한-컨텍스트)
  - [작업시간 설정한 컨텍스트](#작업시간-설정한-컨텍스트)
  - [작업취소 가능한 컨텍스트](#작업취소-가능한-컨텍스트)
  - [취소도 되면서 값도 설정하는 컨텍스트 만들기](#취소도-되면서-값도-설정하는-컨텍스트-만들기)


### 고루틴

- 스레드란?
  - 고루틴: 경량 스레드로 함수나 명령을 동시에 실행 시 사용. main()도 고루틴에 의해 실행 됨
  - 프로세스 1개(싱글테스킹) vs 멀티프로세스 (멀티태스킹).
  - 프로세스 : 메모리 공간에 로딩되어 동작하는 프로그램. 1개 이상의 스레드를 가지고 있음
    - 스레드는 프로세스의 세부작업 단위 또는 실행 흐름. 1개, 2개, 멀티 스레드 있을 수 있음
  - CPU는 하나의 스레드 실행가능, 여러개 스레드를 번갈아가면서 수행하여 동시실행 처럼 보임

- 컨텍스트 스위칭 비용
  - 하나의 CPU가 여러개 thread 전환하면서 수행시 컨텍스트 스위칭 비용 발생
  - 스레드 전환시 현재 상태 보관해야 함 -> 스레드 컨텍스트를 저장
  - 스레드 컨텍스트 : 명령 포인터, 스택메모리 등
  - 컨텍스트 저장 및 복원 시 비용 발생
  - CPU 코어마다 OS스레드를 하나만 할당해서 사용하기 때문에, **Go언어**는 컨텍스트 스위칭 비용 발생 없음!

- 고루틴 사용
  - 모든 프로그램은 main()을 고루틴으로 하나 가지고 있음
  - `go 함수호출`로 새로운 고루틴 추가 가능.
  - 현재 고루틴이 아닌 새로운 고루틴에서 함수가 수행 됨

- Goroutines
  - Goroutines are lightweight concurrent functions or threads in Go.
  - They allow you to perform multiple tasks concurrently, leveraging parallelism on multi-core systems.
  - Unlike traditional threads, Goroutines are managed by the Go runtime, which dynamically allocates and manages their stack size.
  - Goroutines communicate using channels, which prevent race conditions when accessing shared memory.
  - Example: Fetching data from multiple APIs concurrently using Goroutines:

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://api.example.com/users",
		"https://api.example.com/products",
		"https://api.example.com/orders",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}

func fetchAPI(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}
```


[↑ Back to top](#)
<br><br>

- main 고루틴외에 PrintHangul, PrintNumber 고루틴 2개 추가생성
  - 3개 고루틴이 동시에 실행됨(CPU 3코어이상)
  - 1-코어 CPU에서는 동시에 실행되는 것 처럼 보임
  - main 고루틴은 종료하면 모든 고루틴이 즉시 종료되고, 프로그램이 종료됨
    - main 고루틴에서 3초간 기다려서 나머지 2개 고루틴도 실행되는 동안 wait함

```go
package main

import (
  "fmt"
  "time"
)

func PrintHangul() {
  hanguls := []rune{'가','나','다','라','마','바', '사'}
  for _, v := range hanguls {
    time.Sleep(300 * time.Millisecond)
    fmt.Printf("%c", v)
  }
}
func PrintNumbers(){
  for i:=1; i<=5; i++ {
    time.Sleep(400 * time.Millisecond)
    fmt.Printf("%d ", i)
  }
}

func main() {
  // PrintHangul, PrintNumbers가 동시에 실행
  go PrintHangul()
  go PrintNumbers()

  // 기다리지 않으면 메인함수 종료되고 모든 고루틴도 종료되기 때문에 대기시간 지정
  // 하지만 3초 라는 시간처럼 항상 시간을 계산할 필요는 없음
  //  -> sync패키지 WaitGroup 객체 사용!
  time.Sleep(3*time.Second)
}
```

[↑ Back to top](#)
<br><br>

### sync.WaitGroup

- [synchronization primitives in go](https://medium.com/better-programming/using-synchronization-primitives-in-go-mutex-waitgroup-once-2e50359cb0a7)

- 서브 고루틴이 종료될 때까지 기다리기
  - 항상 고루틴의 종료시간에 맞춰 time.Sleep(종료까지걸리는시간) 호출할 수 없음
  - 고루틴이 끝날때까지 wait할수 있음: sync.WaitGroup 객체

- A `WaitGroup` is a synchronization primitive in Go that allows you
  - to wait for a collection of goroutines to finish executing.
  - It keeps track of the number of goroutines that are currently active.
    - `Add(n int)` increases the internal counter by n, indicating that n goroutines will be added to the group.
    - `Done()` decreases the counter by 1 when a goroutine completes its work.
    - `Wait()` blocks until the counter becomes zero (i.e., all goroutines have finished).

```go
// sync.WaitGroup 객체 사용
var wg sync.WaitGroup

// Add()로 완료해야 하는 작업개수 설정하고, 각 작업이 완료 될때마다 Done() 호출하여
// 남은 작업개수를 하나씩 줄여줌. Wait()은 전체 작업이 모두 완료될때까지 대기하게 됨
wg.Add(3)   // 작업개수 설정
wg.Done()   // 작업이 완료될 때마다 호출
wg.Wait()   // 모든 작업이 완료될 때까지 대기
```


```go
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
  fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a,b,sum)

  wg.Done() // wg의 남은 작업개수를 1씩 감소시킴
}

func main() {
  // Set the total work count: 10 goroutines will be created, increasing CPU utilization
  wg.Add(10) // 총 작업개수 설정: 10개의 고루틴 생성하여 CPU 점유율 증가

  for i:=0; i<10; i++ {
    go SumAtoB(1, 1000000)
  }

  wg.Wait() // 모든 작업이 완료 될때까지 (남은작업개수 = 0) 종료하지 않고 대기
  fmt.Println("모든 계산이 완료됐습니다.")
}
```


[↑ Back to top](#)
<br><br>

- synchronization primitives in go
  - https://medium.com/better-programming/using-synchronization-primitives-in-go-mutex-waitgroup-once-2e50359cb0a7

- `wait-group.example.go`

```go
package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	filesInHomeDir, err := os.ReadDir(homeDir)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(filesInHomeDir))

	fmt.Println("Printing files in", homeDir)

	for _, file := range filesInHomeDir {
    // wg.Add(1) // this also works instead of `wg.Add(len(filesInHomeDir))`
		// anon function with parameter type `os.DisEntry`
		// The `(file)` at the end of the expression immediately invokes
		// the anonymous function with the value of the file variable.
		go func(f os.DirEntry) {
			defer wg.Done()
			fmt.Println(f.Name())
		}(file)
	}

	wg.Wait()
	fmt.Println("finished....")
}
```


[↑ Back to top](#)
<br><br>


### sync.Once

- `once-example.go`
  - Say you’re building a REST API using the Go net/http package and you want a piece of code to be executed
  - only when the HTTP handler is called (e.g. a get a DB connection).
  - You can wrap that code with once.Do and rest assured that
  - it will be only run when the handler is invoked for the first time.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var once sync.Once

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("http handler start")
		once.Do(oneTimeOp)
		fmt.Println("http handler end")
		w.Write([]byte("done!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func oneTimeOp() {
	fmt.Println("one time op start")
	time.Sleep(3 * time.Second)
	fmt.Println("one time op end")
}

```


[↑ Back to top](#)
<br><br>



### 고루틴의 동작방법

- 고루틴은 명령을 수행하는 단일 흐름으로, OS 스레드를 이용하는 경량 스레드
- 2-Core 컴퓨터 가정
- Goroutine 1개 일때:
  - `Core1-OsThread1-GoRoutine1`
- Goroutine 2개 일때:
  - `Core1-OsThread1-GoRoutine1`
  - `Core2-OsThread2-GoRoutine2`
- Goroutine 3개 일때
  - `Core1-OsThread1-GoRoutine1`
  - `Core2-OsThread2-GoRoutine2`
  - 작업끝날때 까지 대기 후 끝난 2코어에 '고루틴_3' 배정됨
  - `Core1-OsThread1-GoRoutine1`
  - `Core2-OsThread2-GoRoutine3`
- GoRoutine3가 시스템콜 호출 후 네트워크 대기상태 일때:
  - GoRoutine3가 대기목록으로 빠지고, GoRoutin4가 스레드2를 이용하여 실행됨
  - `Core1-OsThread1-GoRoutine1`
  - `Core2-OsThread2-GoRoutine4`
  - 컨텍스트 스위칭은 CPU가 Thread를 변경할 때 발생하는데, 코어와 스레드는 변경되지않고, 오로지 고루틴만 옮겨 다니기 때문에,
  - **고루틴을 사용하면, 컨텍스트 스위칭 비용이 발생하지않는다!**


- 시스템 콜 호출 시 (운영체제가 지원하는 서비스를 호출)
  - (고루틴으로 시스템콜 호출시; e.g. 네트워크로 데이터 읽을 때 데이터 들어올때 까지 고루틴이 대기상태 됨), 
  - 네트워크 수신 대기상태인 고루틴이 대기목록으로 빠지고, 대기중이던 다른 고루틴이 OS 스레드를 이용하여 실행 됨
  - 코어와 스레드 변경(컨텍스트 스위칭) 없이 고루틴이 옮겨다니기 때문에 효율적
  - 코어가 스레드 옮겨다니는 컨텍스트 스위칭을 하지 않고, 대신 고루틴이 직접 대기상태 <-> 실행상태 스위칭 옮겨다녀서 효율적


[↑ Back to top](#)
<br><br>


### 동시성 프로그래밍 주의점

- 문제점: `하나의/동일한 메모리 자원`에 `여러개 고루틴` 접근!
  - e.g. 입금1000, 출금1000 을 10개의 고루틴이 동시 실행 하는 상황
  - 두개 고루틴이 각각 1000원 입금했는데 2000이 아닌 1000이된상태에서 다시 두번 출금시 < 0 : panic!
  - `race condition`
- 해결책: 한 고루틴에서 값을 변경할때 다른 고루틴이 접근하지 못하도록 `mutex` 활용
  - `mutual exclusion`

- A `race condition` occurs when two or more threads access shared data and attempt to modify it simultaneously. 
  - To prevent race conditions, you typically use locks or synchronization mechanisms.
  - A lock ensures that only one thread can access the shared data at a time.

- Example of `race condition`

```go
package main
import (
  "fmt"
  "sync"
  "time"
)

type Account struct {
  Balance int
}

func DepositAndWithdraw(account *Account) {
  if account.Balance < 0 {
    panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
  }
  account.Balance += 1000
  time.Sleep(time.Millisecond)
  account.Balance -= 1000
}

func main() {
  var wg sync.WaitGroup

  account := &Account{0}
  wg.Add(10)

  for i:=0; i< 10; i++ {
    // 하나의 자원에 다수의 고루틴이 접근 함
    // 뮤텍스를 통해 Lock 걸어 해결
    go func() {
      for {
        DepositAndWithdraw(account)
      }
      wg.Done()
    }()
  }

  wg.Wait()
}
```


[↑ Back to top](#)
<br><br>

### 뮤텍스를 이용한 동시성 문제 해결

- 한 고루틴에서 값을 변경할때 다른 고루틴이 접근하지 못하도록 `mutex` 활용: `mutual exclusion`
- `mutex.Lock()`으로 mutex 획득 `mutext.Unlock()`으로 mutex 반납
- 다른 고루틴이 이미 뮤텍스를 획득했다면 해당 고루틴이 뮤텍스를 놓을 때(`mutex.Unlock()`) 까지 기다림
- 하지만 오직 하나의 고루틴만 공유 자원에 접근하기 때문에 동시성 프로그램 성능 향상 의미가 없어짐..
- 또한 `Deadlock` 발생 가능!

- A mutex (short for “mutual exclusion”) is a synchronization primitive used
- to protect shared resources in concurrent programs. It ensures that
- only one goroutine can access a critical section of code (a shared resource) at any given time.
- Mutexes prevent race conditions by allowing exclusive access to data.

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

var mutex sync.Mutex

type Account struct {
  Balance int
}

func DepositAndWithdraw(account *Account) {
  // 뮤텍스 획득!
  mutex.Lock()

  // 한번 획득한 뮤텍스는 반드시 Unlock() 호출하여 반납
  // `defer`: 함수 종료 전에 뮤텍스 Unlock() 메서드 호출
  defer mutex.Unlock()

  if account.Balance < 0 {
    panic(fmt.Sprint("Balance cannot be negative: %d", account.Balance))
  }

  account.Balance += 1000
  time.Sleep(time.Millisecond)
  account.Balance -= 1000
}


func main() {
  var wg sync.WaitGroup

  account := &Account{0}
  wg.Add(10)
  for i:=0; i< 10; i++ {
    go func() {
      for {
        DepositAndWithdraw(account)
      }
      wg.Done()
    }()
  }
  wg.Wait()
  
}
```


[↑ Back to top](#)
<br><br>

### mutex to ensure atomic access to a shared variable

A mutex helps achieve atomic access by allowing only one thread to hold the lock (mutex) at any given time.
The following code achieve 1.concurrency and 2.preventing race conditions.

### Using mutex and WaitGroup

```go
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
```

[↑ Back to top](#)
<br><br>

### Using mutex and done channel

- The done channel is used without a sync.WaitGroup.
- Each worker goroutine sends a signal to the done channel when it completes its task.
- The main goroutine waits for all workers to finish by receiving from the done channel.
- This approach allows us to manage concurrency without using a sync.WaitGroup

```go
package main
 
import (
    "fmt"
    "runtime"
    "sync"
)
 
const initialValue = -500
 
type counter struct {
    i int64
    mu sync.Mutex  // 공유 데이터 i를 보호하기 위한 뮤텍스
    once sync.Once // 한 번만 수행할 함수를 지정하기 위한 Once 구조체
}
 
// counter 값을 1씩 증가시킴
func (c *counter) increment() {
    // i 값 초기화 작업은 한 번만 수행되도록 once의 Do() 메서드로 실행
    c.once.Do(func() {
        c.i = initialValue
    })
     
    c.mu.Lock()   // i 값을 변경하는 부분(임계 영역)을 뮤텍스로 잠금
    c.i += 1      // 공유 데이터 변경
    c.mu.Unlock() // i 값을 변경 완료한 후 뮤텍스 잠금 해제
}
 
// counter의 값을 출력
func (c *counter) display() {
    fmt.Println(c.i)
}
 
func main() {
    // 모든 CPU를 사용하게 함
    runtime.GOMAXPROCS(runtime.NumCPU())
     
    c := counter{i: 0}          // 카운터 생성
    done := make(chan struct{}) // 완료 신호 수신용 채널
   
    // c.increment()를 실행하는 고루틴 1000개 실행
    for i := 0; i < 1000; i++ {
        go func() {
            c.increment()      // 카운터 값을 1 증가시킴
            done <- struct{}{} // done 채널에 완료 신호 전송
        }()
    }
     
    // 모든 고루틴이 완료될 때까지 대기
    for i := 0; i < 1000; i++ {
        <-done
    }
     
    c.display() // c의 값 출력
}
```

[↑ Back to top](#)
<br><br>

### More done channel example

https://gobyexample.com/closing-channels

```go
package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done

    _, ok := <-jobs
    fmt.Println("received more jobs:", ok)
}
```

[↑ Back to top](#)
<br><br>


### 뮤텍스의 문제점

1. 뮤텍스는 동시성 프로그래밍 성능이점 감소시킴
2. `데드락` 발생 가능
- e.g. 식탁에 A와 B가 각각 수저1, 포크1 집고있음.
- A,B가 포크1, 수저1 집으려 할때, A,B 누구하나 양보하지 않아, 밥을 먹을 수 없음: 두개 mutex 각각 차지
- `어떤 고루틴도 원하는 만큼 뮤텍스를 확보하지 못해서 무한히 대기하는 경우`; 데드락
- 멀티코어 환경에서는 여러 고루틴으로 성능 향상 가능
- 같은메모리 접근시 꼬일 수 있음
- 뮤텍스로 고루틴 하나만 접근하도록 하여 꼬이는 문제 해결 가능
- 하지만, 뮤텍스를 잘못 사용하면 성능향상 없이 데드락 발생가능
  - 뮤텍스 사용시 좁은 범위에서 사용하여 데드락 발생 방지
  - 또는 둘다 수저-> 포크 순서로 뮤텍스 락 사용하면 해결 가능

- 멀티코어 컴퓨터에서는 여러 고루틴을 사용하여 성능 향상
- 하지만 같은 메모리를 여러 고루틴이 접근하면 프로그램이 꼬일 수 있음
- 뮤텍스를 이용하면 동시에 고루틴 하나만 접근하도록 저장해 꼬이는 문제를 막을 수 있다.
- 그러나 뮤텍스를 잘못 사용하면 성능 향상도 못하고 데드락이라는 심각한 문제가 생길 수 있다.

- Deadlock in goroutines (not in using mutex): different scenario than above `mutex` deadlock
  - A deadlock occurs when a group of goroutines are waiting for each other, and none of them can proceed.
  - Essentially, they’re stuck in a circular dependency, unable to make progress.

- `Deadlock` 예시


```go
package main
import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)

var wg sync.WaitGroup

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {

  for i:= 0; i<100; i++ {
    fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
    first.Lock()
    fmt.Printf("%s %s 획득\n", name, fisrtName)
    second.Lock()
    fmt.Printf("%s %s 획득\n", name, secondName)
    fmt.Printf("%s 밥을 먹습니다.\n", name)
  
    time.Sleep(time.Duration(rand.Intn(1000))* time.Millisecond)
  
    second.Unlock()
    first.Unlock()
  }

  wg.Done()
}

func main() {
  rand.Seed(time.Now().UnixNano())

  wg.Add(2)
  fork := &sync.Mutex{}
  spoon := &sync.Mutex{}

  go diningProblem("A", fork, spoon, "포크", "수저")
  go diningProblem("B", spoon, fork, "수저", "포크")

  wg.Wait()
}
```


[↑ Back to top](#)
<br><br>



### 또 다른 자원 관리 기법

- 각 고루틴이 서로 다른 자원에 접근하게 만드는 두가지 방법
- mutex없이 동시성 프로그래밍 가능

- 영역을 나누는 방법
  - 각 고루틴은 할당된 작업만 하므로 고루틴(작업자)간 간섭 없음
  - 고루틴 간 간섭이 없어서 뮤텍스도 필요 없음
- 역할을 나누는 방법 : Channel과 함께 설명


- 각 고루틴은 할당된 작업만 하므로 고루틴간 간섭이 발생하지 않아서, Mutex가 필요없음

```go
package main
import (
  "fmt"
  "sync"
  "time"
)

type Job interface {
  Do()
}

type SquareJob struct {
  index int
}

func (j *SquareJob) Do() {
  fmt.Printf("%d 작업 시작\n", j.index)
  time.Sleep(1 * time.Second)
  fmt.Printf("%d 작업 완료 - 작업결과: %d\n", j.index, j.index * j.index)
}

func main() {
  jobList := [10]Job

  for i:=0 ; i< len(jobList); i++ {
    jobList[i] = &SquareJob{i}
  }

  var wg sync.WaitGroup
  wg.Add(10)

  for i:=0; i<10; i++ {
    job := jobList[i]
    go func() {
      job.Do()
      wg.Done()
    }
  }
  wg.Wait()
}
```



[↑ Back to top](#)
<br><br>


### 채널

- 채널: 고루틴끼리 메시지를 전달 할 수 있는 메시지 큐
  - 메시지큐에 메시지가 쌓이게 되고
  - 메시지를 읽을 때는 처음온 메시지부터 차례대로 읽음

- 채널 인스턴스 생성
  - 채널을 사용하기 위해서는 먼저 채널 인스턴스를 만들어야 함


- Channels are reference types, meaning they are already a reference to the underlying data structure.
  - `square(ch chan int) {}`
  - no need to declare channel parameters as pointers since channel is already a reference type

```go
// 채널타입: chan string
//    chan: 채널키워드
//    string: 메시지 타입
var messages chan string = make(chan string)
```


1. 채널에 데이터 넣기

```go
var messages chan string = make(chan string)
messages <- "This is a message"
```

2. 채널에서 데이터 빼기

```go
// 채널에서 빼낸 데이터를 담을 변수
// "채널 messages에 데이터가 없으면 데이터가 들어올떄까지 '대기함'"
var msg string = <- messages
```

1. 생성한 goroutine에서 채널에 데이터 빼기 (consumer)
  - main goroutine에서 데이터 넣기 (provider)

```go
package main
import (
  "fmt"
  "sync"
  "time"
)

func square(wg *sync.WaitGroup, ch chan int) {
  // 데이터를 빼온다
  // 데이터 들어올때까지 대기
  n := <- ch

  time.Sleep(time.Second)
  fmt.Printf("Square: %d\n", n*n)

  wg.Done()
}

// main 고루틴과 square 고루틴이 동시 실행
// main 루틴에서 채널에 9를 넣어줄때까지 square루틴은 대기상태
// 1. main 루틴
// 2. square() 루틴
func main() {
  var wg sync.WaitGroup

  // 크기 0인 채널 생성 : 반드시 다른 고루틴이 채널에서 데이터 꺼내야 정상 종료
  // 어떤 고루틴도 데이터 빼지 않으면 모든 고루틴이 계속대기 하다가 deadlock 발생
  ch := make(chan int)

  wg.Add(1)

  // 데이터를 빼서 처리
  go square(&wg, ch)

  // 데이터 넣는다.
  ch <- 9

  // square내에서 main루틴에서 넣어준 채널 데이터 빼고 wg.Done() 완료 될떄까지 대기
  wg.Wait()
}
```

2. 생성한 goroutine에서 채널에 데이터 넣기 (provider)
  - main goroutine에서 데이터 뺴기 (consumer)

```go
package main

import (
	"fmt"
	"sync"
)

func computeAndSendResult(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	// Perform some computation
	result := 42

	// Send the result through the channel
	ch <- result
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	resultCh := make(chan int)

  // produce
	go computeAndSendResult(&wg, resultCh)

  // consume
	// Receive the result from the channel
	result := <-resultCh
	fmt.Println("Received result:", result)

	wg.Wait()

}
```


[↑ Back to top](#)
<br><br>

### go channel with range and close

 - https://medium.com/better-programming/manging-go-channels-with-range-and-close-98f93f6e8c0c

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func produce(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Printf("producer put i=%d\n", i)
	}
	// Without closing channel, the consumer will wait indefinitely for channel
	close(c)
	fmt.Println("producer finished.\n")
}

func consume(c <-chan int) {
	fmt.Println("consumer sleeps for 5 seconds...")
	time.Sleep(5 * time.Second)
	fmt.Println("consumer started")
	for i := range c {
		fmt.Printf("consumer gets i = %d\n", i)
	}
	fmt.Println("consumer finished. press ctrl+c to exit")
}

func main() {
	// Both producer and consumer goroutine do not have to coexist
	// i.e. even if the producer goroutine finishes (and closes the channel),
	// the consumer goroutine range loop will receive all the values.
	c := make(chan int, 5)

	// producer
	go produce(c)

	// consumer
	go consume(c)

	e := make(chan os.Signal)
	// `signal.Notify` registers a channel `e` to receive specific signals
	// -> list of signals to capture i.e. `syscall.SIGINT`(Ctrl+c), `syscall.SIGTERM`(termination), etc...
	signal.Notify(e, syscall.SIGINT, syscall.SIGTERM)
	// blocks the main goroutine until one of these signals is received.
	<-e
}
```


[↑ Back to top](#)
<br><br>

- Modify above so that it works without using `signal.Notify`
  - 1. use `done` channel of type struct{}
  - 2. use `sync.WaitGroup`


1. use `done` channel of type struct{}

```go
package main

import (
	"fmt"
	"time"
)

func produce(c chan<- int, done chan<- struct{}) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Printf("producer put i=%d\n", i)
	}
	// Without closing channel, the consumer will wait indefinitely for channel
	close(c)
	fmt.Println("producer finished.\n")
	done <- struct{}{}
}

func consume(c <-chan int, done chan<- struct{}) {
	fmt.Println("consumer sleeps for 5 seconds...")
	time.Sleep(5 * time.Second)
	fmt.Println("consumer started")
	for i := range c {
		fmt.Printf("consumer gets i = %d\n", i)
	}
	// fmt.Println("consumer finished. press ctrl+c to exit")
	fmt.Println("consumer finished.")
	done <- struct{}{}
}

func main() {
	// Both producer and consumer goroutine do not have to coexist
	// i.e. even if the producer goroutine finishes (and closes the channel),
	// the consumer goroutine range loop will receive all the values.
	done := make(chan struct{})
	c := make(chan int, 5)

	// producer
	go produce(c, done)

	// consumer
	go consume(c, done)

	// in other cases use sync.WaitGroup to make main goroutine wait
	// e := make(chan os.Signal)
	// // `signal.Notify` registers a channel `e` to receive specific signals
	// // -> list of signals to capture i.e. `syscall.SIGINT`(Ctrl+c), `syscall.SIGTERM`(termination), etc...
	// signal.Notify(e, syscall.SIGINT, syscall.SIGTERM)
	// // blocks the main goroutine until one of these signals is received.
	// <-e
	<-done
	<-done
}
```

2. use `sync.WaitGroup`

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func produce(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		c <- i
		fmt.Printf("producer put i=%d\n", i)
	}
	// Without closing channel, the consumer will wait indefinitely for channel
	close(c)
	fmt.Println("producer finished.\n")
}

func consume(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("consumer sleeps for 5 seconds...")
	time.Sleep(5 * time.Second)
	fmt.Println("consumer started")
	for i := range c {
		fmt.Printf("consumer gets i = %d\n", i)
	}
	// fmt.Println("consumer finished. press ctrl+c to exit")
	fmt.Println("consumer finished.")
}

func main() {
	var wg sync.WaitGroup

	// Both producer and consumer goroutine do not have to coexist
	// i.e. even if the producer goroutine finishes (and closes the channel),
	// the consumer goroutine range loop will receive all the values.
	c := make(chan int, 5)

	wg.Add(2)

	// producer
	go produce(c, &wg)

	// consumer
	go consume(c, &wg)

	// in other cases use sync.WaitGroup to make main goroutine wait
	// e := make(chan os.Signal)
	// // `signal.Notify` registers a channel `e` to receive specific signals
	// // -> list of signals to capture i.e. `syscall.SIGINT`(Ctrl+c), `syscall.SIGTERM`(termination), etc...
	// signal.Notify(e, syscall.SIGINT, syscall.SIGTERM)
	// // blocks the main goroutine until one of these signals is received.
	// <-e
	wg.Wait()
}
```


[↑ Back to top](#)
<br><br>

- Both producer and consumer goroutine do not have to coexist
  - i.e. even if the producer goroutine finishes (and closes the channel),
	- the consumer goroutine range loop will receive all the values.

- We can simulate this scenario by using a combination of:
  - A buffered channel in the producer
  - Delaying the consumer goroutine by adding a time.Sleep()


[↑ Back to top](#)
<br><br>


### 채널 크기

- 기본 채널크기 0
  - 채널크기0: 채널에 데이터 보관할 곳이 없으므로 데이터 빼갈때까지 대기
- 채널 만들때 버퍼 크기 설정 가능

```go
package main

import (
  "fmt"
)

func main() {
  ch := make(chan int)
  // **DEADLOCK 채널에 데이터 넣기만 하고 뺴지 않을때
  // 채널에 데이터 넣을떄 기본사이즈 0이기 때문에
  // 보관할 수 없으므로 채널에서 데이터 빼는 코드가 있어야 진행가능!: goroutine 1개 생성해서 해당 func()에서 <-ch 해줘야함
  // 데드락 발생! 택배 보관장소 없으면 문앞에서 기다려야 함
  // 데이터 보관할 수 있는 메모리영역: 버퍼
  ch <- 9
  fmt.Println("Never print")
}
```


[↑ Back to top](#)
<br><br>

- 버퍼 가진 채널
  - `var chan string messages = make(chan string, 2)`
  - 버퍼가 다 차면, 버퍼가 없는 크기 0 채널처럼 빈자리 생길때 까지 대기
  - 데이터를 빼주지 않으면 버퍼없을 때 처럼 고루틴이 멈추게됨



### 채널에서 데이터 대기

- 고루틴에서 데이터를 계속 기다리면서 데이터가 들어오면 작업을 수행

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

func square(wg *sync.WaitGroup, ch chan int) {
  // 채널에 데이터가 들어 올때까지 '계속' 기다림
  // 데드락 방지: square() 호출 밖에서 close(채널)로 채널이 닫히면
  // for문을 종료하여 프로그램 정상 종료하도록 함
  for n := range ch {
    fmt.Printf("Square: %d\n", n*n)
    time.Sleep(time.Second)
  }

  // 위에 for문에서 계속 채널로 들어오는 데이터 기다리는 동안은, 실행되지 않음
  // 데이터를 채널 ch에 모두 넣은 다음에 close(ch)로 채널을 닫으면
  // for eange에서 데이터를 처리하고나서 채널이 닫힌 상태라면 for문 종료함 -> wg.Done() 실행됨!
  wg.Done()
}

func main() {
  var wg sync.WaitGroup
  ch := make(chan int)

  wg.Add(1)
  go square(&wg, ch)

  // 10번만 데이터를 넣음
  // square 내에서 채널 데이터 계속 기다림
  for i :=0; i< 10; i++ {
    ch <- i * 2
  }

  // 작업완료를 기다리지만,
  // square() 내에서 wg.Done()이 실행 되지 않고 deadlock발생
  // 하지만 채널을 닫아서 데드락 방지 가능
  // 채널에서 데이터를 모두 빼낸 상태이고, 채널이 닫혔으면
  // for range 문을 빠져나가게 됨 -> wg.Done()이 실행됨!!
  close(ch)

  wg.Wait()
}
```


[↑ Back to top](#)
<br><br>

### SELECT 문

- 여러 채널을 동시에 기다릴 수 있음.
- 어떤 채널이라도 하나의 채널에서 데이터를 읽어오면 해당 구문을 실행하고 select문이 종료됨.
- 하나의 case만 처리되면 종료되기 때문에, 반복해서 데이터를 처리하고 싶다면 for문과 함께 사용 해야함
- 채널에 데이터가 들어오길 기다리는 대신, 다른작업 수행하거나, 여러채널 동시대기
- 여러개 채널을 동시에 기다림. 하나의 케이스만 처리되면 종료됨
- 반복된 데이터 처리를 하려면 for문도 같이 사용
- ch채널과 quit채널을 모두 기다림, ch채널먼저 시도하기 때문에
  - ch채널 먼저 읽다가 모두 읽고나서 quit에 true가 들어와서 읽으면서 return 실행

```go
package main
import (
  "fmt"
  "sync"
  "time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
  for {
    // ch와 quit 양쪽을 모두기다림
    //    ch채널의 데이터를 모두 읽으면,
    //    quit 채널데이터를 읽고, square() 함수가 종료됨
    select {
    case n := <-ch: // ch 채널에서 데이터를 빼낼 수 있을 때 실행
      fmt.Printf("Squared: %d\n", n*n)
      time.Sleep(time.Second)
    case <- quit: // quit 채널에서 데이터를 빼낼 수 있을 때 실행
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

  for i:=0; i<10; i++ {
    ch <- i
  }

  quit <- true
  wg.Wait()
}
```


[↑ Back to top](#)
<br><br>

### 일정간격으로 실행

```go
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
      fmt.Println("tick")
    case <- terminate:
      fmt.Println("terminated...")
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
```


[↑ Back to top](#)
<br><br>

### 채널로 생산자 소비자 패턴 구현

- 역할 나누는 방법
  - 컨베이어벨트: 차체생산->바퀴설치->도색->완성

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

// 공정순서:
//     Body -> Tire -> Color
// 생산자 소비자 패턴 (Producer Consumer Pattern) 또는 pipeline pattern
// MakeBody()루틴이 생산자, InstallTire()루틴은 소비자
// InstallTire()는 PaintCar()루틴에 대해서는 생산자
type Car struct {
  Body string
  Tire string
  Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

// 1초간격 차체생산하여 tireCh 채널에 데이터 넣음
// 10초 후 tireCh 채널 닫고 루틴종료
func MakeBody(tireCh chan *Car) { // 차체생산
  tick := time.Tick(time.Second)
  after := time.After(10 * time.Second)

  for {
    select {
    case <- tick:
      // Make a body
      car := &Car{}
      car.Body = "Sports car"
      tireCh <- car
    case <-after:
      close(tireCh)
      wg.Done()
      return
    }
  }

}

// tireCh채널에서 데이터 읽어서 바퀴설치하고 paintCh채널에 넣어줌
// tireCh채널 닫히면 루틴종료하고 paintCh채널 닫아줌
func InstallTire(tireCh, paintCh chan *Car) { // 바퀴설치
    for car := range tireCh {
      // Make a body
      time.Sleep(time.Second)
      car.Tire = "Winter tire"
      paintCh <- car
    }
    wg.Done()
    close(paintCh)
}


// paintCh채널에서 데이터 읽어서 도색을 하고, 완성된 차 출력
func PaintCar(paintCh chan *Car) { // 도색
  for car := range paintCh {
    // Make a body
    time.Sleep(time.Second)
    car.Color = "Red"
    duration := time.Now().Sub(startTime) // 경과 시간 출력
    fmt.Printf("%.2f Complete Car: %s %s %s\n",duration.Seconds(), car.Body, car.Tire, car.Color)
  }

  wg.Done()
}

func main() {
  tireCh := make(chan *Car)
  paintCh := make(chan *Car)

  fmt.Println("Start the factory")

  wg.Add(3)
  go MakeBody(tireCh)
  go InstallTire(tireCh, paintCh)
  go PaintCar(paintCh)

  wg.Wait()
  fmt.Println("Close the factory")
}
```


[↑ Back to top](#)
<br><br>

### mutex to ensure atomic access to a shared variable

A mutex helps achieve atomic access by allowing only one thread to hold the lock (mutex) at any given time.

```go
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

func main() {
	c := counter{i: 0}

	for i := 0; i < 1000; i++ {
		c.wg.Add(1)
		go func(num int) {
			defer c.wg.Done()
			c.mu.Lock()
			c.i += 1
			c.mu.Unlock()
		}(i)
	}

	c.wg.Wait()

	fmt.Println("Final Counter Value:", c.i)
}
```

[↑ Back to top](#)
<br><br>

### buffered vs. unbuffered channel

If you don't explicitly close an unbuffered (non-buffered) channel in Go,
it won't cause any immediate issues. However, there are important implications:

1. **Blocking Behavior**:
   - If a goroutine tries to receive from an unbuffered channel and no other goroutine is sending to it, the receiver will block indefinitely.
   - Similarly, if a goroutine tries to send to an unbuffered channel and no other goroutine is receiving from it, the sender will block until another goroutine starts receiving.

2. **Resource Leaks**:
   - If you forget to close an unbuffered channel, it remains open indefinitely.
   - This can lead to resource leaks, especially if the channel is used for synchronization or signaling.
   - Properly closing channels ensures that goroutines can exit gracefully when their work is done.

3. **Signaling Completion**:
   - Closing an unbuffered channel is often used to signal the end of communication between goroutines.
   - It allows the receiving goroutine to know that no more values will be sent.
   - When the sender closes the channel, the receiver can detect this and exit gracefully.

In summary, while not explicitly closing an unbuffered channel won't cause immediate errors, it's good practice to close channels when they are no longer needed. This helps prevent blocking issues and ensures proper resource management. 😊

[↑ Back to top](#)
<br><br>

### 컨텍스트

- 컨텍스트 사용하기
  - 컨텍스트는 작업을 지시할때 작업가능 시간, 작업 취소 등의 조건을 지시할 수 있는 작업명세서 역할
  - 새로운 고루틴으로 작업 시작할떄 일정시간 동안만 작업지시하거나, 외부에서 작업 취소시 사용.
  - 작업 설정에 대한 데이터도 전달 가능
  - The `context` package provides tools for managing concurrent operations.
  - It allows you to control the lifecycle, cancellation, and propagation of requests across multiple Goroutines.
  - Key features:
    - Cancellation: Propagate cancellation signals across Goroutines.
    - Deadlines: Set deadlines for operations.
    - Values: Pass request-scoped data down the chain.
  - Example: Managing timeouts for API requests using context.WithTimeout.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://api.example.com/users",
		"https://api.example.com/products",
		"https://api.example.com/orders",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}

func fetchAPI(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}
```


[↑ Back to top](#)
<br><br>


### 작업취소 가능한 컨텍스트

- 이 컨텍스트를 만들어, 작업자에게 전달하면 작업 지시한 지시자가 원할때 작업취소 알릴 수 있음

```go
package main

import (
  "fmt"
  "sync"
  "time"
  "context"
)

var wg sync.WaitGroup

// 작업이 취소될 때까지 1초마다 메시지 출력하는 고루틴
func PrintEverySecond(ctx context.Context) {
  tick := time.Tick(time.Second)
  for {
    select {
    case <-ctx.Done():
      wg.Done()
      return
    case <-tick:
      fmt.Println("tick")
    }
  }
}

func main() {
  wg.Add(1)

  // 취소 가능한 컨텍스트 생성 : 컨텍스트 개체와 취소함수 반환
  ctx, cancel := context.WithCancel(context.Background())

  go PrintEverySecond(ctx)
  time.Sleep(5 * time.Second)

  // 작업취소
  //     컨텍스트의 Done()채널에 시그널을 보내, 작업자가 작업 취소하도록 알림
  //    <-ctx.Done() 채널
  cancel()

  wg.Wait()
}
```


[↑ Back to top](#)
<br><br>

### 작업시간 설정한 컨텍스트

- 일정시간 동안만 작업을 지시할 수 있는 컨텍스트 생성
- WithTimeout() 두번째 인수로 시간을 설정하면, 그시간이 지난 뒤
- 컨텍스트의 Done()채널에 시그널을 보내서 작업 종료
- WithTimeout() 역시 두번째 반환값으로 cancel함수 반환하기 때문에
- 작업 시작전에 원하면 언제든지 작업 취소 가능

```go
func main() {
  wg.Add(1)

  // 5초 후 컨텍스트의 Done()채널에 시그널을 보내 작업종료 요청
  ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
  go PrintEverySecond(ctx)

  wg.Wait()
}
```

[↑ Back to top](#)
<br><br>

### 특정 값을 설정한 컨텍스트

- 작업자에게 지시할때 별도의 지시사항 추가 가능
- 컨텍스트에 특정키로 값을 읽어오도록 설정 가능
- context.WithValue()로 컨텍스트에 값 설정 가능

```go
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
```

[↑ Back to top](#)
<br><br>


### 취소도 되면서 값도 설정하는 컨텍스트 만들기

- 컨텍스트를 만들때 항상 상위 컨텍스트 객체를 인수로 넣어줘야 했음
- 일반적으로 context.Background()를 넣어줬는데, 여기에 이미 만들어진 컨텍스트 객체 넣어도 됨
- 이를통해 여러 값을 설정하거나, 기능을 설정할 수 있음
- 구글: "golang concurrency patterns"

```go
ctx, cancel := context.WithCancel(context.Background())
ctx = context.WithValue(ctx, "number", 9)
ctx = context.WithValue(ctx, "keyword", "Lilly")
```

[↑ Back to top](#)
<br>
