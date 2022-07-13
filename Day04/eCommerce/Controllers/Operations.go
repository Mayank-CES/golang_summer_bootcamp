package Controllers

import (
	"fmt"
	"net/http"

	"github.com/Mayank-CES/golang_summer_bootcamp/Day04/eCommerce/Models"
	"github.com/Mayank-CES/golang_summer_bootcamp/Day04/eCommerce/Services"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Controller struct {
	service Services.CustomerServices
}

func NewController(services Services.CustomerServices) *Controller {
	return &Controller{service: services}
}

// CreateCustomerAccount ... Create Customer Account
func (cnt *Controller) CreateCustomerAccount(c *gin.Context) {
	var customer Models.Customer
	c.BindJSON(&customer)
	err := cnt.service.CreatCustomerAccount(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

/*
//BuyProduct ... Buy the product
func (cnt *Controller) BuyProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := cnt.service.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
		fmt.Println("An Error occurred while fetching the product")
	}
	c.BindJSON(&product)
	err = cnt.service.BuyProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
*/

// CheckOrderByID ... Get the product by id
func (cnt *Controller) CheckOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Transaction
	err := cnt.service.CheckOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

/*------- Products ------*/

// GetAllProducts  ... Get all Products
func (cnt *Controller) GetAllProducts(c *gin.Context) {
	var products []Models.Product
	err := cnt.service.GetAllProducts(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

// GetProductByID ... Get the product by id
func (cnt *Controller) GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := cnt.service.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

/*---------- Retailer ----------*/

// CreateRetailerAccount  ... Create Customer Account
func (cnt *Controller) CreateRetailerAccount(c *gin.Context) {
	var retailer Models.Retailer
	c.BindJSON(&retailer)
	err := cnt.service.CreatRetailerAccount(&retailer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, retailer)
	}
}

// AddProduct ... Create Customer Account
func (cnt *Controller) AddProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := cnt.service.AddProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// PatchProduct ... Buy the product
func (cnt *Controller) PatchProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := cnt.service.GetProductByID(&product, id)
	if err != nil {
		// c.JSON(http.StatusNotFound, product)
		fmt.Println("An Error occurred while fetching the product")
	}
	var updatedProduct Models.PatchProd
	c.BindJSON(&updatedProduct)
	err = cnt.service.PatchProduct(&updatedProduct, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// GetRHistoryByID ... Get the History by id
func (cnt *Controller) GetRHistoryByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Transaction
	err := cnt.service.GetRHistoryByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

// //UpdateInfo ... Update the user information
// func (cnt
// *Controller) UpdateInfo(c *gin.Context) {
//	var student Models.Student
//	id := c.Params.ByName("id")
//	err := Models.GetStudentByID(&student, id)
//
//	if err != nil {
//		c.JSON(http.StatusNotFound, student)
//	}
//	c.BindJSON(&student)
//	err = Models.UpdateInfo(&student, id)
//	if err != nil {
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.JSON(http.StatusOK, student)
//	}
// }

/*
//DeleteStudent ... Delete the user
func (cnt *Controller) DeleteStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&student, id)

	// ********   ALSO DELETE THE MARKS TABLE ASSOCIATED WITH THE STUDENT   ********

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var marks []Models.Marks
		err = Models.DeleteAllMarks(&marks,id)
		if err==nil{
			c.JSON(http.StatusOK, gin.H{"All marks for Student " + id: " are deleted"})
		}else{
			c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}

}


//GetAllMarks ... Get all marks
func (cnt *Controller) GetAllMarks(c *gin.Context) {
	var marks []Models.Marks							// Check for availability of marks
	err := Models.GetAllMarks(&marks)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}


//GetMarksByID ... Get the user by id
func (cnt *Controller) GetMarksByID(c *gin.Context) {

	studentId := c.Params.ByName("student_id")
	var student Models.Student						// Check for availability of student
	err := Models.GetStudentByID(&student, studentId)
	if err == nil {
		var marks Models.Marks							// Check for availability of marks
		id := c.Params.ByName("id")
		err = Models.GetMarksByID(&marks, id)

		if err==nil {
			var empty = Models.Marks{}
			if marks != empty {
				c.JSON(http.StatusOK, marks)
			} else {
				c.JSON(http.StatusOK, "No Marks Data found for the student")
			}
		}else{
			c.AbortWithStatus(http.StatusNotFound)
		}
	}else{
		c.JSON(http.StatusOK, "Student ID Not Found")
	}

}




//UpdateMarks ... Update the user information
func (cnt *Controller) UpdateMarks(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")

	studentId := c.Params.ByName("student_id")
	err := Models.GetStudentByID(&student, studentId)		// Check for availability of student
	if err == nil {
		var marks Models.Marks							// Check for availability of marks
		c.BindJSON(&marks)

		err = Models.UpdateMarks(&marks, id)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, marks)
		}
	}else{
		c.AbortWithStatus(http.StatusNotFound)
	}


}

//DeleteMarks ... Delete the marks by id		*********** 1 More case-----No result found for marks for the id
func (cnt *Controller) DeleteMarks(c *gin.Context) {
	studentId := c.Params.ByName("student_id")
	id := c.Params.ByName("id")
	var student Models.Student						// Check for availability of student
	err := Models.GetStudentByID(&student, studentId)
	if err==nil {
		var marks Models.Marks // Check for availability of marks
		err = Models.GetMarksByID(&marks, studentId)
		if err==nil{
			err = Models.DeleteMarks(&marks,id)
			if err==nil{
				c.JSON(http.StatusOK, gin.H{"Marks for Student " + studentId: "and Marks ID " +id+ " are deleted"})
			}else{
				c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			}
		}else{
			c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		}

	}else{
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}

}

func (cnt *Controller) DeleteAllMarks(c *gin.Context) {
	studentId := c.Params.ByName("student_id")
	var student Models.Student						// Check for availability of student
	err := Models.GetStudentByID(&student, studentId)
	if err==nil {
		var marks []Models.Marks // Check for availability of marks
		err = Models.DeleteAllMarks(&marks,studentId)
		if err==nil{
			c.JSON(http.StatusOK, gin.H{"All marks for Student " + studentId: " are deleted"})
		}else{
			c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		}
	}else{
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}

}

*/
