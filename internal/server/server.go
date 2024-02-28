package server

import (
	"context"
	"net/http"

	v1 "github.com/MarcoBuarque/monolito/internal/server/routes/v1"
	"github.com/gin-gonic/gin"
)

func setupHandler() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1.SetupRoutes(router)

	return router
}

func SetupServer() *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: setupHandler(),
	}
}

func Shutdown(ctx context.Context) error {
	return nil
}
