package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3) // 함수리터럴 3개를 가진 슬라이스
}

func CaptureLoop2() {

}

func main() {
	CaptureLoop()
	CaptureLoop2()
}