package main

import (
	"fmt"
	"sync"
)

func worker(workerId int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job
		fmt.Printf("Worker(%d) is processing job-%d\n", workerId, job)
	}
}

func main() {
	numOfWorkers := 3
	numOfJobs := 10

	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	var wg sync.WaitGroup

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
		close(results)
	}()

	for result := range results {
		fmt.Printf("Reading result: job-%d\n", result)
	}
}
