package main

import (
	"sync"
	"time"
)

func worker() {
	time.Sleep(time.Second * 3)

	// wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		// go worker(&wg)
		go func() {
			wg.Done()

			worker()
		}()
	}

	wg.Wait()
}
