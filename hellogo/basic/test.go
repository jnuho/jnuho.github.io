package main

import "sync"

func main() {
	numOfJobs := 10
	numOfWorkers := 3

	var wg sync.WaitGroup
	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= numOfJobs; j++ {
				println("Job", j, "done")
			}
		}()
	}

}
