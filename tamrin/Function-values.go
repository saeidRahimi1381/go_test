package main

import "fmt"

func main() {
	f := func(b, p int) int {
		return b * p
	}
	fmt.Println(f(2, 4))

	////////////////////////
	fmt.Print("\n\n\n")

	compute(6, print)

	////////////////////////
	fmt.Print("\n\n\n")

	a := func(t int) {
		fmt.Println("t", t)
	}
	a(35)
}

func compute(x int, f func(r int)) int {
	result := x * 2
	f(result)
	return result
}

func print(d int) {
	fmt.Println("d", d)
}
