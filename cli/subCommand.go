package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable the foo")
	fooName := fooCmd.String("name", "mamad", "change the name of the foo")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "change the level of the bar")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand:", "foo")
		fmt.Println("enable:", *fooEnable)
		fmt.Println("name:", *fooName)
		fmt.Println("tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand:", "bar")
		fmt.Println("level:", *barLevel)
		fmt.Println("tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
