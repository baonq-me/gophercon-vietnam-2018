package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func AddMore(students []Student, name string) {
	students = append(students, Student{
		Name: name,
	})
}

func main() {
	//students := make([]Student, 0, 1000)
	students := make([]Student, 1000)

	fmt.Printf("Current capacity: %d\n", cap(students))
	fmt.Printf("Current length: %d\n", len(students))


	students[0] = Student {
		Name: "Bao",
	}

	students[1] = Student {
		Name: "Tuyen",
	}

	// error ?
	students = append(students, Student{
		Name: "Tuyen",
	})

	// [{Name:Bao} {Name:} {Name:Bao}]
	//fmt.Printf("%v \n", students)

	// [{Bao} {} {Bao}]
	//fmt.Printf("%+v \n", students)

	// [{Name:Tuyen}]
	//fmt.Printf("%+v \n", students[1:2])


}