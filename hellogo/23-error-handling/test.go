package main

import "fmt"

// Error() 메소드 있는 error 인터페이스를 구현
type PasswordError struct {
  Len int
  RequireLen int
}

func (err PasswordError) Error() string {
  return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
  if len(password) < 8 {
    return PasswordError{len(password), 8}
  }
  return nil
}

func main() {
  err := RegisterAccount("myID", "myPw")
  if err != nil {
    if errInfo, ok := err.(PasswordError); ok {
      fmt.Printf("%v Len: %d RequiredLen: %d\n", errInfo, errInfo.Len, errInfo.RequireLen)
    }
  } else {
    fmt.Println("회원 가입됐습니다.")
  }

}