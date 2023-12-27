// Struct
package main

import "fmt"

func main() {
	type Student struct {
		LastName     string
		FirstNmae    string
		Age          uint
		ID           uint
		NationNumber string
	}

	var u Student
	fmt.Println(u.Age)
	fmt.Println(u.LastName)

	////////////////////////
	fmt.Print("\n\n\n")

	var a = Student{
		LastName:     "Rahimi",
		FirstNmae:    "Saeid",
		Age:          20,
		ID:           123,
		NationNumber: "1234",
	}
	fmt.Println(a.Age)
	fmt.Println(a.LastName)

	////////////////////////
	fmt.Print("\n\n\n")

	var i Student
	i.Age = 12
	i.LastName = "karimi"
	fmt.Println(i.Age)
	fmt.Println(i.LastName)

	////////////////////////
	fmt.Print("\n\n\n")

	var j = Student{"ali", "rad", 18, 1234, "123456"}
	fmt.Println(j)

	////////////////////////
	fmt.Print("\n\n\n")

	fmt.Println("My Name is", j.FirstNmae, "and I'm", j.Age, "years old.", "My National code is", j.NationNumber)

	////////////////////////
	fmt.Print("\n\n\n")

	fmt.Printf("My Name is %s and I'm %d years old. My National code is %s \n", j.FirstNmae, j.Age, j.NationNumber)

}
