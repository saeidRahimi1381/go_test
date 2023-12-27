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

	/////////////////////////
	fmt.Print("\n\n\n")

	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	/////////////////////////
	fmt.Print("\n\n\n")

	j := 0
	for j < 100 {
		j = j + 1
	}
	fmt.Println(j)

}
