package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	ctx1 := context.Background()
	ctx2, c2 := context.WithCancel(ctx1)
	ctx3, _ := context.WithCancel(ctx2)
	ctx4, _ := context.WithCancel(ctx3)

	defer c2()

	var wg sync.WaitGroup
	wg.Add(3)
	go first(ctx2, &wg)
	go second(ctx3, &wg)
	go third(ctx4, &wg)

	time.Sleep(time.Second * 5)

	wg.Wait()
}

func first(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			log.Println("first fn canceled!")
			ticker.Stop()
			wg.Done()
			return
		case <-ticker.C:
			log.Println("first fn is ticking...")
		}
	}
}

func second(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			log.Println("second fn canceled!")
			ticker.Stop()
			wg.Done()
			return
		case <-ticker.C:
			log.Println("second fn is ticking...")
		}
	}
}

func third(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			log.Println("third fn canceled!")
			ticker.Stop()
			wg.Done()
			return
		case <-ticker.C:
			log.Println("third fn is ticking...")
		}
	}
}
