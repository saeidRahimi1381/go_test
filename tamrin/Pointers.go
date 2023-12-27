// Pointers

package main

import "fmt"

func main() {
	var p *int
	i := 43

	fmt.Println("p", p)
	fmt.Println("i", i)

	p = &i
	fmt.Println("p", p)
	fmt.Println("value that p references to", *p)

	////////////////////////
	fmt.Print("\n\n\n")

	name := "saeid"

	fmt.Println(name)
	changeName(&name, "Ali")
	fmt.Println(name)
}

func changeName(nameAddress *string, newName string) {
	*nameAddress = newName
}
