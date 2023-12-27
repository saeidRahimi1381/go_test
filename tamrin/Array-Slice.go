// Array & Slice

package main

import (
	"fmt"
)

func main() {
	//Array
	var a [6]int
	a = [6]int{1, 2, 3, 4, 5, 6}
	for index, value := range a {
		fmt.Println(index, value)
	}

	////////////////////////
	fmt.Print("\n\n\n")

	//Slice

	var b []int
	b = append(b, 5)
	fmt.Println(b)
	fmt.Println(cap(b))
	fmt.Println(len(b))

	c := make([]int, 5)
	fmt.Println(cap(c))
	fmt.Println(len(c))
	c = []int{1, 2, 3, 4, 5}
	fmt.Println(len(c), cap(c))
	c = append(c, 6)
	fmt.Println(len(c), cap(c))
	fmt.Println(c[3])
	fmt.Println(c[3:5])
	cc := c[1:4]
	fmt.Println(cc)
}
