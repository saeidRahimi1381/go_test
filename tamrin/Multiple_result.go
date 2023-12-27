// Multiple result
package main

import "fmt"

func main() {
	a, b := swap(4, 6)
	fmt.Println(a, b)
}

func swap(x, y int) (int, int) {
	return y, x
}
