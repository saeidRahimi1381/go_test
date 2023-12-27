// Variadic Function

package main

import "fmt"

func main() {
	pre()
	pre("saeid")
	pre("hello", "world")
}

func pre(values ...string) {
	for _, value := range values {
		fmt.Println("v", value)
	}
}
