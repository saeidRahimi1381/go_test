// Function
package main

import "fmt"

func main() {
	fmt.Println(division(10, 2))
	fmt.Println(sum(4, 7))
}

func division(x int, y int) int {
	return x / y
}

func sum(w, z int) int {
	return w + z
}
