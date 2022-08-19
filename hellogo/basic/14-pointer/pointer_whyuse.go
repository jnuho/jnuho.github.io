package main

import (
  "fmt"
)

type Data struct {
  value int
  data [200]int
}

func ChangeData(arg Data) {
  // 매개'변수' arg 데이터 변경
  arg.value = 999
  arg.data[100] = 999
}

func ChangeData2(arg *Data) {
  // 매개'변수' arg 데이터 변경
  arg.value = 999 // (*arg).value 와 같음
  arg.data[100] = 999 // (*arg).data[100] 와 같음
}

func main() {

  // 포인터는 변수대입 또는 함수 인수전달 시 메모리 공간을 적게 사용하면서 사용 가능
  var data Data

  // 1.매개변수 arg는 인자로 전달하는 data와 다른 주소값을 가지기 때문에
  // arg값을 변경 해도, data는 변경되지 않음
  // 2.함수호출할 때 마다 data사이즈 만큼 메모리에 복사됨: 성능에 안좋음
  ChangeData(data)
  fmt.Printf("value= %d\n", data.value)
  fmt.Printf("data[100]= %d\n", data.data[100])

  // 포인터 사용하여 1,2 해결
  ChangeData2(&data)
  fmt.Printf("value= %d\n", data.value)
  fmt.Printf("data[100]= %d\n", data.data[100])

  // 1.기존방식
  // var data Data
  // var p *Data = &data

  // 2.구조체를 생성해 초기화하는 방식
  // 포인터 변수 p만으로도 구조체의 필드값에 접근하고 변경 가능
  var p *Data = &Data{}
  p.value =  77
  p.data[100] =  77
  fmt.Printf("value= %d\n", p.value)
  fmt.Printf("data[100]= %d\n", p.data[100])

}