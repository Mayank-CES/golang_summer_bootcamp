package Controllers

import (
	"fmt"
	"net/http"
	"studentAPI/Models"

	"github.com/gin-gonic/gin"
)

//GetStudents ... Get all students
func GetStudents(c *gin.Context) {
	var students []Models.Student
	err := Models.GetAllStudents(&students)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

//CreateStudent ... Create User
func CreateStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.CreateStudent(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//GetStudentByID ... Get the user by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetStudentByID(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//UpdateInfo ... Update the user information
func UpdateInfo(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentByID(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models.UpdateInfo(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}



//DeleteStudent ... Delete the user
func DeleteStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

//GetMarksByID ... Get the user by id
func GetMarksByID(c *gin.Context) {
	id := c.Params.ByName("id")

	var student Models.Student						// Check for availability of student
	err := Models.GetStudentByID(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}

	var marks Models.Marks							// Check for availability of marks
	err = Models.GetMarksByID(&marks, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}


//UpdateMarks ... Update the user information
func UpdateMarks(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")

	err := Models.GetStudentByID(&student, id)		// Check for availability of student
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}

	var marks Models.Marks							// Check for availability of marks
	//err = Models.GetMarksByID(&marks, id)
	//if err != nil {
	//	c.JSON(http.StatusNotFound, marks)
	//}
	c.BindJSON(&marks)

	err = Models.UpdateMarks(&marks, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

