package main

import (
	"fmt"
	"time"
)

func doSomething(done chan<- bool) {
	fmt.Println("doing something...")
	time.Sleep(time.Second * 3)
	done <- true
}

func main() {
	done := make(chan bool, 1)

	go doSomething(done)

	// if our channel use after the function, our code will be synchronized.
	<-done

	fmt.Println("Done!")
}
