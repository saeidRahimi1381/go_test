// Type conversion
package main

import "fmt"

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

func main() {
	const pattern = "%T(%v)\n"
	fmt.Printf(pattern, i, i)
	fmt.Printf(pattern, f, f)
	fmt.Printf(pattern, u, u)
}
