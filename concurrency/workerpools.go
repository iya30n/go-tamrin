package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("The worker %d working on job %d \n", id, j)
		time.Sleep(time.Second * 3)
		fmt.Printf("The worker %d finished the job %d \n", id, j)
		results <- j
	}
}

func main() {
	const jobsCount = 5
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= jobsCount; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 1; i <= jobsCount; i++ {
		<-results
	}
}
