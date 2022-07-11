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
      fmt.Printf("v is []int Slice %T:%v\n", t, t)
    case [5]int:
      fmt.Printf("v is [5]int Array %T:%v\n", t, t)
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