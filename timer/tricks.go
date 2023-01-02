package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 300)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at ", t)
			}
		}
	}()

	time.Sleep(time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped!")
}
