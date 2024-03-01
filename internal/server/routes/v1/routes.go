package v1

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		productManagement := v1.Group("/product-management")
		{
			productManagement.POST("/", CreateProduct)
			productManagement.PATCH("/:productID", UpdateSalesPrice)
		}

		storeCatalog := v1.Group("/store-catalog")
		{
			storeCatalog.GET("/products", ListProducts)
			storeCatalog.GET("/products/:productID", GetProduct)
		}

		clientManagement := v1.Group("/client-management")
		{
			clientManagement.POST("/", CreateClient)
			// clientManagement.GET("", )
			clientManagement.GET("/:clientID", GetClient)
		}
	}
}
