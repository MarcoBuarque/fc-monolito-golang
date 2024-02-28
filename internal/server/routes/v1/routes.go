package v1

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/v1")
	products := api.Group("/products")
	// products.GET("")
	// products.GET("/:productID")
	products.POST("/", AddProduct)

	// clients := api.Group("clients")
	// clients.POST("")
}
