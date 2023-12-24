package main

import "fmt"

var (
	name = "saeid"
)

func main() {
	name2 := &name
	*name2 = "mohammad"
	fmt.Println(name)
	fmt.Println(name2)

	////////////////////
	fmt.Print("\n\n\n")

	fmt.Println(name)
	changeName(&name, "Ali")
	fmt.Println(name)
}

func changeName(nameAddress *string, newName string) {
	*nameAddress = newName
}
