package main

import (
	"flag"
	"fmt"
)

func main() {
	st := flag.String("option", "something", "test description")
	intnum := flag.Int("count", 10, "count of something")

	flag.Parse()

	fmt.Println(*st)
	fmt.Println(*intnum)
}