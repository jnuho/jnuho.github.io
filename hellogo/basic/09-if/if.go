package main



import (
  "fmt"
)

func getMyAge() (int, bool) {
  return 33, true
}

func main() {
  if age, ok := getMyAge(); ok && age < 20 {
    fmt.Println("Very Young")
  } else if ok && age >=20 && age < 30 {
    fmt.Println("Young")
  } else if ok && age >=30 && age < 40 {
    fmt.Println("Old")
  } else if ok && age >=40 {
    fmt.Println("Very Old")
  } else {
    fmt.Println("ERROR")
  }
}

