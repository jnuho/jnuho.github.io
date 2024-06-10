---
layout: post
title: Worker pool pattern
categories: [project]
tags: [Concurrency, golang]
fullview: false
comments: true
shortinfo: Golang Concurrency pattern
preview_worker_pool: true
---


|<img src="https://d17pwbfgewyq5y.cloudfront.net/worker_pool_pattern.drawio.png" alt="pods" width="550">|
|:--:| 
| *Worker pool pattern* |

<br>

- The worker pool pattern involves creating a group of worker goroutines to process tasks concurrently,
- limiting the number of simultaneous operations. This pattern is valuable when you have a large number of tasks to execute.
- Some examples of using the Worker Pool Pattern in Real-world Applications
  - Handling incoming HTTP requests in a web server.
  - Processing images concurrently.

```go
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
		go func(workerId int) {
			defer wg.Done()
			worker(workerId, jobs, results)
		}(i)
	}

	// Add jobs to the jobs channel
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
  // close channel since all values sent to jobs channel
	close(jobs)

	// Wait for all workers to finish and close the results channel
	// when the counter reaches zero, wg.Wait() unblocks
	// the counter reaches zero when `worker` finishes putting values from `jobs` to `results`
	// since all is sent to `results` channel, let's close results channle here
	go func() {
		wg.Wait()
    // close channel since all values sent to results channel
    // because the fact that wg.Wait() is called when all wg.Done() is executed implies
    // that `worker(workerID, jobs, results)` function's `results <- job * 2` operation is complete
    // since all values sent to the `results` channel we can close `results` channel here!
		close(results)
	}()

	// Process results from the results channel
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
```

[↑ Back to top](#)
<br><br>


