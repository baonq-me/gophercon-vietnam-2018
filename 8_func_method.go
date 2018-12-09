package main

import "fmt"

type Student1 struct {
	Name string
}

// is a method of struct Student1
func (s Student1) PrintGreeting() {
	fmt.Printf("Student name: %s\n", s.Name)
}

func main() {
	var s Student1 = Student1{
		Name: "Bao",
	}

	Student1.PrintGreeting(s)
}
