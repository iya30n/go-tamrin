package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)

		c <- "something"
	}()

	// if we don't define the default case, our select will be block until a case being run or close the channel.
	// if we define the default case, the select won't be blocked and running it.

	// close(c)

	select {
	case <-c:
		fmt.Println(<-c)
	default:
		fmt.Println("default")
	}
}
