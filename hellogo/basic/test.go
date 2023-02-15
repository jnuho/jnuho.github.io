package main

import (
	"fmt"
	"time"
)

// jobs for performing work
// results for sending result back to main goroutine
func worker(id int, jobs chan int, results chan int) {
	// wait until channels are populated
	fmt.Printf("Start worker %d \n", id)
	for j := range jobs {
		time.Sleep(time.Second)
		results <- j*2
	}
	fmt.Printf("End worker %d \n", id)
}

func main() {
	const numOfJobs = 5
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	// Start up to 3 workers
	// with empty channel jobs and results
	for i:=0; i<3; i++ {
		go worker(i, jobs, results)
	}

	// 5개 
	for i:=0; i< numOfJobs; i++ {
		fmt.Printf("Main : put %d value into jobs\n", i)
		jobs <- i
	}

	close(jobs)

	// The workers perform the jobs by receiving them from the jobs channel
	// and sending the results to the results channel.
	// The use of channels ensures that the jobs are performed in order
	// and that the main goroutine waits for all the workers to finish before exiting.
	for i:=0; i< numOfJobs; i++ {
		a := <- results
		fmt.Printf("Result:  value from channel : %d\n", a)
	}
}
