// Method

package main

import "fmt"

type student struct {
	Name   string
	ID     int
	Scores []int
}

func main() {
	st := student{
		Name:   "saeid",
		ID:     123456,
		Scores: []int{5, 10, 18, 7},
	}

	// printStudent(st)
	// fmt.Println(studentAvg(st))
	// fmt.Println(isStudentEligible(st))

	st.printStudent()
	fmt.Println(st.StudentAvg())
	//fmt.Println(st.isStudentEligible())
}

func (s student) printStudent() {
	fmt.Printf("the student name is %s and the id is %d\n", s.Name, s.ID)
}

func (s student) StudentAvg() float64 {
	sum := 0
	for _, s := range s.Scores {
		sum += s
	}

	return float64(sum) / float64(len(s.Scores))
}

// func (s student) isStudentEligible() bool {
// 	if StudentAvg(s) > 12 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
