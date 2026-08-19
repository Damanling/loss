package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.1"

func list() {
	fmt.Println("1500")
	fmt.Println("food")
	fmt.Println("кофе")
}

type Categories string

const (
	CategoryFood      = "food"
	CategoryTransport = "transport"
	CategoryFun       = "fun"
	CategoryBills     = "bills"
	CategoryOther     = "other"
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

		switch *cat {
		case CategoryFood, CategoryTransport, CategoryFun, CategoryBills, CategoryOther:
			fmt.Println("Категория выбрана верно")
		case "empty":
			fmt.Fprintln(os.Stderr, "cat flag is required")
			os.Exit(1)
		default:
			fmt.Fprintln(os.Stderr, "unknown category, allowed: food, transport, fun, bills, other")
			os.Exit(1)
		}

		fmt.Println(*amount)
		fmt.Println(*cat)
		fmt.Println(*note)

	case "list":
		list()

	case "cats":
		fmt.Println(CategoryFood)
		fmt.Println(CategoryTransport)
		fmt.Println(CategoryFun)
		fmt.Println(CategoryBills)
		fmt.Println(CategoryOther)

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
	fmt.Fprintln(os.Stderr, "loss cats")
}
