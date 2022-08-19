package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PrintFile("hamlet.txt")
}

// 파일을 읽어서 출력
func PrintFile(filename string) {
	file, err := os.Open(filename) // 파일 열기
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		return
	}

	defer file.Close() // 함수 종료전 파일 닫기

	scanner := bufio.NewScanner(file) // 스캐너를 생성해서 한줄 씩 읽기
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}