package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    line, _ := reader.ReadString('\n')
    line = strings.TrimSuffix(line, "\n")

    fmt.Println(line)

    // we can run our code with file as an input
    // go run main.go < ./inputs/1.txt
    // so, stdin contains any input
   // if we run go run main.go < ./inputs/1.txt > ./outputs/1.txt
    // it writes the output in the ./outputs/1.txt
    // because fmt.Println(line) is an stdout also files are I/O
}
