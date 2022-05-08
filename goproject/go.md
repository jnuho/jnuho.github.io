

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
byte | uint8의 별칭 1바이트 데이터 나타낼때 사용 | 0~255
rune | int32별칭 UTF-8로 문자하나 나타낼때 사용 | -214..648~214..647
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


