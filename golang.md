
- [AWS SDK](#aws-sdk)

### Learn Golang

- [go module](https://medium.com/@fonseka.live/getting-started-with-go-modules-b3dac652066d)

```sh
# fmt 라이브러리만 사용할때는 go run 파일명.go 실행 가능
# thrid party 라이브러리 사용 시 에는

docker run --rm -it golang:1.18
apt-get update && apt-get install vim -y
echo $GOPATH
  /go

# GOPATH 가 아닌 디렉토리에서 실행 /root/TestApp
# fmt 라이브러리만 사용하므로 정상 실행
cd ~
mkdir TestApp && cd /root/TestApp


# 외부 라이브러리 사용 시 It complains about not finding the package because :
#  1.GOPATH밖 2.패키지 다운로드 안했기 떄문
#  no required module provides package github.com/Pallinder/go-randomdata:
#  go.mod file not found in current directory or any parent directory;
#  use go module to fix this
# go mod init /root/TestApp
go mod init github.com/jnuho/goproject/TestApp

# go mod tidy fetch all the dependencies that you need for testing in your module.
go mod tidy

# 특정 버전의 패키지 사용
# 디폴트 : latest
go get github.com/bwmarrin/discordgo@v0.20.3

cat go.mod
  require github.com/bwmarrin/discordgo v0.20.3

go get github.com/bwmarrin/discordgo@latest

# 빌드 하기 전 go.mod 파일 필요
# go build 커멘드로 실행 파일 생성
go build
```


- Vscode > Extension > Go

```sh
### CREATE MODULE ###
cd goproject/calculator
# MODULE_NAME should ideally be on the web
go mod init github.com/jnuho/jnuho.github.io/goproject/calculator
cat go.mod

### CREATE PACKAGE ###
# e.g. two packages
mkdir operations display
# to use 'format.go' inside 'simple.go', use:
# import "github.com/jnuho/jnuho.github.io/goproject/calculator/display"
touch operations/simple.go
touch display/format.go

### COMMIT AND PUSH TO REPO ###

```

```go
### USE ABOVE PACKAGES ###
# mkdir simplego && cd simplego
# touch main.go
package main

import "github.com/jnuho/jnuho.github.io/goproject/calculator/operations"

func main() {
}
```


```sh
# look through import statements
# and downloads all missing packages/dependencies
# 외부 패키지 다운로드
go mod tidy

cat go.mod
```



- 코드 실행 단계
  - 폴더생성
  - .go 파일생성 및 작성
  - Go 모듈 생성 (.mod 생성: 모듈명,Go버전,필요패키지 목록 등 정보)
  - 빌드 (`go build`)
    - GOOS GOARCH 환경변수 조정하여 다른 OS및 아키텍쳐에서 실행가능한 파일 생성 가능
    - e.g. `GOOS=linux GOARCH=amd64 go build`
  - 실행

- Go1.16버전 부터 Go 모듈 사용이 기본이 됨
  - 모듈은 패키지 종속성 관리 등 패키지 관리 시스템 역할
  - `Ver < 1.16`: 모든 Go코드는 $GOPATH/src 아래
  - `Ver >= 1.16`: 모든 Go코드는 Go모듈 아래
    - `go mod init MODULE_NAME`으로 Go모듈 생성
    - 모듈이름은 유니크해야 함; GitHub 저장소 주소를 활용하거나 URL을 모듈 이름으로 사용

이름| 설명| 값의 범위
--|--|--
uint8 | 1바이트 부호 없는 정수 | 0~255
uint16 | 2바이트 부호 없는 정수 | 0~65535
uint32 | 4바이트 부호 없는 정수 | 0~4294...95
uint64 | 8바이트 부호 없는 정수 | 0~1844...615
int8 | 1바이트 부호 있는 정수 | -128~127
int16 | 2바이트 부호 있는 정수 | -32768~32767
int32 | 4바이트 부호 있는 정수 | -214...648~214..647
int64 | 8바이트 부호 있는 정수 | -922....808 ~ 922..807
float32 | 4바이트 실수 | IEEE~754 32비트 실수
float64 | 8바이트 실수 | IEEE~754 64비트 실수
complex64 | 8바이트 복소수(진수,가수) | 진수와 가수 범위는 float32 범위와 같음
complex128 | 16바이트 복소수(진수,가수) | 진수와 가수 범위는 float64 범위와 같음
byte | uint8의 별칭 1바이트 데이터 나타낼때 사용 | 0~255 rune | int32별칭 UTF-8로 문자하나 나타낼때 사용 | -214..648~214..647
int |  32비트 컴퓨터에서 int32. 64비트 컴퓨터에서 int64 |
uint |  32비트 컴퓨터에서 uint32. 64비트 컴퓨터에서 uint64 |


- 그외 타입
  - boolean
  - 문자열
  - 배열
  - 슬라이스
  - 구조체
  - 포인터
  - 함수타입
  - 인터페이스
  - 맵
  - 채널


- 변수선언 형태

```go
package main

import "fmt"

func main() {
  var a int = 3
  var b int // 타입별 기본값으로 대체 (초깃값 생략)
  var c = 4 // 우변값의 타입으로 대체 (타입생략)
  d := 5    // 우변값의 타입으로 대체 (var 키워드와 타입 생략)

  fmt.Println(a, b, c ,d)
}
```


### 6.연산자

- numericOperator.go

```go
package main

import (
  "fmt"
)

func main(){
  var x int32 = 7
  //var y float32 = 3.14

  fmt.Println(x + 1)
  fmt.Println(x - 1)
  fmt.Println(x * 1)
  fmt.Println(x / 1)
  fmt.Println(x % 1)

  fmt.Println(x & 1)
  fmt.Println(x | 1)
  fmt.Println(x ^ 1)
  fmt.Println(x &^ 1)

  fmt.Println(x << 1)
  fmt.Println(x >> 1)
}
```

- compOperator.go

```go
package main

import (
  "fmt"
)

func main(){
  var x int32 = 7
  var y float32 = 3.14

  fmt.Println(x == y)
  fmt.Println(x != y)
  fmt.Println(x * y)
  fmt.Println(x < y)
  fmt.Println(x > y)
  fmt.Println(x <= y)
  fmt.Println(x >= y)

}
```


- operationErr1.go
  - 컴퓨터는 2진수로만 표현-> 실수 표현시 오차발생 (e.g. 0.376 != 2^-2 + 2^-3+ ...)
  - 서로 다른 값이 오차에 의해 같다고 나오는 현상 테스트
  - epsilon 값을 정의하여 두 값의 차이가 해당 값보다 작으면 equal처리
-- operationErr2.go
  - 비트표현시 '지수부+실수부' 중 가장 오른쪽 비트 하나차이 이내이면 같다고 판단
  - 마지막 비트가 1만큼 차이 나는지 확인
    - `func Nextafter(x,y float64) (r float64)`
    - x에서 y를향 해서 1비트 조정 (+ or - 한 값을 반환)

```go
package main

import (
  "fmt"
  "math"
)

const epsilon = 0.000001

func equal(a, b float64) bool {
  diff := math.Abs(a-b)

  if diff <= epsilon {
    return true
  } else {
    return false
  }

}

func main() {
  var a float64 = 0.1
  var b float64 = 0.2
  var c float64 = 0.3

  fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
  fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b,c))

  a= 0.0000000000004
  b= 0.0000000000002
  c= 0.0000000000007

  fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b,c))

}
```

- operationErr2.go


```go

package main

import (
  "fmt"
  "math"
)

func equal(a, b float64) bool {
  return math.Nextafter(a,b) == b
}

func main() {
  var a float64 = 0.1
  var b float64 = 0.2
  var c float64 = 0.3

  fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
  fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b,c))

  a= 0.0000000000004
  b= 0.0000000000002
  c= 0.0000000000007

  fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b,c))
}
```

- operationErr3.go


```go

package main

import (
  "fmt"
  "math/big"
)

func main() {
  a, _ :=  new(big.Float).SetString("0.1") // a
  b, _ :=  new(big.Float).SetString("0.2") // b
  c, _ :=  new(big.Float).SetString("0.3") // c

  d := new(big.Float).Add(a,b) // a+b
  fmt.Println(a, b, c, d)
  fmt.Println(c.Cmp(d)) // 0: equals (정밀도 높음)
}
```

- 논리 연산자
  - `|| && !`


- 대입연산자
  - `a,b =3,4`
  - `a,b = b,a`
  - `a=10 b=a`

- 증감연산자
  - `a++ a-- += -=`
  - `a++` 는 값을 반환하지 않음!

- 연산자 우선순위

### 7.함수


```go
package main
import "fmt

func factorial(n int) int {

  if n <= 1  {
    return 1
  } else {
    return n * factorial(n-1)
  }
}

func main() {

  fmt.Println(factorial(5))
}
```

### 8.상수


```go
const Pig int = 0
const Cow int = 1
const Chicken int = 2

const PI = 3.14
const PI2 float64 = 3.14

// iota로 간편하게 열거값 사용 하기
// 1씩증가. 소괄호 벗어나면 다시 초기화 됨
const (
	Red   int = iota // 0
	Blue  int = iota // 1
	Green int = iota // 2
)

// const를 소괄호()로 묶고 iota 사용하면
// 0부터 1씩 차례로 증가하며 값이 초기화 됨
const (
	C1 uint = iota + 1 // 1 = 0 + 1
	C2                 // 2 = 1 + 1
	C3                 // 3 = 2 + 1
)

const (
	BitFlag1 uint = iota + 1 // 1 = 1 << 0
	BitFlag2                 // 2 = 1 << 1
	BitFlag3                 // 4 = 1 << 2
	BitFlag4                 // 8 = 1 << 3
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	}

}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)

	var a int = PI * 100
	// var b int = PI2 * 100 // ERROR!
	fmt.Println(a)
	// fmt.Println(b)

	fmt.Println(C1)
	fmt.Println(C2)
	fmt.Println(C3)

	fmt.Println(BitFlag1)
	fmt.Println(BitFlag2)
	fmt.Println(BitFlag3)
}
```


### 9.if조건문

```go
  // 초기문을 실행하고 그 결과를 확인할때 조건문 사용
  if 초기문; 조건문 {
    // 초기문의 결과를 확인하는 목적으로 활용가능
  }
```


### 10.switch문

```go
package main

import "fmt"


type ColorType int

const (
  Red ColorType = iota // 1
  Blue                 // 2
  Yellow                 // 3
  Green                 // 4
)

func SelectColor(color ColorType) {
  switch color {
    case Red:
      fmt.Println("Red!!!rrrr")
    case Blue:
      fmt.Println("Blue!!!bbb")
    case Yellow:
      fmt.Println("Yellow!!!")
    case Green:
      fmt.Println("Greeen!!!ggg")
  }
}

func main() {
  mycolor := Red
  fmt.Println(mycolor)

}

```

### 11.for문

```go
// label사용은 지양 - 혼동 올 수 있음
func find45(a int) (int, bool) {
	for b := 1; b < 10; b++ {
		if b*a == 45 {
			return b, true
		}
	}
	return 0, false
}

func main() {
	a := 1
	b := 0
	found := false

	for ; a <= 9; a++ {
		if b, found = find45(a); found {
			break
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
```


- continue, break

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		// 입력 된 값이 변수에 저장되기 전에 버퍼에 저장
		// 숫자는 number에 저장되고 
		// 숫자 + 
		fmt.Print("숫자를 입력하세요: [-1로 종료]")
		var number int
		_, err := fmt.Scanln(&number)

		if err != nil {
			fmt.Println("숫자가 아닙니다!")

			//키보드 버퍼 비움
			stdin.ReadString('\n')
			continue
		}

		fmt.Printf("입력 하신 숫자는 %d 입니다.", number)

		if number%2 == 0 {
			fmt.Printf("입력 하신 숫자는 짝수, %d, 입니다.\n", number)
		} else if number%2 == 1 {
			fmt.Printf("입력 하신 숫자는 홀수, %d, 입니다.\n", number)
		} else if number == -1 {
			fmt.Printf("for문 종료\n")
			break
		}
	}
}
```


### 12.배열 array


- array copy

```go
package main

import (
	"fmt"
)

func main() {
	a := [5]int{1,2,3,4,5}
	b := [5]int{500,400,300,200,100}

	// a
	for i,v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	// b
	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a

	// a가 복사된 b
	fmt.Println()
	for i,v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}
```


- multi-array

```go
package main

import (
  "fmt"
)

func main() {
  a := [2][5] int {
    {1,2,3,4,5},
    {6,7,8,9,10},
    // 행바뀜 있을때는 마지막에 , 필요
    // 행바뀜없는 inline 일때는 필요 없음 e.g. { {...},{.., 9,10} }
  }

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, "")
		}
		fmt.Println()
	}
}
```

### 13.구조체 struct


- 컴퓨터가 데이터에 효과적으로 접근하고자, 메모리를 일정 크기 간격으로 정렬
  - 64비트 컴퓨터의 경우 8바이트 단위로 정렬 되어 있는게 효율적
  - int32, float64필드가 있는 구조체의 경우 12바이트(4+8) 가 아닌 16바이트(8+8) 할당
  - 순서를 float64, int32로 하면 12바이트(8+4)

```go
package main

import (
  "fmt"
  "unsafe"
)


type User struct {
  A int32 // 4바이트
  B float64 // 8바이트
}

type User2 struct {
  A int8    // 1바이트
  B int     // 8바이트
  C int8    // 1바이트
  D int     // 8바이트
  E int8    // 1바이트
}

type User3 struct {
  A int8    // 1바이트
  C int8    // 1바이트
  E int8    // 1바이트
  B int     // 8바이트
  D int     // 8바이트
}

func main() {
  user := User {34, 97.5}
  // 12바이트가 아닌 16바이트 출력 : 메모리정렬
  // 변수시작주소를 변구크기의 배수로 맞춤
  fmt.Println("user size : ", unsafe.Sizeof(user))

  user2 := User2{1,2,3,4,5}
  fmt.Println("user2 size : ", unsafe.Sizeof(user2))

  user3 := User3{1,2,3,4,5}
  fmt.Println("user3 size : ", unsafe.Sizeof(user3))

}
```

### 14.포인터

```go
func main() {
  var a int = 500
  var p *int

  p = &a

  fmt.Printf("p의 값: %p\n", p) // 메모리 주솟값
  fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p) // p가 가리키는 메모리의 값 *p == a

  *p = 100 // p가 가리키는 메모리의 값 변경
  fmt.Printf("===p가 가리키는 메모리의 값 변경->100===\n")
  fmt.Printf("a의 값: %d\n", a) // a값 변화 확인 *p == a

  var p2 *int = &a
  fmt.Printf("p2도 a를 가리키는 *int 포인터 입니다: \n\tp2= %p, *p2=%d, &a=%p\n", p2,*p2, &a)
}
```


- `==` 연산으로 포인터가 같은 메모리 공간 가리키는지 확인가능


```go
func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2 : %v\n", p1==p2)
	fmt.Printf("p2 == p3 : %v\n", p2==p3)


	// 포인터의 기본값 nil
	var p *int
	if p != nil {
		fmt.Println("포인터 p는 유효한 메모리 공간을 가리킴")
	} else {
		fmt.Println("포인터 p == nil 어떤 메모리 공간도 가리키고 있지 않음")
	}

	// 포인터는 변수대입 또는 함수 인수전달 시 메모리 공간을 적게 사용하면서 사용 가능
}
```


```go
package main

import (
  "fmt"
)

type Data struct {
  value int
  data [200]int
}

func ChangeData(arg *Data) {
  // 매개'변수' arg 데이터 변경
  arg.value = 999
  arg.data[100] = 999
}

func main() {
  // 인스턴스: 메모리에 할당된 '데이터의 실체'
  // 포인터를 이용하여 인스턴스에 접근
  // Garbage Collector에 의해 일정 간격으로 메모리에서 쓸모없는 데이터 청소 (인스턴스 등)
  // 생성된 변수 (인스턴스) 메모리는 함수 종료시 사라짐
  // 1.인스턴스는 메모리에 생성된 데이터의 실체
  // 2.포인터를 이용해서 인스턴스를 가리키게 할 수 있음
  // 3.함수 호출 시 포인터 인수를 통해서 인스턴스를 입력받고, 그값을 변경할 수 있게 됨
  // 4.쓸모 없어진 인스턴스는 가비지 컬렉터가 자동으로 지워줌

  // Data 인스턴스 생성
  var data Data
  
  // p: 새로운 Data인스턴스가 생성된 것이 아닌, 기존에 있던 data 인스턴스를 가리키는 포인터 변수
  var p *Data = &data

  fmt.Printf("p.value = %d\n", p.value)
  fmt.Printf("data.value = %d\n", data.value)

  ChangeData(&data)

  // 0->999 확인
  fmt.Printf("p.value = %d\n", p.value)
  fmt.Printf("data.value = %d\n", data.value)

  // 인스턴스 하나만 생성하고 포인터 3개가 하나의 인스턴스 가리킴
  var p1 *Data = &Data{}
  var p2 *Data = p1
  var p3 *Data = p1

  fmt.Printf("주솟값1 = %p\n", p1)
  fmt.Printf("주솟값2 = %p\n", p2)
  fmt.Printf("주솟값3 = %p\n", p3)

  // 포인터 초기화 방법
  // & 사용한 초기화 방법:
  //    포인터 값을 별도의 변수로 선언하지 않고 초기화 하는 방법
  // new  사용한 초기화 방법
  //    new 초기화 방법은 내부 필드값 초기화는 안됨
  //    & 사용한 초기화 방법과 동일 &Data{} == new(Data)
  //    & 사용한 초기화 방법은 내부 필드값도 초기화 가능 p1 := &Data{3,4}
  p4 := &Data{}
  ChangeData(p4)
  p5 := new(Data)
  ChangeData(p5)

  fmt.Printf("주솟값4 = %p\n", p4)
  fmt.Printf("p4.value = %d\n", p4.value)
  fmt.Printf("p4.value = %d\n", p4.value)

  fmt.Printf("주솟값5 = %p\n", p5)
  fmt.Printf("p5.value = %d\n", p5.value)
  fmt.Printf("p5.value = %d\n", p5.value)
}
```


- 스택 메모리 : 힙메모리 영역보다 훨씬 효율적이지만, 함수 내부에서만 사용 가능
- 힙 메모리 : 함수외부로 공개되는 메모리 공간 할당 시

```
* C/C++에서는 malloc() 호출하여 힙메모리 공간 할당
* java에서는 클래스타입을 힙에, 기본타입을 스택에 할당
* Go언어는 탈출검사(escape analysis)로 어느 메모리에 할당 할 지 결정
  예를들어 퍼블릭 함수가 내부에서 생성된 객체를 return하면
  탈출검사에서 확인되기 때문에 함수가 종료 되어도 메모리에서 사라지지 않음
  (GC가 해당 객체를 분별 할 수 있음)

* 함수외부로 공개되는 인스턴스의 경우 함수가 종료 되어도 사리지지 않음
```


### 15.문자열

```go
s1 := "Hello string variable!"
s2 := `Special \t characters \t\t \n inside backticks are ignored\n
	backticks also ignores new lines
	line3
	line4
	line5`

fmt.Println(s1)
fmt.Println(s2)
```

- rune타입 : 4-byte int32타입과 별칭
  - `type rune int32`

```go
	// 문자 하나를 표혀하는데 rune 타입 사용
	// UTF-8은 1~3바이트
	// 알파벳 문자열크기 =1, 한글 문자열 크기 = 3
	var c rune = '한'

	fmt.Printf("%T\n", c) // 타입 출력 : int32
	fmt.Println(c)				// int32(rune) 형식 값 출력
	fmt.Printf("%c\n", c) // 문자출력

	// 12 = (1*3) +  (3*3)
	a := "abc가나다"
	fmt.Println(len(a))


	// string <-> []rune 타입 변환
	// 		[]rune(str)
	// 		string([]rune{...})
	str := "Hello World!"
	// []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	runes := []rune(str)

	fmt.Println(str)
	fmt.Println(runes)
	fmt.Println(string(runes))
	fmt.Println(len(runes))

	for _,v := range runes {
		fmt.Printf("타입: %T, 값: %d, 문자열: %c\n", v,v,v)
	}

	// strng <-> []byte 변환
	// 모든 문자열은 1바이트 배열로 변환 가능
```

- Compare String

```go
	s1 := "Hello"
	s2 := "Hell"
	s3 := "Hello"

	fmt.Printf("%s == %s, %v\n", s1, s2, s1 == s2)
	fmt.Printf("%s == %s, %v\n", s1, s3, s1 == s3)

	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBAD"
	str4 := "ZZZ"

	fmt.Printf("%s > %s, %v\n", str1, str2, str1 > str1)
	fmt.Printf("%s <= %s, %v\n", str3, str4, str3 <= str4)
```


- string 끼리 대입하기
  - Data, Len 같은지 확인

```go
import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
	str1 := "Hello World"
	str2 := str1


	// string -> unsafe.Pointer -> *reflect.StringHeader
	stringHeader1 := (*reflect.StringHeader) (unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader) (unsafe.Pointer(&str2))
	
	// 두 값이 같음
	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}
```


- string is immutable


```go
func main() {
	
	var str string = "Hello World"
	str= "How are you"
	// ERROR: string is immutable
	// str[2] = 'a'

	fmt.Println(str)

	var slice []byte = []byte(str)
	slice[2] = 'a'
	fmt.Printf("%s\n", slice)

	// str = string(slice)
	// fmt.Printf("%s\n", str)
}
```

### 16.패키지

- go build 로 실행 파일 만들때 go mod 필요

```go
var c rune = '한'

fmt.Printf("%T\n", c) // char타입 출력
fmt.Println(c)
fmt.Printf("%c\n", c) // 문자출력
```

```sh
go mod init [패키지명]
go mod init 'github.com/jnuho/goproject'
# Go모듈에 필요한 패키지를 찾아서 다운로드, 필요한 패키지 정보를 go.mod, go.sum 파일에 적어줌
go mod tidy
cat go.mod
```


- 패키지 내 필드 외부 공개 여부
  - const, 변수, 함수, 별칭 타입


```sh
# ch16
#   ㄴ usepkg/usepkg.go
#   ㄴ usepkg/go.mod
#   ㄴ usepkg/go.sum
#     ㄴ custompkg/custompkg.go
cd 16-package/usepkg

# 퍼블릭 github 있는 경우
# go mod init github.com/reponame/16-package/usepkg
go mod init 16-package/usepkg
go mod tidy

# ch16
#   ㄴ ex16.2
#     ㄴ ex16.2.go
#     ㄴ publicpkg/publicpkg.go
cd 16-package/ex16.2
go mod init 16-package/ex16.2
go mod tidy
```

- 패키지 초기화
  - 전역변수 초기화
  - init() 호출 (매개변수, 반환값 없는 함수)

```go
package main

import (
	"fmt"
	"math/rand"

	// 겹치는 패키지 이름 별칭으로 묶기
	"text/template"
	htemplate "html/template"

	// 패키지를 사용하지는 않지만 부과효과 떄문에 import 하는 경우
	"database/sql"

  // 밑줄 _을 이용해서 init() 함수 호출
  // 밑줄 _ 이용하여 unused 에러 방지
	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	db *sql.DB
}

func main() {

	template.New("foo").Parse(`{{define "T"}}Hello`)
	htemplate.New("foo").Parse(`{{define "T"}}Hello`)

	fmt.Println(rand.Int())
}
```

- Go 모듈만들고, 외부 패키지 활용하기

```sh
cd 16-package/usepkg
go mod init 16-package/usepkg
mkdir custompkg
cat custompkg/custompkg.go
  package custompkg

  import "fmt"

  func PrintCustom() {
    fmt.Println("This is custom package!")
  }

cat usepkg.go
  package main

  import (
    "fmt"
    "16-package/usepkg/custompkg" // 모듈 내 패키지
    "github.com/guptarohit/asciigraph" // 외부 저장소 패키지
    "github.com/tuckersGo/musthaveGo/16-package/expkg"
  )

  func main() {
    custompkg.PrintCustom()
    expkg.PrintSample()

    data := []float64{3,4,5,6,9,7,5,8,5,10,2,7,2,5,6}
    graph := asciigraph.Plot(data)
    fmt.Println(graph)
  }
```

- custompkg.go

```go
package custompkg
import "fmt"
func PrintCustom() {
	fmt.Println("This is custom package!")
}
```


- 패키지명과 패키지 외부 공개

```go
cd 16-package/ex16.2/publicpkg
cat publicpkg.go
  package publicpkg

  import "fmt"

  const (
    PI = 3.1415
    pi = 3.1415
  )

  var ScreenSize int = 1080
  var screenHeight int

  func PublicFunc() {
    const MyConst = 100
    fmt.Println("This is a public function", MyConst)
  }

  func privateFunc() {
    fmt.Println("This is a private function")
  }

  type MyInt int
  type myString string

  type MyStruct struct {
    Age int
    name string
  }

  func (m MyStruct) PublicMethod() {
    fmt.Println("This is a public method")
  }

  func (m MyStruct) privateMethod() {
    fmt.Println("This is a private method")
  }

  type myPrivateStruct struct {
    Age int
    name string
  }

  func (m myPrivateStruct) PrivateMethod() {
    fmt.Println("This is a private method")
  }

cat 16-package/ex16.2
cat ex16.2.go
  package main

  import (
    "fmt"
    "16-package/ex16.2/publicpkg"
  )

  func main() {
    fmt.Println("PI: ", publicpkg.PI)
    publicpkg.PublicFunc()

    var myint publicpkg.MyInt = 10
    fmt.Println("myint:", myint)

    var mystruct = publicpkg.MyStruct{Age: 18}
    fmt.Println("mystrcut:", mystruct)
  }
```

- 패키지 임포트하면 다음 진행
  1. 전역변수 초기화
  2. 패키지에 init() 정의되어 있다면 호출
    - 만약 어떤 패키지의 초기화 함수인 init() 함수기능만 사용하길 원할 경우 밑줄 _ 로 임포트

```go
```



### 18.슬라이스

- slice: 동적배열
  - 일반적인 배열 `var array [10]int`은 일정한 길이에서 늘어나지 않는 문제
  - 동적배열 슬라이스 : `var slice []int`

```go
// 초기화 방법
var slice1 = []int{1,2,3} //len:3 cap: 3
var slice2 = []int{1, 5:2, 10:3} // [1 0 0 0 0 2 0 0 0 0 3]
var slice3 = make([]int, 3, 5) // len: 3 cap: 5

for i := 0; i<len(slice1); i++ {
  fmt.Print(slice1[i], " ")
}
for i,v := range slice1 {
  slice[i] = v*2
}
```

- 슬라이스 동작원리
  - func(arr [5]int) 배열 파라미터로 넘기면 전체 배열 COPY : array 요소들 변경 적용 X
  - func(slice []int) 슬라이스 파라미터 포인터까지 복제되므로: array 요소들 변경 적용 됨

```go
type SliceHeader struct {
  Data uintptr  // 실제 배열을 가리키는 포인터
  Len int       // 요소개수
  Cap int       // 실제 배열의 길이
}
```


- 문제발생 CASE1

```go
package main
import "fmt"

func main() {
  slice1 := make([]int, 3, 5)
  slice2 := append(slice1, 4, 5)

  // len() cap()
  fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))
  fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))

  // 문제1: slice1변경 시 둘다 바뀜
  slice1[1] = 100
  fmt.Println("After slice1[1]=100")
  fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))
  fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))

  // 문제2: slice1 append()하면 1,2 둘다 바뀌는데,
  // len에 따라 append 위치가 다름
  slice1 = append(slice1, 500)
  fmt.Println("After append(slice1, 500)")
  fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))
  fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))
}

// slice1:  [0 0 0] 3 5
// slice2:  [0 0 0 4 5] 5 5
// After slice1[1]=100
// slice1:  [0 100 0] 3 5
// slice2:  [0 100 0 4 5] 5 5
// After slice1 = append(slice1, 500)
// slice1:  [0 100 0 500] 4 5
// slice2:  [0 100 0 500 5] 5 5
```


- 문제발생 CASE2

```go
// cap이 다 찼을때, append() 하면
// 2배만큼 len 증가
package main
import "fmt"

func main() {
  // len=3 cap=3
  slice1 := []int{1,2,3}

  // cap=len으로 슬라이스 용량이 다 찬 상태에서
  //  append() 실행 시 cap*2만큼 용량(capacity) 증가
  //  cap: 3-> 6
  //  새로운 용량 (6)의 배열이 생성되면서,
  //  slice1과 slice2는 다른배열을 가리키게됨
  slice2 := append(slice1, 4, 5)
  fmt.Println("After slice2 := append(slice1, 4, 5) ")
  fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))
  fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))

  slice1[1] = 100
  fmt.Println("After slice1[1] = 100")
  fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))
  fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))

  slice1 = append(slice1, 500)
  fmt.Println("After slice1 = append(slice1, 500)")
  fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))
  fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))
}

// After slice2 := append(slice1, 4, 5) 
// slice1:  [1 2 3] 3 3
// slice2:  [1 2 3 4 5] 5 6
// After slice1[1] = 100
// slice1:  [1 100 3] 3 3
// slice2:  [1 2 3 4 5] 5 6
// After slice1 = append(slice1, 500)
// slice1:  [1 100 3 500] 4 6
// slice2:  [1 2 3 4 5] 5 6

```

- 슬라이싱

```go
package main

import "fmt"

func main() {
  array := [5]int{1,2,3,4,5}
  slice := array[1:2]


  // 슬라이싱 cap길이 = 시작인덱스 ~끝인덱스 까지의 길이
  //   여기에서는 cap = 시작인덱스 1부터 끝까지 길이 = 4
  fmt.Println("After slice := array[1:2]")
  fmt.Println("array: ", array)
  fmt.Println("slice: ", slice, len(slice), cap(slice))

  array[1] = 100
  fmt.Println("After array[1] = 100")
  fmt.Println("array: ", array)
  fmt.Println("slice: ", slice, len(slice), cap(slice))

  slice = append(slice, 500)
  fmt.Println("After slice = append(slice, 500)")
  fmt.Println("array: ", array)
  fmt.Println("slice: ", slice, len(slice), cap(slice))


  // IF 슬라이스 append를 배열 길이 5이상으로 실행하면
  // array는 그대로, slice는 동적으로 크기 증가: cap*2 용량 증가 및 len 증가
}


// After slice := array[1:2]
// array:  [1 2 3 4 5]
// slice:  [2] 1 4
// After array[1] = 100
// array:  [1 100 3 4 5]
// slice:  [100] 1 4
// After slice = append(slice, 500)
// array:  [1 100 500 4 5]
// slice:  [100 500] 2 4
```



- 슬라이스를 슬라이싱 하기

```go
package main

import "fmt"

func main() {
  slice1 := []int{1,2,3,4,5}
  fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))

  // 중간
  slice2 := slice1[1:3]
  fmt.Println("slice1[1:3]: ", slice2, len(slice2), cap(slice2))

  // 처음부터 i까지
  slice3 := slice1[:3]
  fmt.Println("slice1[:3]: ", slice3, len(slice3), cap(slice3))

  // i부터 끝까지
  slice4 := slice1[1:]
  fmt.Println("slice1[1:]: ", slice4, len(slice4), cap(slice4))

  // 전체
  slice5 := slice1[:]
  fmt.Println("slice1[:]", slice5, len(slice5), cap(slice5))
}
// slice1:  [1 2 3 4 5] 5 5
// slice1[1:3]:  [2 3] 2 4
// slice1[:3]:  [1 2 3] 3 5
// slice1[1:]:  [2 3 4 5] 4 4
// slice1[:] [1 2 3 4 5] 5 5
```




- 슬라이스 복제 (서로 다른 배열 가리키도록)

```go
package main
import "fmt"

func main() {
  slice1 := []int{1,2,3,4,5}

  // 복제
  slice2 := make([]int, len(slice1))
  for i,v := range slice1 {
    slice2[i] = v
  }
  fmt.Println("After slice1-> slice2 복제(for문)")
  fmt.Println("slice1", slice1, len(slice1), cap(slice1))
  fmt.Println("slice2", slice2, len(slice2), cap(slice2))

  // slice1,2는 서로 다른 배열 가리키므로 slice1의 배열만 반영
  slice1[1] = 100
  fmt.Println("After slice1[1] = 100")
  fmt.Println("slice1", slice1, len(slice1), cap(slice1))
  fmt.Println("slice2", slice2, len(slice2), cap(slice2))
}

// After slice1-> slice2 복제(for문)
// slice1 [1 2 3 4 5] 5 5
// slice2 [1 2 3 4 5] 5 5
// After slice1[1] = 100
// slice1 [1 100 3 4 5] 5 5
// slice2 [1 2 3 4 5] 5 5
```

- 슬라이스 복제 (append()로 코드개선)

```go
package main
import "fmt"

func main() {
  slice1 := []int{1,2,3,4,5}
  slice2 := append([]int{}, slice1...)
  fmt.Println("After slice2 := append([]int{}, slice1...)")
  fmt.Println("slice1", slice1, len(slice1), cap(slice1))
  fmt.Println("slice2", slice2, len(slice2), cap(slice2))
}
```




- 슬라이스 복제 (copy()로 코드개선)
  - `func copy(dest, src []Type) int`

```go
package main
import "fmt"

func main() {
  slice1 := []int{1,2,3,4,5}
	slice2 := make([]int, 3, 10) // len:3, cap:10
	slice3 := make([]int, 10) // len:10, cap:10

  // 3개만 복사됨
	cnt1 := copy(slice2, slice1)
  // 5개 복사됨
	cnt2 := copy(slice3, slice1)

	fmt.Println("slice2: ", cnt1, slice2)
	fmt.Println("slice3: ", cnt2, slice3)


  // 같은 cap, len으로 복제하려면:
  // for문으로 실행 하거나
  // copy이용
  slice2 = make([]int, len(slice1))
  copy(slice2, slice1)
  fmt.Println("slice2: ", slice2)
}

// slice2:  3 [1 2 3]
// slice3:  5 [1 2 3 4 5 0 0 0 0 0]
```



- 슬라이스 요소 삭제
  1. 삭제한 요소 인덱스 이후를 1칸씩 당기고, 마지막 요소 슬라이싱 기능으로 제거
  2. 또는 append() 조합으로 요소삭제 후 reassign


- 방법 1. 한칸씩 당기는 작업: for문

```go
package main
import "fmt"

func main() {
  slice1 := []int{1,2,3,4,5}
  del := 2

  for i:=del; i<len(slice1)-1; i++ {
    slice1[i] = slice1[i+1]
  }
  slice1 = slice1[:len(slice1)-1]

  fmt.Printf("After slice1[%d] is deleted: %v  %d %d\n", del, slice1, len(slice1), cap(slice1))
```

- 방법 2. append() 활요

```go
package main
import "fmt"

func main() {
  slice1 := []int{1,2,3,4,5}
  del := 2
  
  slice1 = append(slice1[:del], slice1[del+1:]...)

  fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))
}
```


- 슬라이스 요소 추가
  1. 추가한 요소 인덱스 이후를 1칸씩 밀기: for문
  2. 또는 append() 조합으로 요소추가 후 reassign


- 방법 1. 한칸씩 미는 작업: for문
  - len+1이 되지만 cap*2 가 됨!

```go
package main
import "fmt"
func main() {
  slice1 := []int{1,2,3,4,5}
  add := 2
  slice1 = append(slice1, 0)


  for i:= len(slice1)-1 ; i > add; i-- {
    slice1[i] = slice1[i-1]
  }
  slice1[add] = 100
  slice1 = slice1[:len(slice1)]

  // slice1 [1 2 100 3 4 5] 6 10
  fmt.Println("slice1", slice1, len(slice1), cap(slice1))
}
```

- 방법 2. append() 활용

```go
package main
import "fmt"
func main() {
  slice1 := []int{1,2,3,4,5}
  add := 2
  slice1 = append(slice1[:add], append([]int{100}, slice1[add:]...)...)

  // slice1 [1 2 100 3 4 5] 6 10
  fmt.Println("slice1", slice1, len(slice1), cap(slice1))
}
```



- 슬라이스 요소 (int형) 정렬

```go
package main
import (
  "fmt"
  "sort"
)
func main() {
  slice := []int{5, 1, 2, 3, 4}

  sort.Ints(slice)
  fmt.Println(slice)
}
```

- 구조체 타입 슬라이스 정렬
  - 구조체를 Sort()로 정렬 하려면 '구조체 리스트' 대해 Len(), Less(), Swap() 메소드 구현필요
  - 구조체리스트 = `[]구조체`

```go
package main
import (
  "fmt"
  "sort"
)

type Student struct {
  Name string
  Age int
}

type Students []Student

func (s Students) Len() int {
  return len(s)
}

func (s Students) Less(i, j int) bool {
  return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}

// 구조체 정렬 시 필요한 정의 메소드
// Len(), Less(), Swap()
func main() {
  // 슬라이스안에 {"A", 1} 선언시
  // 구조체 type 명 생략
  slice := []Student {
    {"B", 2},
    {"C", 9},
    {"A", 1},
    {"D", 5},
  }

  sort.Sort(Students(slice))
  fmt.Println(slice)
}
```


### 19.메소드

- `func (r Rabbit) info() int`
  - 리시버: `(r Rabbit)`
  - 메소드명: `info`

- 메소드와 함수의 구분: '소속'
  - 함수는 어디에도 속하지 않지만, 메소드는 리시버에 속함
  - 리시버를 통해 해당 메소드의 소속 명시
  - 메소드를 통해 object와 기능을 묶을 수 있음
- 응집도 높힘
- 객체지향 (절차중심 -> 관계중심)
  - 클래스 상속은 지원 X, 메소드와 인터페이스 O

```go
package main
import "fmt"

type account struct {
  balance int
}

func withdrawFunc(a* account, amount int) { // 일반함수 표현
  a.balance -= amount
}

func (a *account) withdrawMethod(amount int) { // 메소드 표현
  a.balance -= amount
}

func main() {
  a := &account{100}

  withdrawFunc(a, 30)
  a.withdrawMethod(30)

  fmt.Printf("%d \n", a.balance)
}
```


- 별칭 리시버 타입
  - int 타입도 별칭을 통해 리시버로 기능할 수 있음

```go
package main
import "fmt"

// 사용자 정의 별칭 타입
type myInt int

func (a myInt) add(b int) int {
  return int(a) + b
}

func main() {
  var a myInt = 10 // myInt타입변수
  res := a.add(30)
  fmt.Println("myInt + 30 = ", res)

  var b int = 20
  myInt(b)
}
```


- 객체지향: 절차중심 -> 관계중심

```go
type Student struct {
  FirstName string
  LastName string
  Age int
}

// Student 구조체에 속하는 메소드들
func (s *Student) EnrollClass(c *Subject) {
  // ...
}

func (s *Student) SendReport(p *Professor, r *Report) {
  // ...
}
```

- 포인터 메소드 vs 값타입 메소드

```go
package main
import "fmt"

type account struct {
  balance int
  firstName string
  lastName string
}

// 1. 포인터타입 메소드
func (a1 *account) withdrawPointer(amount int) {
  a1.balance -= amount
}

// 2. 값타입 메소드
func (a2 account) withdrawValue(amount int) {
  a2.balance -= amount
}

// 3. 변경된 값을 반환하는 값 타입 메소드
//  값타입 메소드를 보완하지만, 메모리의 복사가 여러번 일어나 메모리 비효율적
func (a3 account) withdrawReturnValue(amount int) account {
  a3.balance -= amount
  return a3
}

func main() {
  var mainA *account = &account{100, "Joe", "Park"}

  mainA.withdrawPointer(30)
  fmt.Println(mainA)

  // 고언어에서는 (*mainA).withdrawValue(30) 와 같음
  // 복사가 일어나므로 mainA에 변경값 반영안됨
  mainA.withdrawValue(30)
  fmt.Println(mainA)

  mainB := mainA.withdrawReturnValue(30)
  fmt.Println(mainB)
}
```



### 20.인터페이스

1. 메소드는 반드시 메소드명이 있어야 함
2. 매개변수와 반환이 다르더라도 이름이 같은 메소드 있을 수 없음
3. 인터페이스에서는 메소드 구현 포함 X

- 추상화
  - 내부동작을 감춰서 서비스를 제공하는 쪽과 사용하는 쪽 모두에게 자율르 주는 방식
- 덕테이핑 - 서비스 사용자 중심 코딩
  - 구조체 타입이 인터페이스를 구현한다는걸 명시할 필요 없음

1. 포함된 인터페이스
2. 빈 인터페이스
3. 인터페이스 기본값


```go
package main
import "fmt"

type Stringer interface {
  String() string
}

type Student struct {
  Name string
  Age int
}


func (s *Student) String() string {
  return fmt.Sprintf("안녕, 나는 %d살, %s라고 해\n", s.Age, s.Name)
}

func main() {
  // s := &Student{"Jake", 9}
  s := Student{"Jake", 9}

  res := s.String()
  fmt.Println(res)
}
```


- 우편업체 fedex.FedexSender, koreaPost.PostSender

1. FedexSender

```go
// github.com/tuckersGo/musthaveGo/ch20/fedex
package fedex
import "fmt"

type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
  fmt.Printf("Fedex에서 %s를 보냅니다.\n", parcel)
}
```

2. PostSender

```go
// github.com/tuckersGo/musthaveGo/ch20/koreaPost
package koreaPost
import "fmt"

type PostSender struct{
}

func (p *PostSender) Send(parcel string) {
  fmt.Printf("Post우편에서 %s를 보냅니다.\n", parcel)
}
```


- 인터페이스 없이 사용

```go
package main

import (
  "fmt"
  "github.com/tuckersGo/musthaveGo/ch20/fedex"
  "github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

func SendFedexParcel(parcel string, f *fedex.FedexSender) {
  f.Send(parcel)
}

func SendKoreaPostParcel(parcel string, p *koreaPost.PostSender) {
  p.Send(parcel)
}

func main() {
  f := &fedex.FedexSender{}
  p := &koreaPost.PostSender{}

  // f.Send("fedex parcel")
	// p.Send("koreaPost parcel")
  SendFedexParcel("fedex parcel", f)
  SendKoreaPostParcel("koreaPost parcel", p)
}

```

- 인터페이스 사용
  - 두개 struct 타입이 기능을 공유하되 implemntation이 조금 다름


```go
package main
import (
  "fmt"
  "github.com/tuckersGo/musthaveGo/ch20/fedex"
  "github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

type Sender interface {
  Send(parcel string)
}

func SendBook(parcel string, sender Sender) {
  sender.Send(parcel)
}

func main() {
  f := &fedex.FedexSender{}
  p := &koreaPost.PostSender{}

  SendBook("fedex소포", f)
  SendBook("koreaPost소포", p)
}


- 덕타이핑
  - FedexSender, PostSender에 implements Sender라는 명시를 하지 않아도
  - Sender인터페이스 의 Send() 메소드만 정의했다면 자동으로 duckTyping 됨

- 서비스 사용자 중심 코딩
  - Sender 인터페이스를 서비스 제공자인 Fedex나 Post가 아닌 패키지를 이용하는 쪽에서 만듦
  - 덕타이핑에서는 인터페이스 구현 여부를 타입선언에서 하는게 아니라, 인터페이스가
  - 사용될 때, 해당 타입이 인터페이스에 정의된 메서드를 포함했는지 여부로 결정
  - 서비스 제공자가 인터페이스를 정의할 필요없이, 구체적인 객체만 제공하고,
  - 서비스 이용자가 필요에 따라 그때그떄 인터페이스를 정의해서 사용할 수 있음
```

```go
// A회사: B,C회사의 제품 성능 비교하고자 함
//   A회사가 직접 Database 인터페이스 정의하여
//   TotalTime 함수 사용하도록 구현
// 구조체 BDatabase, CDatabase가 달라서
// 한 함수의 인수로 쓸수 없기때문에, Database 인터페이스 정의
func TotalTime(db Database) int {
  db.Get()
  db.Set()
  return ?
}

func Compare() {
  BDB := &BDatabase{}
  CDB := &CDatabase{}

  if TotalTime(BDB) < TotalTime(CDB) {
    fmt.Println("B회사 제품이 더 빠름")
  } else {
    fmt.Println("C회사 제품이 더 빠름")
  }
}

// 덕타이핑을 지원하지 않는다면
//    B 인터페이스 정의, C 인터페이스 정의하여
//    고객에게 알려줘야하는 불편함
//    또한, C에게 B인터페이스를 지원하도록 요청 해야함
//  덕타이핑을 활용하면, 인터페이스 지원여부를 사용하는 쪽에서 판단
```


1. 포함된 인터페이스


```go
package main

type Reader interface {
  Read()(n int, err error)
  Close() error
}

type Writer interface {
  Write()(n int, err error)
  Close() error
}

type ReadWriter interface {
  Reader
  Writer
}
// 1 Read() Write() Close() 포함 타입
//    => Reader, Writer, ReadWriter 모두 사용가능
// 2 Read() Close() 포함 타입
//    => Reader만 사용 가능
// 3 Write Close() 포함 타입
//    => Writer만 사용 가능
// 4 Read() Write() 포함 타입
//    => Close()없기 떄문에 3개 interface 모두 사용 불가능
```

2. 빈 인터페이스

- `interface{}`는 메서도를 가지고 있지 않은 빈 인터페이스
  - 가지고 있어야할 메서드가 하나도 없기 때문에 모든 타입이 빈인터페이스로 쓰일 수 있음
  - 어떤값이든 받을 수 있는 함수, 메소드, 변숫값을 만들 떄 사용



- 빈 인터페이스를 인수로 받기

```go
package main
import "fmt"


type Student struct{
  Age int
}

func PrintVal(v interface{}) {

  switch t:= v.(type) {
    case int:
      fmt.Printf("v is int %d\n", int(t))
    case float64:
      fmt.Printf("v is float64 %f\n", float64(t))
    case string:
      fmt.Printf("v is string %s\n", string(t))
    case []int:
      fmt.Printf("v is []int Slice %T:%v", t, t)
    case [5]int:
      fmt.Printf("v is [5]int Array %T:%v", t, t)
    default:
      fmt.Printf("Not supported type %T:%v\n", t, t)
  }
}

func main() {
  PrintVal(10)
  PrintVal(3.14)
  PrintVal("Hello")
  PrintVal(Student{15})

  arr := [5]int{1,2,3,4,5}
  slice := arr[2:]
  PrintVal(arr)
  PrintVal(slice)
}
```


3. 인터페이스 기본 값 nil

```go
package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker // 인터페이스 기본값은 nil입니다.
	att.Attack()     // att가 nil이기 때문에 런타임 에러가 발생합니다.
}
```



- 인터페이스 변환 하기

1. 구체화된 다른 타입으로 타입 변환
  - 인터페이스를 본래의 구체화된 타입으로 복원할 때
  - 인터페이스 변수 a 를 ConcreteType 타입으로 변환하여, ConcreteType 변수 t 생성하여 대입

```go
var a Interface
t := a.(ConcreteType)
```

```go
package main
import "fmt"

type Stringer interface {
  String() string
}

// Student 구조체가  Stringer 인터페이스 구현
//    String()메소드 정의되어 있기 때문
type Student struct {
  Age int
}

func (s *Student) String() string {
  return fmt.Sprintf("학생 나이는: %d\n", s.Age)
}

// 인터페이스 -> 구현체 타입으로 변환
// Stringer -> *Student
func PrintAge(stringer Stringer) {
  // 특정 구현체 타입으로 변환
  // PrintAge의 매개변수로 어떤 타입이 들어온지 알수 없기 때문에 
  s := stringer.(*Student)
  fmt.Printf("Age: %d\n", s.Age)
  // fmt.Println(stringer.String())
}

func main() {
  s := &Student{100}
  PrintAge(s)
}
```

- 컴파일 에러
  - 인터페이스 변수를 구체화된 타입으로 변환 하려면
  - 해당 타입이 인터페이스 메서드 집합을 포함하고 있어야함, 그렇지 않은 경우 컴파일 에러

```go
type Stringer interface {
  String() string
}

type Student struct {
}

func main() {
  var Stringer stringer
  stringer.(*Student) // 컴파일 에러
}

```


- 런타임 에러
  - go build는 되지만, 실행 중 에러 발생

```go
package main
import "fmt"

type Stringer interface {
  String() string
}

type Student struct {
}

func (s *Student) String() string {
  return "Student"
}

type Actor struct {
}

func (a *Actor) String() string {
  return "Actor"
}

func ConvertType(stringer Stringer) {
  student := stringer.(*Student)
  fmt.Println(student)
}

func main() {
  actor := &Actor{}

  // 런타임 에러:
  // Actor -> Student 타입 에러!!!
  //  panic: interface conversion: main.Stringer is *main.Actor, not *main.Student
  ConvertType(actor)
}
```

2. 다른 인터페이스로 변환
  - 구체화된 타입 뿐아니라, 다른 인터페이스로도 변환 가능
  - `AInterface -> ConcreteType <- BInterface`

```go
var AInterface = ConcreteType{}
b := a.(BInterface)
```


```go
package main
import "fmt"

type Reader interface {
  Read()
}

type Closer interface {
  Close()
}

type struct File {
}

func (*f File) Read() {
}

func ReadFile(reader Reader) {
  c := reader.(Closer)
  // ERROR! 발생!
  // reader 인터페이스변수는 File타입을 가리키고 있고,
  // File 타입은 Close() 메서드를 포함하고 있지 않기 때문에 에러발생
  c.Close()
}

func main() {
  file := &File{}

  // ERROR!
  ReadFile(file)
}
```

- 타입 변환 성공 여부 반환

```go
package main
import (
  "fmt"
  "os"
)

type Reader interface {
  Read()
}

type Closer interface {
  Close()
}

type File struct {
}

func (f *File) Read() {
}

func ReadFile(reader Reader {
  if c, ok := reader.(Closer); ok {
    c.Close()
  }
}

func main() {
  f := &File{}
  ReadFile(f)
}
```


### 21.함수고급편

- 매개변수 `...` 키워드 사용


```go
package main
import "fmt"

func sum(nums ...int) int {
  sum := 0

  fmt.Printf("nums 타입: %T\n", nums)
  for _,v := range nums {
    sum += v
  }
  return sum
}

func PrintT(args ...interface{}) {
  for _,arg := range args {
    switch f := arg.(type) {
      case int:
        val := arg.(int)
        fmt.Println(val, int(f))
        fmt.Println(val == int(f))
        // Print로직
      case float64:
        val := arg.(float64)
        fmt.Println(val == float64(f))
      case string:
        val := arg.(string)
        fmt.Println(val == string(f))
    }
  }
}

func main() {
  fmt.Println(sum(1,2,3,4,5))
  fmt.Println(sum(10,20))
  fmt.Println(sum())

  PrintT(1,3.14, "abc")

}

```


- `defer` 지연 실행

- 함수타입 변수
  - 함수 시작지점을 가리키는 프로그램 카운터 있음 e.g. 1번라인, 2번라인...
  - 함수 시작지점은 숫자로 표현되며, 가리키는 값을 함수포인터라고 함
  - 함수 타입은 함수정의로 표현 e.g. `func (int, int) int`
  - 함수 alias 지정가능 `type opFunc func (int, int) int`
  - 함수리터럴 (익명함수, lambda)
    - 함수리터럴 외부변수를 자동으로 리터럴 함수 내부상태로 가져와 저장 ('capture') : 값복사가 아닌 참조형태로 가져옴

```go
package main
import (
  "fmt"
  "os"
)

func main() {
  f,err := os.Create("test.txt")
  if err != nil {
    fmt.Println("Failed to Create a file")
    return
  }

  defer fmt.Println("반드시 호출됩니다.") //4
  defer f.Close() //3
  defer fmt.Println("파일을 닫았습니다.") //2


  fmt.Println("파일에 Hello World를 씁니다.") // 1
  fmt.Fprintln(f, "Hello World!")
}


// 1->2->3->4
// 파일에 Hello World를 씁니다.
// 파일을 닫았습니다.
// 반드시 호출됩니다.
```

- 함수 타입 변수

```go
package main

import "fmt"

func add(a, b int) int {
  return a + b
}

func mul(a, b int) int {
  return a * b
}

type opFunc func(int, int) int

func getOperator(op string) opFunc {
  if op == "+" {
    return add
  } else if op == "*" {
    return mul
  } else {
    return nil
  }
}

func main() {
  // var opFnc func(int,int) int = getOperator("*")
  opFnc := getOperator("*")

  res := opFnc(2, 9)
  fmt.Println("2*9 = ", res)
}

```

- 함수 리터럴
  - 다른 언어 에서는 lambda(익명 함수)

```go
package main
import "fmt"

func getOperator(operator string) func(int, int) int {
  if operator == "*" {
    return func(a, b int) int {
      return a * b
    }
  } else if operator == "+" {
    return func(a, b int) int {
      return a + b
    }
  } else {
    return nil
  }
}

func main() {
  fn := getOperator("*")
  a, b := 3,4
  fmt.Printf("%d x %d = %d\n", a,b, fn(a,b))
}
```


- 함수리터럴 내부 상태
  - 함수 내부에서 사용 되는 외부 변수는 자동으로 함수 내부 상태로 저장 됨

```go
package main
import "fmt"
func main() {
  i := 0

  // 함수
  f := func() {
    i += 10
  }

  i++
  f()
  fmt.Println(i)
}
```


- 함수 리터럴 내부상태 주의점
  - 캡쳐: 함수 리터럴 외부 변수를 내부로 가져옴
  - 다만 캡쳐는 복사가 아닌 참조 형태로 가져옴



```go
package main
import "fmt"

func CaptureLoop() {
  // make slice of functions with length 3
  f := make([]func(), 3)
  fmt.Println("CaptureLoop")

	// i 가 복사되는것이 아닌 참조되기 떄문에
  for i :=0; i< len(f); i++ {
    f[i] = func() {
			// 캡쳐할 떄 캡쳐하는 순간의 i값(1,2,3)이
			// 복제 되는 것이 아니라, 변수가 참조로 캡쳐되므로
			// i가 최종적으로 3이 되었을떄
			// i를 참조하는 f[0], f[1], f[2]는 모두 i=3를 참조하게 됨
      fmt.Println(i)
    }
  }

  for i :=0; i< len(f); i++ {
		f[i]()
	}
}


func CaptureLoop2() {
  // make slice of functions with length 3
  f := make([]func(), 3)
  fmt.Println("CaptureLoop2")

	// i 가 복사되는것이 아닌 참조되기 떄문에
  for i :=0; i< len(f); i++ {
		v := i
    f[i] = func() {
			// 캡쳐할 떄 캡쳐하는 순간의 i값(1,2,3)이
			// 복제 되는 것이 아니라, 변수가 참조로 캡쳐되므로
			// i가 최종적으로 3이 되었을떄
			// i를 참조하는 f[0], f[1], f[2]는 모두 i=3를 참조하게 됨
      fmt.Println(v)
    }
  }

  for i :=0; i< len(f); i++ {
		f[i]()
	}
}

func main() {
	// 3 3 3
  CaptureLoop()
	// 1 2 3
  CaptureLoop2()
}
```


- 파일 핸들을 내부 상태로 사용하는 예

```go
package main
import (
  "os"
  "fmt"
)

type Writer func(string)

func writeHello(writer Writer) {
  writer("Hello World")
}

func main() {
  f, err := os.Create("test.txt")
  if err != nil {
    fmt.Println("Failed to create a file")
    return
  }

  defer f.Close()

  writeHello(func (msg string) {
    // 함수리터럴 외부 변수 f사용
    fmt.Fprintf(f, msg)
  })
}
```

### 22.자료구조

- 리스트
  - 리스트 구현하는 구조체 코드

```go
type Element struct {
  Value interface{}
  Next *Element
  Prev *Element
}
```

- 리스트 라이브러리

```go
// Element
func (e *Element) Next() *Element
func (e *Element) Prev() *Element

// List
// Package list implements a doubly linked list.
func New() *List
func (l *List) Back() *Element
func (l *List) Front() *Element
func (l *List) Init() *List
func (l *List) InsertAfter(v any, mark *Element) *Element
func (l *List) InsertBefore(v any, mark *Element) *Element
func (l *List) Len() int
func (l *List) MoveAfter(e, mark *Element)
func (l *List) MoveBefore(e, mark *Element)
func (l *List) MoveToBack(e *Element)
func (l *List) MoveToFront(e *Element)
func (l *List) PushBack(v any) *Element
func (l *List) PushBackList(other *List)
func (l *List) PushFront(v any) *Element
func (l *List) PushFrontList(other *List)
func (l *List) Remove(e *Element) any
```

- 리스트 기본 사용법

```go
package main
import (
  "container/list"
  "fmt"
)

func main() {
  // PushBack, PushFront returns Pointer to newly insertedElement
  // 1 4
  //    1 <- e1
  //    4 <- e4
  v := list.New()
  e4 := v.PushBack(4)
  e1 := v.PushFront(1)

  v.InsertBefore(3, e4)
  v.InsertAfter(2, e1)

  for e := v.Front(); e != nil; e= e.Next() {
    fmt.Print(e.Value, " ")
  }
  fmt.Println()

  // 역순
  for e := v.Back(); e != nil; e = e.Prev() {
    fmt.Print(e.Value, " ")
  }
}
```

- 배열,슬라이스 vs. 리스트
  - 요소삽입 : O(N) vs. O(1)
  - 요소삭제 : O(N) vs. O(N)
  - 요소접근 : O(1) vs. O(N)
    - 배열시작주소 + (인덱스*타입크기)

- 큐
  - First In First Out

```go
package main

import (
  "fmt"
  "container/list"
)

// list.List : 링크드리스트
//    .PushBack(n)
//    .Front()
//        front = .Front()
//    .Remove(front)
type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{ list.New()}
}

func main() {
	queue := NewQueue()

	for i :=1; i<=4; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for v != nil {
		fmt.Print(v, " ")
		v = queue.Pop()
	}
	fmt.Println()
}
```

- 스택
  - First In Last Out

```go
package main

import (
  "fmt"
  "container/list"
)

type Stack struct {
  v *list.List
}

func (s *Stack) Push(val interface{}) {
  s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
  back := s.v.Back()
  if back !=nil {
    return s.v.Remove(back)
  }
  return nil
}

func NewStack() *Stack {
  return &Stack{ list.New() }
}

func main() {
  stack := NewStack()
  for i:=1; i<=4; i++ {
    stack.Push(i)
  }

  val := stack.Pop()
  for val != nil {
    fmt.Print(val , " ")
    val = stack.Pop()
  }
}
```

- 링

```go
func New(n int) *Ring
func (r *Ring) Do(f func(any))
func (r *Ring) Len() int
func (r *Ring) Link(s *Ring) *Ring
func (r *Ring) Move(n int) *Ring
func (r *Ring) Next() *Ring
func (r *Ring) Prev() *Ring
func (r *Ring) Unlink(n int) *Ring
```

```go
package main

import (
  "fmt"
  "container/ring"
)

func main() {
  r := ring.New(5)
  n := r.Len()

  for i:=0; i< n; i++ {
    r.Value = 'a' + i
    r = r.Next()
  }

  for i :=0; i<n; i++ {
    fmt.Printf("%c", r.Value)
    r = r.Next()
  }
  fmt.Println()

  for i :=0; i<n; i++ {
    fmt.Printf("%c", r.Value)
    r = r.Prev()
  }
}
```

- 맵
  - 같은입력이 들어오면 같은값이 나옴
  - 입력범위는 무한, 결과는 특정범위

```go
package main
import "fmt"

func main() {
  m := make(map[string]string)
  m["이하나"] = "A100"
  m["송하나"] = "A101"
  m["박하나"] = "A102"
  m["최하나"] = "A103"

  s := "송하나"
  fmt.Print(s, ": ")
  fmt.Println(m[s])
}
```

- 요소 삭제와 없는 요소 확인

```go
package main
import "fmt"
func main() {
  m := make(map[int]int)
  m[1] = 0
  m[2] = 1
  m[3] = 2

  // delete(맵, key값)
  delete(m, 3)
  delete(m, 4) // 없는요소 삭제시도

  v, ok := m[3]
  fmt.Println(v, ok) // 삭제된 요소값 출력

  v, ok = m[1]
  fmt.Println(v, ok) // 존재하는 요소값 출력


  // 맵 순환하기
  for k, v := range m {
    fmt.Printf("Key: %d, Value: %d\n", k, v)
  }
}
```

- 배열,슬라이스 vs. 리스트 vs. 맵
  - 삽입 : O(N) vs. O(1) vs. O(1)
  - 삭제 : O(N) vs. O(N) vs. O(1)
  - 접근 : O(1) vs. O(N) vs. O(1) 키로 접근
    - 배열시작주소 + (인덱스*타입크기)

- 맵의 원리 - 해시함수
  1. 같은 입력 -> 같은 결과
  2. 다른 입력 -> 다른 결과
  3. 입력값의 범위는 무한대, 결과는 특정 범위
    - 배열로 구현시 같은 해시값(인덱스)에 1개의 값만 맵핑
    - m[hash(23)] m[hash(33)] 해시충돌
      - 이 문제는 배열안에 리스트 저장하여 해결가능
    - 해시함수는 요소개수와 상관없이 고정된 시간을 갖는 함수이므로 맵이 읽기,쓰기에서 O(1)의 시간


- 해시로 맵 만들기

```go
const M = 10

func hash(d int) int {
  return d % M
}

func main() {
  var m [M]int
  
  m[hash(23)] = 10
}
```


```go
package main
import "fmt"

const M = 10

func hash(d int) int {
  return d % M
}

func main() {
  m := [M]int{}
  m[hash(23)] = 10
  m[hash(33)] = 10 // 덮어 씌워짐 (해시충돌)
  m[hash(259)] = 50
  fmt.Printf("m[hash(23)] = %d\n", m[hash(23)])
  fmt.Printf("m[hash(259)] = %d\n", m[hash(259)])
}
```


### 23.에러 핸들링

```go
package main

import (
  "fmt"
  "os"
  "bufio"
)

func ReadFile(filename string) (string, error) {
  file, err := os.Open(filename)
  if err != nil {
    return "", err
  }

  defer file.Close()

  rd := bufio.NewReader(file)
  line, _ := rd.ReadString('\n')
  return line, nil
}

func WriteFile(filename string, line string) error {
}

func main() {
}
```

### 24.고루틴과 동시성 프로그래밍

- 스레드란?
  - 고루틴: 경량 스레드로 함수나 명령을 동시에 실행 시 사용. main()도 고루틴에 의해 실행 됨
  - 프로세스 1개(싱글테스킹) vs 멀티프로세스 (멀티태스킹).
  - 프로세스 : 메모리 공간에 로딩되어 동작하는 프로그램. 1개 이상의 스레드를 가지고 있음
    - 스레드는 프로세스의 세부작업 단위 또는 실행 흐름. 1개, 2개, 멀티 스레드 있을 수 있음
  - CPU는 하나의 스레드 실행가능, 여러개 스레드를 번갈아가면서 수행하여 동시실행 처럼 보임

- 컨텍스트 스위칭 비용
  - 한번에 여러개 일 수행 시 스레드는 컨텍스트를 통해 저장 (명령 포인터, 스택메모리 등 저장)
  - 컨텍스트 저장 및 복원 시 비용 발생
  - Go언어는 한개의 CPU 코어당 하나의 OS스레드 할당 하므로 컨텍스트 비용 발생 없음

- 고루틴 사용
  - 모든 프로그램은 main()을 고루틴으로 하나 가지고 있음
  - `go 함수호출`로 새로운 고루틴 추가 가능. 현재 고루틴이 아닌 새로운 고루틴에서 함수가 수행 됨



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

- 모든 고루틴 작업 완료할때 까지 대기
  - 항상 고루틴의 종료시간에 맞춰 time.Sleep(종료까지걸리는시간) 호출할 수 없음

```go
// sync.WaitGroup 객체 사용
var wg sync.WaitGroup

// Add()로 완료해야 하는 작업개수 설정하고, 각 작업이 완료 될때마다 Done() 호출하여
// 남은 작업개수를 하나씩 중여줌. Wait()은 전체 작업이 모두 완료될때까지 대기하기 됨
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
  wg.Add(10) // 총 작업개수 설정

  for i:=0; i<10; i++ {
    go SumAtoB(1, 1000000)
  }

  wg.Wait() // 모든 작업이 완료 될때까지 (남은작업개수 = 0) 종료하지 않고 대기
  fmt.Println("모든 계산이 완료됐습니다.")
}
```

- 고루틴의 동작방법
  - 고루틴은 OS 스레드를 이용하는 경량 스레드
  - 2코어 시스템, 고루틴 3개일때 작업끝날때 까지 대기 후 끝난 1코어에 고루틴_3 배정됨
  - 시스템 콜 호출 시 (고루틴으로 시스템콜 호출시; e.g. 네트워크로 데이터 읽을 때 데이터 들어올때 까지 고루틴이 대기상태 됨), 
    - 네트워크 수신 대기상태인 고루틴이 대기목록으로 빠지고, 대기중이던 다른 고루틴이 OS 스레드를 이용하여 실행 됨
    - 코어와 스레드 변경(컨텍스트 스위칭) 없이 고루틴이 옮겨다니기 때문에 효율적


- 동시성 프로그래밍 주의점
  - 하나의 자원의 여러개 고루틴 접근!

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

- 뮤텍스를 이용한 동시성 문제 해결

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
  mutex.Lock()
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

- 뮤텍스와 데드락
  - 멀티코어 환경에서는 여러 고루틴으로 성능 향상 가능
  - 같은메모리 접근시 꼬일 수 있음
  - 뮤텍스로 고루틴 하나만 접근하도록 하여 꼬이는 문제 해결 가능
  - 하지만, 뮤텍스를 잘못 사용하면 성능향상 없이 데드락 발생가능
    - 뮤텍스 사용시 좁은 범위에서 사용하여 데드락 발생 방지
    - 또는 둘다 수저-> 포크 순서로 뮤텍스 락 사용하면 해결 가능

```go
package main
import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)

var wg sync.WaitGroup

func diningProblem(
  name string
  , first, second *sync.Mutex,
  , firstName, secondName string) {

  for i:= 0; i<100; i++ {
    fmt.Printf("%s 밥을 먹으로 합니다.\n", name)
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

- 또 다른 자원 관리 기법
  - 영역을 나누는 방법
  - 역할을 나누는 방법

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
  fmt.Printf("%d 작업 완료 - 작업결괴: %d\n", j.index, j.index * j.index)
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

### 25.채널과 컨텍스트


- 채널: 고루틴끼리 메시지를 전달 할 수 있는 메시지 큐
  - 메시지큐에 메시지가 쌓이게 되고 메시지를 읽을 때는 처음온 메시지부터 차례대로 읽음

- 채널 인스턴스 생성

```go
// 채널타입: chan string
//    chan: 채널키워드
//    string: 메시지 타입
var messages chan string = make(chan string)
```


- 채널에 데이터 넣기

```go
var messages chan string = make(chan string)
messages <- "This is a message"
```

- 채널에서 데이터 빼기

```go
// 채널에서 빼낸 데이터를 담을 변수
// 채널 messages에 데이터가 없으면 데이터가 들어올떄까지 대기함
var msg string = <- messages
```

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

func main() {
  var wg sync.WaitGroup
  ch := make(chan int)

  wg.Add(1)
  go square(&wg, ch)
  ch <- 9

  // square내에서 채널 데이터 빼고 wg.Done() 완료 될떄까지 대기
  wg.Wait()
}
```


- 채널 크기
  - 기본 채널크기 0
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
  // 보관할 수 없으므로 채널에서 데이터 빼는 코드가 있어야 진행가능!
  // 데드락 발생! 택배 보관장소 없으면 문앞에서 기다려야 함
  // 데이터 보관할 수 있는 메모리영역: 버퍼
  ch <- 9
  fmt.Println("Never print")
}
```

- 버퍼 가진 채널
  `var chan string messages = make(chan string, 2)`

- 채널에서 데이터 대기

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

func square(wg *sync.WaitGroup, ch chan int) {
  // 채널에 데이터가 들어 올때까지 계속 기다림
  // square() 호출 밖에서 close(채널)로 채널이 닫히면
  // for문을 종료하여 프로그램 정상 종료하도록 함
  for n := range ch {
    fmt.Printf("Square: %d\n", n*n)
    time.Sleep(time.Second)
  }

  // 실행되지 않음 :
  // 위에 for문에서 계속 채널로 들어오는 데이터 기다림
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
  // for range 문을 빠져나가게 됨
  close(ch)

  wg.Wait()
}
```

- SELECT 문
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
    select {
    case n := <-ch:
      fmt.Printf("Squared: %d\n", n*n)
      time.Sleep(time.Second)
    case <- quit:
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

- 일정간격으로 실행

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

- 채널로 생산자 소비자 패턴 구현
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
// 		Body -> Tire -> Color
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

- 컨텍스트 사용하기
  - 컨텍스트는 작업을 지시할때 작업가능 시간, 작업 취소 등의 조건을 지시할 수 있는 작업명세서 역할
  - 새로운 고루틴으로 작업 시작할떄 일정시간 동안만 작업지시하거나, 외부에서 작업 취소시 사용.
  - 작업 설정에 대한 데이터도 전달 가능


- 작업취소 가능한 컨텍스트
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
	// 		컨텍스트의 Done()채널에 시그널을 보내, 작업자가 작업 취소하도록 알림
	//		<-ctx.Done() 채널
	cancel()

	wg.Wait()
}
```


- 작업시간 설정한 컨텍스트
  - 일정시간 동안만 작업을 지시할 수 있는 컨텍스트 생성
  - WithTimeout() 두번째 인수로 시간을 설정하면, 그시간이 지난 뒤
  - 컨텍스트의 Done()채널에 시그널을 보내서 작업 종료
  - WithTimeout() 역시 두번째 반환값으로 cancel함수 반환하기 때문에
  - 작업 시작전에 원하면 언제든지 작업 취소 가능

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
```

- 특정 값을 설정한 컨텍스트
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

- 취소도 되면서 값도 설정하는 컨텍스트 만들기
  - 컨텍스트를 만들때 항상 상위 컨텍스트 객체를 인수로 넣어줘야 했음
  - 일반적으로 context.Background()를 넣어줬는데, 여기에 이미 만들어진 컨텍스트 객체 넣어도 됨
  - 이를통해 여러 값을 설정하거나, 기능을 설정할 수 있음
  - 구글: "golang concurrency patterns"

```go
ctx, cancel := context.WithCancel(context.Background())
ctx = context.WithValue(ctx, "number", 9)
ctx = context.WithValue(ctx, "keyword", "Lilly")
```


### 26.단어검색 프로그램 만들기

- 와일드카드
  - * : 0개이상 문자
  - ? : 1개 문자
- `os.Args []string` 변수와 실행 인수
  - Args는 os 패키지의 전역변수로 각 실행인수가 []string 슬라이스에 담김
  - 첫번째 항목으로 실행 명령이 들어감
  - 두번쨰 부터 입력한 인수들이 차례로 들어감
    - `find word filepath`
      - os.Args[0] : find
      - os.Args[1] : word
      - os.Args[2] : filepath

- 파일 핸들링

```go
// 파일열기
func Open(name string) (*File, error)

// 파일 목록가져오기 (slice)
func Glob(pattern string) (matches []string, err error)
filepaths, err := filepath.Glob("*.txt")

// 파일 내용 한줄 씩 읽기
func NewScanner(r io.Reader) *Scanner

type Scanner
  func (s *Scanner) Scan() bool
  func (s *Scanner) Text() string
```


- 단어 포함여부 검사
  - `func Contains(s, substr string) bool`


- 실행 인수 읽고 파일목록 가져오기
  - 단어프로그램 실행파일 만들기 :
  - cd 26-search-word-project/ex26.1
  - touch ex26.1.go
  - go mod init 26-search-word-project/ex26.1
  - go mod tidy
  - go build
  - ex26.1 word ex*
  - 찾으려는 단어 : word
  - 찾으려는 파일 리스트 :
  -   ex26.1.exe
  -   ex26.1.go

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.1 word filepath")
		return
	}

	// 실행인수 가져오기
	//	찾으려는 단어
	word := os.Args[1]
	//	검색할 파일리스트 (슬라이스)
	files := os.Args[2:]
	fmt.Println("- 찾으려는 단어: ", word)
	PrintAllFiles(files)
}

// 파일리스트 및 에러 반환
func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

// 찾은 파일리스트 출력
func PrintAllFiles(files []string) {
	fmt.Println("- 찾으려는 파일 리스트")
	for _, path := range files {
		filelist, err := GetFileList(path)
		if err != nil {
			fmt.Println("파일경로가 잘못되었습니다. err:", err, "path:", path)
			return
		}

		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}
```

- 파일을 열어서 라인 읽기
  - 파일을 열고 bufio 패키지의 NewScanner()로 스캐너를 만들어, 파일내용을 한단어씩 읽기

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PrintFile("hamlet.txt")
}

// 파일을 읽어서 출력
func PrintFile(filename string) {
	file, err := os.Open(filename) // 파일 열기
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		return
	}

	defer file.Close() // 함수 종료전 파일 닫기

	scanner := bufio.NewScanner(file) // 스캐너를 생성해서 한줄 씩 읽기
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
```

- 파일 검색 프로그램 완성 하기

```go
```


### 27.객체지향원칙 SOLID

```go

```


### 28.테스트와 벤치마크


### 29.Go언어로 만드는 웹서버


### 30.Restful API 서버 만들기


### 31.TODO리스트 웹사이트 만들기


### A.Go문법 보충


### B.생각하는 프로그래밍


