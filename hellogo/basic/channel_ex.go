


























package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    // Process jobs from the jobs channel until it is closed
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2 // Send the result of the job to the results channel
    }
}

func main() {
    numWorkers := 3
    numJobs := 10

    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Start workers
    var wg sync.WaitGroup
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            worker(workerID, jobs, results)
        }(i)
    }

    // Add jobs to the jobs channel
    for i := 1; i <= numJobs; i++ {
        jobs <- i
    }
    close(jobs)

    // Wait for all workers to finish and close the results channel
    go func() {
        wg.Wait()
        close(results)
    }()

    // Process results from the results channel
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
}

