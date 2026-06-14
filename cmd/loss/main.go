package main

import (
	"fmt"
	"os"
)

const version = "0.1"

func main() {
	if len(os.Args) < 2 {
		return
	}

	 commandName:= os.Args[1]
	if commandName == "hello" {
		fmt.Printf("hello, brother - это loss v%s\n", version)
		return
	}

	fmt.Printf("got: %s\n", commandName)
}
