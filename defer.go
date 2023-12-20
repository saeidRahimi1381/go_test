// defer
package main

import "fmt"

func main() {
	tamrin_1()
	fmt.Println("\n\n")
	tamrin_2()
	fmt.Println("\n\n")
	tamrin_3()
	fmt.Println("\n\n")
}

func tamrin_1() {
	fmt.Println("tamrin 1:")

	defer func() {
		fmt.Println("hi golang")
	}()
	fmt.Println("hi saeid")
}

func tamrin_2() {
	fmt.Println("tamrin 2:")

	i := 1
	defer fmt.Println(i)
	i = 2
	defer fmt.Println(i)
	i = 3
	defer fmt.Println(i)
}

func tamrin_3() {
	fmt.Println("tamrin 3:")

	i := 1
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	fmt.Println("First")
}
