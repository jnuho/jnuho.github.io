
- Go 설계 목표
	- 대규모 코드베이스 빠르게 처리할 수 있는 컴파이러
	- 최적화 활용하여 빠르게 코드를 생성 할 수 있는 컴파일러

- Go 특징
	- 코드문법에 엄격 (strict)
	- 간단한 컴파일러 구조: Go 최적화는 어셈블리 코드에서 이뤄짐
	- 어셈블리는 텍스트 형태로의 중간과정 없이 바이너리 형태로 생성됨
	- 코드 안정성 : 컴파일러에 오버헤드 부여하고, 런타임시 부하가 발생하는 Swift와 달리
		문법이 간략해서, 컴파일 시간과 런타임에서 매우 간단하게 안정성 구현가능

- Go 적용가능 플랫폼, 아키텍쳐
	- Linux, Windows, MacOS
	- x86/x86_64, ARM, ...

- Go 컴파일러와 런타임
	- 간결한 문법 (장단점)
	- 쿠버네티스, 도커 등 기술들과 결합하여 사용 용이
	- 컴파일러 빠른 동작
	- 메모리 수동 관리, 포인터 직접 다룸

- Go 메모리 관리와 가비지컬렉터
	- 프로그램에서 사용되지 않아 참조되지 않을 메모리 영역 찾아 삭제
	- 동작시, 애플리케이션을 중지 해야하므로, 언제 중단될 지 알 수 없으며, 메모리영역 또한 불확실성을 갖게 됨
	- 다만 Go의 가비지컬렉터는 매우 효율적으로 동작하여, 동작 중 지연 또는 끊김 현상이 거의 없음

- Go 동시성

- 프로젝트 구조
	- https://github.com/golang-standards/project-layout
	- 패키지 : 단일 소스 파일로서 확장자를 제외한 파일명이 패키지 이름이 됨
		- e.g. 패키지 `main` 디렉토리 안에 `main.go` 파일 존재

- 변수
	- string: 문자형태
	- int8/int16/int32/int64 : n-bit의 부호있는 (signed) 정수형
	- uint8/uint16/uint32/uint64 : n-bit의 부호없는 (unsigned) 정수형
	- uint/int : 32-bit 에서는 uint32/int32, 64-bit에서는 uint64/int64

```go
package main

// 전역 변수
var (
	name1 = "Tanmay Bakshi"
	age1 = 16
)

// 전역 상수
const name1 = "Tanmay Bakshi const"
const age2 = 100
```


- 슬라이스 slice
	- append로 배열의 재할당 작업 없이 길이를 알고 있다면 `make`을 활용하여 길이지정하여 생성
		- `arr := make([]int, 2)`
		- `arr[0] = 0`
		- `arr[1] = 1`
	- append로 특정 capacity까지는 요소의 추가로 인한 재할당 발생하지 않게 할 수 있음
		- `make([]int, 0, CAPACITY)`
		- `arr := make([]int, 2)`
		- `arr = append(arr, 0)`
		- `arr = append(arr, 1)`

- if/switch문

- 반복문
	- for문하나로 `for, while, do-while/repeat-while` 기능 모두 구현

```go
names := []string{"A", "B", "C"}

for i := range names {
	fmt.Println(i)
}

for i,v := range names {
	fmt.Println(i, v)
}

for _, v := range names {
	fmt.Println(v)
}
```

- 함수

- 다른함수에 함수 포인터 전달하기

```go
func runMathOp(a int, b int, op func(int, int)int) int {
	return op(a,b)
}

func add(a int, b int) int {return a+b}
func sub(a int, b int) int {return a-b}
func mul(a int, b int) int {return a*b}
func div(a int, b int) int {return a/b}

func main() {
	a,b := 9, 6
	fmt.Println(runMathOp(a,b, add))
	fmt.Println(runMathOp(a,b, sub))
	fmt.Println(runMathOp(a,b, mul))
	fmt.Println(runMathOp(a,b, div))
}

```

- `defer`
	- 함수를 반환하기 직전에 어떤 코드를 실행하라는 명령


- `generic`
	- functions and data types that can work with different types of data
	- type paramters are placeholders for a type that are defined when a generic function or type is used

```go
func MaxInts(nums[] int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}
```

- Example of a gneric function that finds the maximum value in a slice of integers
	- make it work with any type that implements `Comparable` interface.

```go
type Comparable interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string
}

func Max[T Comparable](nums []T) {
	max := nums[0]

	for _, n := range nums {
		if n> max {
			max = n
		}
	}

	return max
}
```




