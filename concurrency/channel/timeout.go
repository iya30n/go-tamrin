package main

import (
	"fmt"
	"time"
)

func somethingToDo(done chan<- bool) {
	fmt.Println("doing something...")
	time.Sleep(time.Second * 3)
	done <- true
}

func main() {
	done := make(chan bool)

	go somethingToDo(done)

	// if we close the channel before select, it'll get default value and and the case will be run.
	// we should close the channel after all (using defer or after select).
	defer close(done)

	select {
	case <-done:
		fmt.Println(<-done)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout Error")
	}
}
