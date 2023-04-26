package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "some_key", "some value")
	ctx = context.WithValue(ctx, "another_key", "another value")

	// myVal will be nil because we don't have any value with the key "nothing"
	// myVal := ctx.Value("nothing")

	// myVal, ok := ctx.Value("nothing_else").(string)

	fmt.Println(ctx.Value("another_key"))
}