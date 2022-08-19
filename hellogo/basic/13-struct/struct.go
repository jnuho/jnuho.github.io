package main

import (
  "fmt"
  "unsafe"
)

// 대문자로 시작하는 필드는 외부에 공개 됨
type User struct {
  Name string // 16
  ID string // 16
  Age int32 // 4바이트
}

type VIPUser struct {
  // UserInfo User

  //포함된 필드 방식: 접근시 한단계 . 으로 접근가능 e.g. vip.Name vip.ID vip.Age
  User    // 36바이트
  VIPLevel int  // 8바이트
  Price int     // 8바이트
}

func main() {
  user := User{"송하나", "hana", 23}
  vip := VIPUser {
    User {"화랑", "hwarang", 34},
    3,
    250,
  }
  fmt.Println("user 구조체 사이즈=", unsafe.Sizeof(user))
  fmt.Println("vip 구조체 사이즈=", unsafe.Sizeof(vip))

}