package main

import (
	"fmt"
	"sync"
)

var x uint64

/* if we just want to sum an integer, we can use the atomic.AddInt64 instead*/
/* if we need goroutines communicate together, it's better to use channel synchronization, else mutex is a better choice */
/* My advice would be to choose the tool for the problem and do not try to fit the problem for the tool :) */

func incX(wg *sync.WaitGroup, mt *sync.Mutex) {
	mt.Lock()
	x ++
	mt.Unlock()

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mt sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incX(&wg, &mt)
	}

	wg.Wait()

	fmt.Println(x)
}