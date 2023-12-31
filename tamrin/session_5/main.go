package main

import (
	"apptodo/function"
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello to TODO app")

	command := flag.String("command", "no cammand", "command to run ")
	flag.Parse()

	for {
		function.Runcommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}

	fmt.Printf("UserStorage: %v \n", function.UserStorage)
}
