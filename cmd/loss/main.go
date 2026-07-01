package main

import (
	"fmt"
	"os"
)

const version = "0.1"

 func main(){
	if( len(os.Args) < 2){
	usage()
	os.Exit(1)
	}

	if arg := os.Args[1]; arg == "hello" {
		fmt.Println("hello, brother — это loss v"+ version)
	} else if arg := os.Args[1]; arg == "echo" {
			fmt.Println("echo: ", os.Args[2:])
			if len(os.Args) < 3 {
				usage()
				os.Exit(1)
		}
	} else {
		usage()
		os.Exit(1)
	}
 }

 func usage() {
	fmt.Fprintln(os.Stderr,"usage: ")
	fmt.Fprintln(os.Stderr,"loss hello")
	fmt.Fprintln(os.Stderr,"loss echo args")
 }
