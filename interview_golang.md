

























1. What is Go language and what makes it different from other programming languages?
  - golang is statically-typed: data types must be known at compile-time (performance)
    - type saftety
    - improved performance
    - better documentation
    - easier refactoring
  - dynamicallyed-typed: types can change at runtime 
  - concurrency
  - garbage collection
  - standard library
  - cross platform
  - strong typing

2. Can you explain the Go routine and how it works?
  - A Goroutine is a lightweight thread of execution in Go.
  - It is a function that is capable of running concurrently with other functions.
  - Unlike traditional threads, Goroutines are cheaper and more efficient to create and manage,
  - as they are multiplexed onto multiple OS threads.
  - Goroutines are implemented using a system of channels and select statements,
  - which allow them to communicate with each other and coordinate their execution.

```go
package main

import (
  "fmt"
  "time"
)

func printNumbers(c chan int) {
  for i:=0; i< 10; i++ {
    c <- i
    time.Sleep(100*time.Millisecond)
  }

  close(c)
}

func main() {
  c := make(chan int)
  go printNumbers()

  for n:= range c {
    fmt.Println(n)
  }
}
```

3. What is the difference between a slice and an array in Go?
  - array: fixed size
  - slice : dynmically sized

4. How do you handle errors in Go? Can you give an example?
  - in Go, errors are represented as values of the built-in error type.
  - to handle errors, Go provides several techniques, including:
  * Return values (recommended)
  * Panic and recover
    - Go provides a panic function that can be used to halt
    - the normal flow of execution and jump to a recovery function

```go
func divide1(a, b int) (int, error) {
  if b == 0 {
    return 0, fmt.Errorf("cannot divide by zero")
  }
  return a/b, nil
}

func main() {
  // 1
  result, err := divide1(10, 0)
  if err != nil {
    fmt.Println("error: " + err)
  }

  // 2 panic and recover
}
```

5. How would you implement concurrency in Go?
  - go provides several built-in features for implementing concurrency
  - Goroutines
    - lightweight threads of execution that can be created with the go keyword.
    - They are used to run multiple functions simultaneously within a single program.
    - printNumbers function is run as a goroutine,
    - and the main function receives values from the channel c to print the numbers.

```go

func printNumbers(ch chan int) {
  for i := 0; i<10; i++ {
    ch <- i
  }
  close(ch)
}

func main() {
  ch := make(chan int)
  go printNumbers(ch)
  for num := range ch {
    fmt.Println(num)
  }
}
```

  - Channels
    - a mechanism for synchronizing and communicating between goroutines.
    - Channels can be used to send and receive values between goroutines.

```go
func add(a,b int) int {
  return a+b
}

func main() {
  result := make(chan int)
  go func() {
    result <- add(1,2)
  }()
  fmt.Println(<- result)
}
```

  - WaitGroup:
    - it is a type from the sync package in Go that allows you
    - to wait for multiple goroutines to finish executing.
    - Example using a WaitGroup to run multiple goroutines concurrently:


```go
import (
  "fmt"
  "sync"
)

func worker(id int, wg *sync.WaitGroup) {
  defer wg.Done()
  fmt.Prinf("Worker %d starting\n", id)
  for i:=0; i< 10; i++ {
    fmt.Prinf("Worker %d working\n", id)
  }
  fmt.Prinf("Worker %d finished\n", id)

}

// Add()로 완료해야 하는 작업개수 설정하고, 각 작업이 완료 될때마다 Done() 호출하여
// 남은 작업개수를 하나씩 중여줌. Wait()은 전체 작업이 모두 완료될때까지 대기하기 됨
func main() {
  var wg sync.WaitGroup

  for i :=0; i< 10; i++ {
    wg.Add(1)
    go worker(i, &wg)
  }
  wg.Wait()
  fmt.Println("All Workers finished")
}
```

6. Can you explain the Go memory model and how it manages memory allocation and garbage collection?

```
The Go memory model is a set of rules that dictate how Go programs should behave when accessing shared memory. It defines the behavior of the Go runtime with regards to memory allocation, garbage collection, and concurrency.

In Go, memory is allocated using the new and make functions. new is used to allocate zero-valued memory for a specified type, while make is used to allocate memory for built-in types such as maps, slices, and channels. Both functions return a pointer to the newly allocated memory.

Go uses a garbage collector to automatically reclaim memory that is no longer being used by the program. The garbage collector tracks the reachability of objects in memory, and frees objects that are not reachable. This means that Go programs don't have to worry about freeing memory manually, and can instead rely on the garbage collector to automatically manage memory usage.

The Go memory model also provides some guarantees with regards to concurrency. For example, it guarantees that a write to a variable by one goroutine will be visible to another goroutine after the first goroutine has finished writing. Additionally, it provides atomic memory operations that ensure that updates to a variable are indivisible, even in the presence of concurrent access. These guarantees make it easier to write concurrent Go programs that are correct and free of race conditions.

Overall, the Go memory model provides a high-level abstraction that allows Go programmers to focus on writing their programs, without having to worry about the low-level details of memory allocation and garbage collection. The Go runtime takes care of these details and provides the necessary guarantees to ensure that Go programs behave correctly and efficiently.
```

7. How do you structure your Go code and organize packages?
8. Can you explain the defer keyword in Go and provide an example of its use?
9. How do you handle HTTP requests in Go and what packages do you typically use for that?
10. Can you explain the difference between interface and struct in Go and when to use each?
11. Can you give an example of Go code that demonstrates the use of channels for communication between go routines?
12. Can you explain the use of Go testing and give an example of a test case?
13. Can you discuss the benefits of using Go for microservices?
14. How do you handle dependencies in Go and what tools do you use for dependency management?
15. Can you provide an example of using reflection in Go?