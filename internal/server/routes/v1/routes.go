package v1

import (
	checkout "github.com/MarcoBuarque/fc-monolito-golang/internal/server/routes/v1/chekout"
	clientmanagement "github.com/MarcoBuarque/fc-monolito-golang/internal/server/routes/v1/client_management"
	productmanagement "github.com/MarcoBuarque/fc-monolito-golang/internal/server/routes/v1/product_management"
	storecatalog "github.com/MarcoBuarque/fc-monolito-golang/internal/server/routes/v1/store_catalog"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")

	checkout.ConfigRoutes(v1)
	clientmanagement.ConfigRoutes(v1)
	productmanagement.ConfigRoutes(v1)
	storecatalog.ConfigRoutes(v1)
}
