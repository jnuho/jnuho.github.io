package main

import (
	"fmt"
	"sync"
)

func worker(workerId int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d is working on job %d\n", workerId, job)
		results <- job * 3.14
	}
}

func main() {
	numOfJobs := 10
	numOfWorkers := 3

	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	var wg sync.WaitGroup

	// jobs
	for i := 1; i <= numOfWorkers; i++ {

		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			worker(workerId, jobs, results)
		}(i)
	}

	for i := 1; i <= numOfJobs; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
	}()
	// results
	for result := range results {
		fmt.Println(result)
	}
}
