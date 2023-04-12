package main

import (
	"fmt"
	"time"
)

func printer(c <- chan string) {
	for {
		time.Sleep(time.Second)

		fmt.Println(<-c)
	}
}

func pinger(c chan <- string) {
	for {
		c <- "ping"
	}
}

func ponger(c chan <- string) {
	for {
		c <- "pong"
	}
}

func main() {
	c := make(chan string)

	go printer(c)
	go pinger(c)
	go ponger(c)

	for {
		var input string
		fmt.Scanln(&input)

		if len(input) == 0 {
			break
		}

		c <- input
	}
}