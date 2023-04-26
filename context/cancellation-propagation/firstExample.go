package main

import (
	"context"
	"log"
)

func main() {
	ctx1 := context.Background()
	ctx2, c2 := context.WithCancel(ctx1) 
	ctx3, _ := context.WithCancel(ctx2)
	ctx4, _ := context.WithCancel(ctx3)

	go doWork(c2)

	for {
		select {
		case <-ctx2.Done():
			log.Println("context 2 canceled")
		case <-ctx3.Done():
			log.Println("context 3 canceled")
		case <- ctx4.Done():
			log.Println("context 4 canceled and the process will be done!")
			return
		}
	}
}

func doWork(c2 context.CancelFunc) {
	c2()
}