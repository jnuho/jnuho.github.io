
- Concurrency vs Parallelism
	- Concurrent :
		- 동시성: 서버에서 여러개 클라이언트 요청을 동시에 처리
		- 서버가 동시에 처리할 수 있는 클라이언트 수를 늘려주지만, 요청하나하나의 속도는 그대로
	- Parallel : 
		- 병렬: 멀티스레드 방식으로 하나의 연산을 Parallel 방식으로 처리


1. What is Go language and what makes it different from other programming languages?
  - statically-typed: data types must be known at compile-time
    - contrary to dynamically-typed in which types can change at runtime
    - type safety
    - improved performance
    - better documentation
    - easier refactoring
	- 스태틱타입 언어로 컴파일 시 데이터타입이 정해져야하며,  type-safety, 성능 등의 이점이 있습니다.
  - concurrency
  - garbage collection
  - standard library
  - cross platform
  - strong typing
2. Can you explain the Go routine and how it works?
  - A Goroutine is a lightweight thread of execution in Go.
	- 고루틴은 경량 스레드의 실행(execution)으로 다른 function들과 
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

func printNumbers(ch chan int) {
  for i:=0; i< 10; i++ {
    ch <- i
    time.Sleep(100*time.Millisecond)
  }

  close(ch)
}

func main() {
  ch := make(chan int)
  go printNumbers(ch)

  for n:= range c {
    fmt.Println(n)
  }
}
```

- 채널
	- 초기화
		- `ch := make(chan int)`
	- 초기화 시 크기 지정
		- `ch := make(chan int, 5)`
	- 고루틴간 1. communication과 2. 동기화를 가능하게 하기 위한 방법으로 채널을 사용

- 1. 채널 커뮤니케이션

```go
import (
	"fmt"
)

func printNumbers(ch chan int) {
	for i:=0; i<10; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int)

	go printNumbers(ch)

	for n := range(ch) {
		fmt.Println(n)
	}
}
```

- 2. 채널 동기화
	- Use channel to wait for a goroutine to finish its work, and then proceeding to the next operations
	- Wait until the goroutine is finished, then proceed to the next operation
	- Sending to the channel when the operation is done, blocking until the message is received

```go
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	// Do some work
	// ... doing work ...
	fmt.Println("Work Started : ", time.Now())
	time.Sleep(3000 * time.Millisecond)

	// Signal that the work is done
	done <- true
}

func main() {
	done := make(chan bool)

	// Start the worker in a goroutine
	go worker(done)

	if <-done {
		fmt.Println("Work done : ", time.Now())
	}
}
```

- 채널 built-in 함수 `close(채널)`

```
In Go, close(채널) is a built-in function that is used to close a channel.
When a channel is closed, it can no longer be written to,
close(채널)은 해당 채널로의 write이 더이상 안되도록 함
다만 이미 write 된 데이터를 읽는 작업은 가능
but it can still be read from until all values that have been written to it have been read.
Closing a channel is important for signaling to the receiver that no more values will be sent.
```

- 채널을 동기화에 사용하는 추가 예시

```go


```

3. What is the difference between a slice and an array in Go?
  - array: fixed size
  - slice : dynmically sized. length can change at rumtime
		- can be created from an array or slice by slicing between indices
		- 크기가 동적으로 관리 되며, array나 slice로 부터 생성 가능 (index 활용)

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
	- built-in

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
  - a mechanism for synchronizing and communicating between goroutines
  - can be used to send and receive values between goroutines.

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

- WaitGroup :
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
  fmt.Printf("Worker %d starting\n", id)
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

```
Organizing your Go code into packages and following a structure
can help to make your code more maintainable and reusable.

In Go, a package is a collection of related Go source files.
Each Go source file starts with a package declaration,
which specifies the name of the package that the file belongs to.

Go provides several standard packages, such as fmt, math, and net/http,
that you can use to perform common tasks.
You can also create your own custom packages to group related functionality.
Here are some guidelines for structuring your Go code and organizing packages:
Group related functionality into separate packages:
For example, you can create separate packages for handling database access,
handling HTTP requests, or handling file I/O.

Use descriptive and meaningful package names:
The name of the package should clearly describe its purpose, such as database, http, or fileio.

Define the API for your package: Define the functions and types
that are part of the public API for your package and keep the implementation details hidden.
You can use Go's visibility rules (capitalized identifiers are exported, lowercase identifiers are not)
to control the visibility of your package's API.

Document your code: Use Go's built-in documentation capabilities to document your code.
This makes it easier for others to understand and use your code.

Keep your code modular: Avoid including too much functionality in a single package,
as this can make your code difficult to maintain and test.
Instead, keep your code modular by breaking it down into smaller,
focused packages that can be easily composed to build larger, more complex systems.

By following these guidelines, you can structure your Go code
in a way that makes it easier to maintain, test, and reuse.
```

8. Can you explain the defer keyword in Go and provide an example of its use?
	- execute a function later, after the function that contains the defer statement has completed
	- even if the surrounding function panics, the deferred function call is executed!
	- clean up resources, closing files, releasing locks, printing log messages
	- e.g. defer closing files, after reading or writing

```go
import (
	"fmt"
	"os"
)
func main() {
	file, err := os.Open("filename.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
```

9. How do you handle HTTP requests in Go and what packages do you typically use for that?

```
'net/http' package for handling HTTP requests
which provides a set of functions and types for building HTTP clients and servers.
`http.ListenAndServe` to create an HTTP Server, and pass it to `http.Handler` interface.

* 3rd-party
github.com/gorilla/mux: A powerful HTTP router and URL dispatcher.
github.com/go-chi/chi: A lightweight and idiomatic HTTP request multiplexer.
github.com/labstack/echo: A fast and flexible HTTP server framework.
github.com/json-iterator/go: A high-performance JSON parsing library.
```


```go
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// writes "Hello World" to the response writer
		fmt.Fprintf(w, "Hello World")
	})

	http.ListenAndServe(":8080", nil)
}
```


10. Can you explain the difference between interface and struct in Go and when to use each?

```
Struct is a user-defined composite data type
	that groups together zero or more values of different types into a single entity.
An interface, on the other hand, is a collection of method signatures
	that describe a set of behaviors that a type can exhibit

Here are some key differences between structs and interfaces:

A struct is a concrete data type that has a fixed set of fields,
	each with a specific type,
whereas an interface is an abstract type that defines a set of method signatures
	that a concrete type can implement.

A struct can be instantiated and its fields can be accessed directly,
whereas an interface cannot be instantiated and its methods
	can only be invoked on a concrete type that implements the interface.

A struct is used to group related data into a single entity,
while an interface is used to define a set of behaviors
	that a type must implement to be considered part of a particular category.

When to use a struct:
	When you need to group related data together into a single entity.
	When you need to access and modify the individual fields of a data type.
	When you need to create a new instance of a data type.

When to use an interface:
	When you want to define a set of behaviors that a type must implement to be considered part of a particular category.
	When you want to abstract away the implementation details of a type and only interact with its public API.
	When you want to create reusable code that works with different concrete types that implement the same interface.

In general, structs are used to represent data
and interfaces are used to represent behaviors.

Structs are useful for organizing and working with data in a program,
while interfaces are useful for defining common functionality across different types.

It's common to use both structs and interfaces together to create a well-organized and flexible codebase.
```

11. Can you give an example of Go code that demonstrates the use of channels for communication between go routines?


```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(){
		for i := 0; i<5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Retrieve values from the channel in the main goroutine
	for val := range(ch) {
		fmt.Println(val)
	}
}
```

```
In this example, we create a channel of type int using the make function.
We then start a new goroutine that sends integers through the channel using the <- operator,
and close the channel after we're done sending values.
In the main goroutine, we use the range keyword
to receive values from the channel until it is closed, and print each received value to the console.

Channels provide a powerful mechanism for communication and synchronization between goroutines.
By sending and receiving values through channels, goroutines can coordinate their behavior
and exchange data in a safe and controlled manner.
Channels also provide a way to block a goroutine until a value is available,
which is useful for coordinating concurrent operations.
```

12. Can you explain the use of Go testing and give an example of a test case?


```
Testing is an important aspect of software development, and Go provides a built-in testing framework to make it easier for developers to write and run tests.

In Go, test files are usually suffixed with _test.go and are organized into packages just like regular source files. Tests can be written using the testing package, which provides a set of functions for creating and running tests.

Here's an example of a test case for a simple sum function:

```

```go
package main

import "testing"

func TestSum(t *testing.T) {
    total := sum(3, 5)
    if total != 8 {
        t.Errorf("sum(3, 5) = %d; expected 8", total)
    }
}

func sum(a, b int) int {
    return a + b
}

```

```
In this example, we have a function called sum that takes two integers and returns their sum. We also have a test function called TestSum that uses the sum function and checks if the result is correct using the if statement and the t.Errorf function.

To run the test, we can use the go test command in the terminal at the directory where the test file is located. This command will discover and run all test functions in the package and report the results.

The output of the go test command should look like this:
```

```sh
go test
```


13. Can you discuss the benefits of using Go for microservices?

```
Yes, here are some benefits of using Go for microservices:

- Fast performance: Go is a compiled language that is optimized for speed and efficient memory usage,
which makes it ideal for building microservices that require high performance.

- Concurrency support: Go has built-in support for concurrency,
which allows developers to easily create concurrent, parallel, and distributed systems.
This is especially important for microservices, which typically need to handle many requests simultaneously.

- Small and simple: Go has a small and simple syntax, which makes it easy to learn and use.
This simplicity also makes it easier to write and maintain microservices,
which tend to be smaller and more focused than monolithic applications.

- Easy to deploy: Go produces statically compiled binaries that can be deployed easily
and run on any platform without requiring a separate runtime environment.

- Strong typing and safety: Go's strong typing and safety features help prevent
common programming errors and make it easier to write reliable and secure microservices.

Overall, Go's combination of speed, concurrency support, simplicity, ease of deployment,
and safety features make it a great choice for building microservices.
```

14. How do you handle dependencies in Go and what tools do you use for dependency management?

```
Dependency management is an important aspect of any software development project, including those written in Go. In Go, there are several ways to handle dependencies and several popular tools for dependency management.

Here are a few ways to handle dependencies in Go:


1. Go modules: Go modules are a built-in feature in Go that provide versioned dependencies and a solution for dependency management. With modules, dependencies are listed in a go.mod file at the root of the project, and the go command automatically downloads and builds the correct versions of the dependencies. Go modules are the recommended way to manage dependencies in Go projects.

2. Vendoring: Vendoring involves copying the dependencies of a project into the project's repository, so that the project has a self-contained set of dependencies. This can be done manually, or with tools like govendor or dep. Vendoring has been largely replaced by Go modules, but it is still supported for older projects.

3. Package managers: There are several third-party package managers for Go that can be used for dependency management, including Godep, Glide, and dep. These tools provide a way to manage dependencies outside of the Go standard library, and can be useful for older projects that have not yet migrated to Go modules.


As mentioned earlier, Go modules are the recommended way to manage dependencies in Go projects. Here's an example of using Go modules to manage dependencies:

1. Initialize the module by running the go mod init command in the root directory of the project. This will create a go.mod file that lists the dependencies of the project.
$ cd my-project
$ go mod init example.com/my-project

2.  Add dependencies to the go.mod file using the go get command.
$ go get github.com/gin-gonic/gin

3.  Import the dependencies in the project's source code.
import "github.com/gin-gonic/gin"

4.  Build and run the project with the go command, which will automatically download and build the correct versions of the dependencies.
$ go build
$ ./my-project

Overall, Go provides several ways to manage dependencies, with Go modules being the recommended solution. Other tools like vendoring and package managers can also be used, but they are not the preferred way to handle dependencies in Go.

```

15. Can you provide an example of using reflection in Go?
- Reflection is a powerful feature in Go that allows you to inspect and manipulate
- the structure of a value at runtime. Here's a simple example of how to use reflection in Go:

```go
package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name    string
    Age     int
    Address string
}

func main() {
    p := Person{
        Name:    "John Doe",
        Age:     35,
        Address: "123 Main St",
    }

    v := reflect.ValueOf(p)
    t := v.Type()

    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        fmt.Printf("%s: %v\n", t.Field(i).Name, field.Interface())
    }
}

```

```
In this example, we have a Person struct with three fields: Name, Age, and Address. In the main function, we create an instance of the Person struct and use the reflect.ValueOf function to get a reflect.Value object representing the struct. We also use the Type method of the reflect.Value object to get a reflect.Type object representing the struct's type.

We then loop through the fields of the struct using the NumField method of the reflect.Value object, and use the Field method to get the value of each field. We also use the Interface method of the reflect.Value object to get the value of the field as an interface{}.

Finally, we use the fmt.Printf function to print the name and value of each field.

When we run this program, we should see output like this:
```

```makefile
Name: John Doe
Age: 35
Address: 123 Main St
```


```
This example demonstrates how reflection can be used to inspect the fields of a struct at runtime, regardless of the struct's type. Reflection can also be used to manipulate objects at runtime, create new instances of types, and much more.
```
