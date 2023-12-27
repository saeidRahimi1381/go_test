// Name return values
package main

import "fmt"

func main() {
	fmt.Println(split(18))
}

func split(sum int) (x, y int) {
	x = sum * 4 / 7
	y = sum - 3
	return
}
