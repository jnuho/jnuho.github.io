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