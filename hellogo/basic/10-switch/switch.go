package main

import (
	"fmt"
)

func getMyAge() int {
  return 22
}

func main() {
	date := 3
	day := "thursday"
  temp := 18

	switch date {
	case 1:
		fmt.Println("첫째")
	case 2:
		fmt.Println("둘째")
	case 3:
		fmt.Println("셋째")
	default:
		fmt.Println("DEFAULT")
	}

	switch day {
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업")
	case "wednesday", "thursday", "friday":
		fmt.Println("수,목,금은 실습")
	default:
		break
	}


  fmt.Println("switch true 또는 명시하지 않으면 디폴트 true 사용")
  switch true {
  case temp < 10, temp >30:
    fmt.Println("바깥 활동하기에 좋은 날씨 아님")
  case temp >=10 && temp < 20 :
    fmt.Println("바깥 약간 쌀쌀")
    // 여기서 case문 실행 후 종료되기 때문에 다음 case문 실행 X
  case temp temp>=15 && temp < 25:
    fmt.Println("바깥 야외 활동하기 좋음")
  }


  switch age := getMyAge(); age {
  case 10:
    fmt.Println("Child")
  case 20:
    fmt.Println("Teenager")
  case 30:
    fmt.Println("20's")
  default
    fmt.Println("My Age is ", age)
  }

  // switch age := getMyAge(); true{
  switch age := getMyAge(); {
  case age < 10:
    fmt.Println("Child")
  case age < 20:
    fmt.Println("Teenager")
  case < 30:
    fmt.Println("20's")
  default
    fmt.Println("My Age is ", age)
  }

}
