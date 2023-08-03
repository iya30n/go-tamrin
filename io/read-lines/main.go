package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    file, err := os.Open("./items.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
