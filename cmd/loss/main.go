package main

import (
	"fmt"
	"os"
)

const version = "0.1"

 func sff(){
	if( len(os.Args) < 2){
	return
	}
	var arg string = os.Args[1]
	if(arg == "hello"){
		fmt.Printf("ddddd%s\n",version)
		return
	}
fmt.Println("got: "+ version)
 }