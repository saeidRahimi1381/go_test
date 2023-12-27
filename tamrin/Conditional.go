// Conditional
package main

import "fmt"

func main() {

	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is vary big")
	} else {
		fmt.Println("x is big")
	}

	a := 11.0
	b := 20.0
	frac := a / b
	if frac > 0.5 {
		fmt.Println("a is more than half b")
	}

	SwitchStatmanets()

	/////////////////////////
	fmt.Print("\n\n\n")

	c := 48
	if c > 100 {
		fmt.Println("AAAA")
	} else if c > 50 && c < 100 {
		fmt.Println("BBBB")
	} else {
		fmt.Println("CCCC")
	}
}

func SwitchStatmanets() {
	x := 1

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("twe")
	case 3:
		fmt.Println("three")

	default:
		fmt.Println("many")
	}

}
