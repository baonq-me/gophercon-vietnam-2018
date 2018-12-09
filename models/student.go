package models

type Student struct {
	Name string	`json:"name"`
	Age int	`json:"my_age"`
}

type Students []Student

var TableStudents Students = Students{}

func (students Students) HasStudentName(name string) (bool, int) {
	// for index, student := range students {
	for index, student := range students {
		if student.Name == name {
			return true, index
		}
	}

	return false, -1
}