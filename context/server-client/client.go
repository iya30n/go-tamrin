package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	rootCtx := context.Background()

	req, err := http.NewRequest("GET", "http://127.0.0.1:8989", nil)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(rootCtx, time.Second * 10)
	defer cancel()

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	log.Println("Response received", res)
}