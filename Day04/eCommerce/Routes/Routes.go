
package Routes

import (
	"eCommerce/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	//authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
	//	"max": "639",
	//}))

	gp := router.Group("/")
	{
		customerRoutes := gp.Group("/customer")
		{
			customerRoutes.POST("/add", Controllers.CreateCustomerAccount)
			customerRoutes.GET("/order/:t_id", Controllers.CheckOrderByID)
			//customerRoutes.PATCH("/:id", Controllers.BuyProduct)

		}
		retailerRoutes := gp.Group("/retailer")
		{
			retailerRoutes.POST("/add", Controllers.CreateRetailerAccount)
			retailerRoutes.PATCH("/add", Controllers.AddProduct)
			retailerRoutes.PATCH("/:id", Controllers.PatchProduct)
			retailerRoutes.GET("/:id", Controllers.GetRHistoryByID)
		}
		productRoutes := gp.Group("/product")
		{
			productRoutes.GET("/", Controllers.GetAllProducts)
			productRoutes.GET("/:id", Controllers.GetProductByID)
		}
	}

	return router
}
