package main

import "fmt"

type cars struct {
	Bebz, BMW string
}

func main() {
	m := make(map[string]cars)
	m["Benz"] = cars{
		"2700.5 $", "4600.7 $",
	}
	fmt.Println(m["Benz"])

	////////////////////////
	fmt.Print("\n\n\n")

	Mutating_Maps()

}

func Mutating_Maps() {
	c := make(map[string]int)

	c["Answer"] = 42
	fmt.Println("The value:", c["Answer"])

	c["Answer"] = 48
	fmt.Println("The value:", c["Answer"])

	delete(c, "Answer")
	fmt.Println("The value:", c["Answer"])

	v, ok := c["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
