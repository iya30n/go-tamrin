package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Handler] request received")

		rCtx := r.Context()
		resChan := make(chan int)
		defer close(resChan)

		go doWork(rCtx, resChan)
		
		select {
		case <- rCtx.Done():
			return
		case result := <-resChan:
			log.Println("[Handler] Received 1000")
			log.Println("[Handler] Send response")
			fmt.Fprintf(w, "Response %d", result)
			return
		}
	})

	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		panic(err)
	}
}

func doWork(rCtx context.Context, resChan chan int) {
	log.Println("[doWork] launch the doWork")
	sum := 0
	for {
		log.Println("[doWork] one iteration")
		time.Sleep(time.Millisecond)

		select {
		case <- rCtx.Done():
			log.Println("[doWork] ctx Done is received inside doWork")
			return
		default:
			sum ++
			if sum > 1000 {
				log.Println("[doWork] sum has reached 1000")
				resChan <- sum
				return
			}
		}
	}
}