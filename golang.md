
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
  s := &Student{"Jake", 9}
  res := s.String()
  fmt.Println(res)
}
```

### 21.함수고급편

- `...` 키워드 사용
- `defer` 지연 실행
- 함수타입 변수
  - 함수시작 지점을 가리키는 프로그램 카운터 있음 e.g. 1번라인, 2번라인...
  - 함수시작 지점은 숫자로 표현되며, 가리키는 값을 함수포인터라고 함
  - 함수타입은 함수정의로 표현 e.g. `func (int, int) int`
  - 함수 alias 지정가능 `type opFunc func (int, int) int`
  - 함수리터럴 (익명함수, lambda)
    - 함수리터럴 외부변수를 자동으로 리터럴 함수 내부상태로 가져와 저장 ('capture') : 값복사가 아닌 참조형태로 가져옴

### 24.고루틴과 동시성 프로그래밍


```go

```
