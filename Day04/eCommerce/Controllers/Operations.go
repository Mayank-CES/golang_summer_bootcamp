package Controllers

import (
	"eCommerce/Models"
	"eCommerce/Services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


//CreateCustomerAccount ... Create Customer Account
func CreateCustomerAccount(c *gin.Context) {
	var customer Models.Customer
	c.BindJSON(&customer)
	err := Services.CService.CreatCustomerAccount(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

/*
//BuyProduct ... Buy the product
func BuyProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Services.CService.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
		fmt.Println("An Error occurred while fetching the product")
	}
	c.BindJSON(&product)
	err = Services.CService.BuyProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
*/


// CheckOrderByID ... Get the product by id
func CheckOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Transaction
	err := Services.CService.CheckOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

/*------- Products ------*/

// GetAllProducts  ... Get all Products
func GetAllProducts(c *gin.Context) {
	var products []Models.Product
	err := Services.CService.GetAllProducts(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}


// GetProductByID ... Get the product by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Services.CService.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

/*---------- Retailer ----------*/

// CreateRetailerAccount  ... Create Customer Account
func CreateRetailerAccount(c *gin.Context) {
	var retailer Models.Retailer
	c.BindJSON(&retailer)
	err := Services.CService.CreatRetailerAccount(&retailer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, retailer)
	}
}

// AddProduct ... Create Customer Account
func AddProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Services.CService.AddProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//PatchProduct ... Buy the product
func PatchProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Services.CService.GetProductByID(&product, id)
	if err != nil {
		//c.JSON(http.StatusNotFound, product)
		fmt.Println("An Error occurred while fetching the product")
	}
	var updatedProduct Models.PatchProd
	c.BindJSON(&updatedProduct)
	err = Services.CService.PatchProduct(&updatedProduct,id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// GetRHistoryByID ... Get the History by id
func GetRHistoryByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Transaction
	err := Services.CService.GetRHistoryByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}


////UpdateInfo ... Update the user information
//func UpdateInfo(c *gin.Context) {
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
//}

/*
//DeleteStudent ... Delete the user
func DeleteStudent(c *gin.Context) {
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
func GetAllMarks(c *gin.Context) {
	var marks []Models.Marks							// Check for availability of marks
	err := Models.GetAllMarks(&marks)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}


//GetMarksByID ... Get the user by id
func GetMarksByID(c *gin.Context) {

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
func UpdateMarks(c *gin.Context) {
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
func DeleteMarks(c *gin.Context) {
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

func DeleteAllMarks(c *gin.Context) {
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

