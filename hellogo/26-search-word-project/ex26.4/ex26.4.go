package main

import (
	"fmt"
	"os"
	"path/filepath"
	"bufio"
	"strings"
)

type LineInfo struct {
	lineNo int
	line string
}

type FindInfo struct {
	filename string
	lines []LineInfo
}

/**
cd 26-search-word-project/ex26.4
go mod init 26-search-word-project/ex26.4
go mod tidy
go build
./ex26.4 word *.txt
*/



func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}

	ch := make(chan FindInfo)
	cnt := len(filelist)
	recvCnt := 0

	for _,filename := range filelist {
		go FindWordInFile(word, filename, ch)
	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		recvCnt++
		if recvCnt == cnt {
			// all received
			break
		}
	}
	return findInfos
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		ch <- findInfo
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNo := 1
	for scanner.Scan() {

		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}

	ch <- findInfo
}


func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.1 word filepath")
		return
	}

	findInfos := []FindInfo{}
	word := os.Args[1]
	files := os.Args[2:]

	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("=====")
		for _, lineInfo := range findInfo.lines {
			fmt.Println(lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("=====")
	}
}