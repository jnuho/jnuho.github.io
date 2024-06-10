package main

import "sync"

func

func main() {
	numOfJobs := 10
	numOfWorkers := 3

	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	var wg sync.WaitGroup

	// jobs
	for i := 1; i <= numOfWorkers; i++ {

		wg.Add(1)
		go func() {

		}(i)
	}

	// results
}
