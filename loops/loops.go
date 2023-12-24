package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		fmt.Println(i)
		if i == 50 {
			break
		}
	}

	////////////////////////
	fmt.Print("\n\n\n")

	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}
}
