package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
)

const version = "0.1"

func list() {

}

var categories = []string{"food", "transport", "fun", "bills", "other"}

func validCategory(c string) bool { return slices.Contains(categories, c) }

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
		add := flag.NewFlagSet("add", flag.ContinueOnError)

		var amount = add.Int64("amount", 0, "потраченная сумма")
		var cat = add.String("cat", "empty", "категории")
		var note = add.String("note", "empty", "что купил")

		if err := add.Parse(os.Args[2:]); err != nil {
			usage()
			os.Exit(1)
		}

		if *amount == 0 {
			fmt.Fprintln(os.Stderr, "amount flag is required")
			os.Exit(1)
		}

		if *amount < 0 {
			fmt.Fprintln(os.Stderr, "amount must be positive")
			os.Exit(1)
		}

		if !validCategory(*cat) {
			fmt.Fprintln(os.Stderr, "unknown category, allowed: food, transport, fun, bills, other")
			os.Exit(1)
		}

		fmt.Printf("%d,%s,%s\n", *amount, *cat, *note)

	case "list":
		listFlags := flag.NewFlagSet("list", flag.ContinueOnError)
		if err := listFlags.Parse(os.Args[2:]); err != nil {
			os.Exit(1)
		}

		list()

	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: ")
	fmt.Fprintln(os.Stderr, "loss hello")
	fmt.Fprintln(os.Stderr, "loss echo args")
	fmt.Fprintln(os.Stderr, "loss add --amount 100 --cat food --note 'что купил'")
	fmt.Fprintln(os.Stderr, "loss list")
}
