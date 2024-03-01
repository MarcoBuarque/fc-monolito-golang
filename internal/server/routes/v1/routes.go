package v1

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		productADM := v1.Group("/product-adm ")
		{
			productADM.POST("/", CreateProduct)
			productADM.PATCH("/:productID", UpdateSalesPrice)
		}

		productCatalog := v1.Group("/catalog")
		{
			productCatalog.GET("", ListProducts)
			productCatalog.GET("/:productID", GetProduct)
		}
	}

	// clients := api.Group("clients")
	// clients.POST("")
}
