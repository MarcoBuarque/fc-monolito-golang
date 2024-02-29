package v1

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("", ListProducts)
			products.GET("/:productID", GetProduct)
			products.POST("/", CreateProduct)
			products.PATCH("/:productID", UpdateSalesPrice)
		}
	}

	// clients := api.Group("clients")
	// clients.POST("")
}
