package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d working on job %d \n", id, j)
		time.Sleep(time.Second * 3)
		fmt.Printf("worker %d finished the job %d \n", id, j)
		results <- j
	}
}

func main() {
	const jobsCount = 5

	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			
			worker(i, jobs, results)
		}()
	}

	for i := 1; i <= jobsCount; i++ {
		jobs <- i
	}
	close(jobs)

	/* for i := 0; i < jobsCount; i++ {
		<-results
	} */
	wg.Wait()
}
