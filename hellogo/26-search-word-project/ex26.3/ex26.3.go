package main

import (
	"fmt"
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// 찾은 라인정보
type LineInfo struct {
	lineNo int
	line string
}

// 파일 내 라인정보
type FindInfo struct {
	filename string
	lines []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행인수가 필요합니다. ex) ex26.3 word filepath")
		return
	}
}

func GetFileList(path string)([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {

}

func FindWordInFile(word, filename string) FindInfo {
}
