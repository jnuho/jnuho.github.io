package main

import (
  "fmt"
)

type Data struct {
  value int
  data [200]int
}

type Student struct {
  Name string
  Age int
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
  
  // 새로운 Data인스턴스가 생성된 것이 아닌, 기존에 있던 data 인스턴스를 가리킴
  var p *Data = &data

  fmt.Printf("p.value = %d\n", p.value)
  fmt.Printf("data.value = %d\n", data.value)

  ChangeData(&data)

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

  p0 := &Student{"abc", 10}
  fmt.Printf("p0.Name = %s, p0.Age= %d\n", p0.Name, p0.Age)
  fmt.Printf("%v\n", *p0)
}