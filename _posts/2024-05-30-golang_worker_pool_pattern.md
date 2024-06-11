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


|<img src="https://i.imgur.com/vPEfkbp.png" alt="pods" width="550">|
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
		fmt.Printf("Worker %d processing job-%d\n", id, job)
		results <- job // Send the result of the job to the results channel
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

  // NOTE:
  // Non-blocking Main Goroutine:
  //    By placing wg.Wait() and close(results) in a `separate` goroutine,
  //    you ensure that the main goroutine can continue to collect results
  //    while waiting for all workers to finish.
  // Blocking Main Goroutine:
  //    The main goroutine blocks on wg.Wait() and waits for all workers
  //      to finish before proceeding to collect results.
  //    Possible Deadlock: If you put wg.Wait() and close(results) in the main goroutine,
  //      it will block on wg.Wait(), potentially causing a deadlock if the buffer in the results channel fills up.
  //      This is because workers will be trying to send results to a full results channel
  //      while the main goroutine is waiting on wg.Wait().
  //    Potential Deadlock Scenario
  //      If the main goroutine blocks on wg.Wait() and the results channel buffer fills up,
  //      the workers will be stuck trying to send results. The main goroutine, in turn,
  //      is stuck waiting for the workers to finish, creating a deadlock.

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
		fmt.Printf("Reading result: job-%d\n", result)
	}
}
```

The reason for creating a separate goroutine for `wg.Wait()` and `close(results)` is to ensure that the main goroutine does not block while waiting for all the worker goroutines to finish. Let's break down the process and explain the differences:

### Understanding the Flow

1. **Creating Channels and Starting Workers:**
   - Channels `jobs` and `results` are created with a buffer size of `numJobs`.
   - Worker goroutines are started and each worker processes jobs from the `jobs` channel and sends the results to the `results` channel.

2. **Enqueuing Jobs:**
   - Jobs are enqueued into the `jobs` channel.
   - Once all jobs are enqueued, the `jobs` channel is closed to signal to the workers that there are no more jobs to process.

3. **Waiting for Workers to Finish:**
   - A separate goroutine is started to wait for all worker goroutines to finish (`wg.Wait()`).
   - After all workers are done, this goroutine closes the `results` channel.

4. **Collecting Results:**
   - The main goroutine collects and processes results from the `results` channel until it is closed.

### Why a Separate Goroutine for `wg.Wait()` and `close(results)`?

By placing `wg.Wait()` and `close(results)` in a separate goroutine, you ensure that the main goroutine can continue to collect results while waiting for all workers to finish. Here's what happens in both cases:

#### Case 1: Separate Goroutine for `wg.Wait()` and `close(results)`

```go
go func() {
 wg.Wait()
 close(results)
}()
```

- **Non-blocking Main Goroutine:** The main goroutine does not block and can continue to collect results from the `results` channel.
- **Graceful Completion:** Once all workers are done, the `results` channel is closed, allowing the main goroutine to exit the `for` loop and terminate gracefully.

#### Case 2: `wg.Wait()` and `close(results)` in Main Goroutine

```go
wg.Wait()
close(results)
```

- **Blocking Main Goroutine:** The main goroutine blocks on `wg.Wait()` and waits for all workers to finish before proceeding to collect results.
- **Possible Deadlock:** If you put `wg.Wait()` and `close(results)` in the main goroutine, it will block on `wg.Wait()`, potentially causing a deadlock if the buffer in the `results` channel fills up. This is because workers will be trying to send results to a full `results` channel while the main goroutine is waiting on `wg.Wait()`.

### Potential Deadlock Scenario

If the main goroutine blocks on `wg.Wait()` and the `results` channel buffer fills up, the workers will be stuck trying to send results. The main goroutine, in turn, is stuck waiting for the workers to finish, creating a deadlock.

### Conclusion

Using a separate goroutine for `wg.Wait()` and `close(results)` allows the main goroutine to continue processing results without blocking, preventing the risk of a deadlock and ensuring smooth program execution. This design pattern is commonly used in concurrent programming to decouple the coordination of goroutines from the main execution flow.


[↑ Back to top](#)
<br><br>


