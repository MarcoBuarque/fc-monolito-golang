package server

import (
	"net/http"

	docs "github.com/MarcoBuarque/fc-monolito-golang/cmd/api/swagger"

	v1 "github.com/MarcoBuarque/fc-monolito-golang/internal/server/routes/v1"
	"github.com/MarcoBuarque/fc-monolito-golang/pkg"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupHandler() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, pkg.BuildCustomResponse("ROUTE_NOT_FOUND", "This route does not exist"))
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1.SetupRoutes(router)

	return router
}

func SetupServer() *http.Server {

	return &http.Server{
		Addr:    ":8080",
		Handler: setupHandler(),
	}
}
