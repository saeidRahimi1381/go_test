package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	fmt.Print("Type a number: ")
	array := []string{}
	for i := 0; i < 5; i++ {
		var temp string
		fmt.Scan(&temp)
		array = append(array, temp)
	}
	char := "A"
	fmt.Println(array)

	var items = findByName(array, char)
	appendFamily(items)

}
func appendFamily(arr []string) {
	for i, element := range arr {

		arr[i] = element + "-" + randomeFamily()
		fmt.Print("\n", arr[i])
	}
	fmt.Print("\n", arr)
}

func findByName(arr []string, character string) []string {

	cloneArray := []string{}

	for _, item := range arr {

		if strings.Contains(item, character) {
			
			cloneArray = append(cloneArray, item)
		}

	}

	return cloneArray
}

func randomeFamily() string {
	familyArray := []string{"hsseinali", "hasan ali", "mehmandoost", "gholamy", "vefgh"}
	var randomeitem = rand.Intn(len(familyArray) - 0)
	c := familyArray[randomeitem]
	return c
}
