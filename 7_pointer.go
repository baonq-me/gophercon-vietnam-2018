package main

import "fmt"

// Student is defined in 4_struct.go

func main() {

	var a *Student = &Student{
		Name: "Bao",
	}
	fmt.Printf("Student name: %d\n", a)

	myAge := 27
	fmt.Printf("My age: %d\n", myAge)
	fmt.Printf("My age address: %d\n", &myAge)

}