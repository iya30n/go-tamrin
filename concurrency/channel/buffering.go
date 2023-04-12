package main

import "fmt"

func main() {
	c := make(chan string, 2)

	c <- "first"
	c <- "second"
	// c <- "third"

	close(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}