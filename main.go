package main

import (
	"./models"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

// https://golang.org/cmd/cgo/

func main() {
	e := echo.New()
	// GET /students
	// GET /students/{name}
	// POST /students
	// PUT /students/{name}
	// DELETE /students/{name}
	fmt.Printf("BUILD a API")

	e.GET("/students", GetStudents)
	e.POST("/students", CreateStudents)
	e.DELETE("/students/:name", DeleteStudents)
	e.Logger.Fatal(e.Start(":1234"))



}

func GetStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, models.TableStudents)
}

func CreateStudents(c echo.Context) error {
	var creatingStudent models.Student
	err := c.Bind(&creatingStudent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed !\n")
	}

	models.TableStudents = append(models.TableStudents, creatingStudent)

	return c.JSON(http.StatusOK, "OK")
}

func DeleteStudents(c echo.Context) error {
	studentName := c.Param("name")
	exist, studentIndex := models.TableStudents.HasStudentName(studentName)
	if !exist {
		return c.JSON(http.StatusNotFound, "Not found")
	}

	// Delete a student at studentIndex in models.TableStudents
	models.TableStudents = append(
		models.TableStudents[:studentIndex],		// is a list
		models.TableStudents[studentIndex+1:]...,   // need ... to be a list
	)

	return c.JSON(http.StatusOK, "OK")
}
