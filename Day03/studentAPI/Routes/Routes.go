package Routes

import (
	"studentAPI/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/student-api")
	{
		grp1.GET("student", Controllers.GetStudents)				// Get the data for all the students
		grp1.POST("student", Controllers.CreateStudent)			// Input data for a new student
		grp1.GET("student/:id", Controllers.GetStudentByID)		// Get the info of a specific student
		grp1.GET("marks/:id", Controllers.GetMarksByID)		// Get the marks of a specific student

		grp1.PUT("student/:id", Controllers.UpdateInfo)			// Update student info
		grp1.PUT("marks/:id", Controllers.UpdateMarks)			// Update student marks
		grp1.DELETE("student/:id", Controllers.DeleteStudent)		// Delete student info
	}
	return r
}
