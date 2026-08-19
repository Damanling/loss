package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.1"

type Categories string
const (
	CategoryFood = "food"
	CategoryTransport = "transport"
	CategoryFun = "fun"
	CategoryBills = "bills"
	CategoryOther = "other"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	hello := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	add := flag.NewFlagSet(os.Args[1], flag.ExitOnError)

	switch os.Args[1] {
	case "hello":
		var name = hello.String("name", "brother", "имя пользователя")

		if err := hello.Parse(os.Args[2:]); err != nil {
			usage()
			os.Exit(1)
		}

		fmt.Printf("hello, %s — это loss v%s \n", *name, version)

	case "echo":
		if len(os.Args) < 3 {
			usage()
			os.Exit(1)
		}

		fmt.Printf("echo: %s\n", os.Args[2:])

	case "add":
		var amount = add.Int("amount", 0, "потраченная сумма")

		if err := add.Parse(os.Args[2:]); err != nil {
			usage()
			os.Exit(1)
		}

		fmt.Println(*amount)

		var cat = add.String("cat", "empty", "категории")
		if *cat {
			
		}
		if *cat != CategoryFood && CategoryTransport && CategoryFun && CategoryBills && CategoryOther {
			fmt.Fprintln(os.Stderr, "unknown category, allowed: food, transport, fun, bills, other")
			os.Exit(1)
		}
		fmt.Println(*cat)

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
