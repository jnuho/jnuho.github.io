package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	filesInHomeDir, err := os.ReadDir(homeDir)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(filesInHomeDir))

	fmt.Println("Printing files in", homeDir)

	for _, file := range filesInHomeDir {
		// anon function with parameter type `os.DisEntry`
		// The `(file)` at the end of the expression immediately invokes
		// the anonymous function with the value of the file variable.
		go func(f os.DirEntry) {
			defer wg.Done()
			fmt.Println(f.Name())
		}(file)
	}

	wg.Wait()
	fmt.Println("finished....")
}
