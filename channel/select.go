package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan string), make(chan string)

	go func() {
		time.Sleep(time.Second*2)

		c1 <- "first"
	}()

	go func() {
		time.Sleep(time.Second*3)

		c1 <- "second"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)
	}
}
