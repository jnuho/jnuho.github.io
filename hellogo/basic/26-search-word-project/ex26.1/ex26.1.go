package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.1 word filepath")
		return
	}

	// 실행인수 가져오기
	//	찾으려는 단어
	word := os.Args[1]
	//	검색할 파일리스트 (슬라이스)
	files := os.Args[2:]

	fmt.Println("- 찾으려는 단어: ", word)
	PrintAllFiles(files)
}

// 파일리스트 및 에러 반환
func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

// 찾은 파일리스트 출력
func PrintAllFiles(files []string) {
	fmt.Println("- 찾으려는 파일 리스트: ")
	for _, path := range files {
		filelist, err := GetFileList(path)
		if err != nil {
			fmt.Println("파일경로가 잘못되었습니다. err:", err, "path:", path)
			return
		}

		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}