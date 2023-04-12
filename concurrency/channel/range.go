package main

import "fmt"

func main() {
	// the range is working only for buffered channels.
	
	c := make(chan string, 3)

	c <- "first"
	c <- "second"
	c <- "third"

	close(c)

	for value := range c {
		fmt.Println(value)
	}
}