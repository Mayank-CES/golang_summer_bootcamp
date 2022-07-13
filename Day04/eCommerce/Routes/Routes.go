package Routes

import (
	"github.com/Mayank-CES/golang_summer_bootcamp/Day04/eCommerce/Controllers"
	"github.com/Mayank-CES/golang_summer_bootcamp/Day04/eCommerce/Repository"
	"github.com/Mayank-CES/golang_summer_bootcamp/Day04/eCommerce/Services"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"max": "639",
	}))

	ctrl := Controllers.NewController(Services.NewCustomerService(Repository.NewRepo()))

	gp := authorized.Group("/")
	{
		customerRoutes := gp.Group("/customer")
		{
			customerRoutes.POST("/add", ctrl.CreateCustomerAccount)
			customerRoutes.GET("/order/:t_id", ctrl.CheckOrderByID)
			// customerRoutes.PATCH("/:id", ctrl.BuyProduct)

		}
		retailerRoutes := gp.Group("/retailer")
		{
			retailerRoutes.POST("/add", ctrl.CreateRetailerAccount)
			retailerRoutes.PATCH("/add", ctrl.AddProduct)
			retailerRoutes.PATCH("/:id", ctrl.PatchProduct)
			retailerRoutes.GET("/:id", ctrl.GetRHistoryByID)
		}
		productRoutes := gp.Group("/product")
		{
			productRoutes.GET("/", ctrl.GetAllProducts)
			productRoutes.GET("/:id", ctrl.GetProductByID)
		}
	}

	return router
}
