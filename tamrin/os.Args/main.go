package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("/n/n")
	args := os.Args
	fmt.Println("lan(Args)", len(args))

	if len(args) < 3 {
		//fmt.Println("error-there is not input")
		//os.Exit(1)
		//fmt.Println("hi!")
		//return
		log.Fatalf("there is no input - len of args: %d", len(args))
	}
	name := args[1]
	fmt.Println("hello", name)
	name2 := args[2]
	fmt.Println("hello2", name2)
}
