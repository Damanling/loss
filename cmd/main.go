package main

import (
	"fmt"
	"os"
)

const a = "0.1"

func main() {
	if len(os.Args) < 2 {
		return
	}

	lol := os.Args[1]
	if lol == "hello" {
		fmt.Printf("hello, brother - это loss v%s\n", a)
		return
	}

	fmt.Printf("got: %s\n", lol)
}
