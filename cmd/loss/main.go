package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.1"

func main() {
	if len(os.Args) < 2 {
		usage()





		os.Exit(1)
	}
	hello := flag.NewFlagSet(os.Args[1], flag.ExitOnError)

	switch os.Args[1] {
	case "hello":
		var name = hello.String("name", "brother", "имя пользователя")
		if err := hello.Parse(os.Args[2:]); err != nil {
			usage()
			os.Exit(1)
		}
		fmt.Println("hello, " + *name + " — это loss v" + version)

		fmt.Println('test')

	case "echo":
		if len(os.Args) < 3 {
			usage()
			os.Exit(1)
		}
		fmt.Println("echo:", os.Args[2:])

	default:
		usage()
		os.Exit(1)
	}
}
func usage() {
	fmt.Fprintln(os.Stderr, "usage: ")
	fmt.Fprintln(os.Stderr, "loss hello")
	fmt.Fprintln(os.Stderr, "loss echo args")
}
