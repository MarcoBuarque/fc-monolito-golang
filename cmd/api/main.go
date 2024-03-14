package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/MarcoBuarque/fc-monolito-golang/config"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	// config.InitLog()
}

//	@title			fc-monolito APIs
//	@version		1.0
//	@description	Implementation in golang of FC monolith course.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

//	@host		localhost:8080
//	@BasePath	/api/v1

// @schemes	http https
func main() {
	response := config.GetDB()
	fmt.Println("RESPONSE", response)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	api := server.SetupServer()

	go func() {
		if err := api.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error: %s\n", err)
		}
	}()

	<-done
	log.Println("Shutting down server...")

	// Close all database connection etc....
	/*
	  Insert Code
	*/

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := api.Shutdown(ctx); err != nil {
		log.Fatal("Error while shutting down Server. Initiating force shutdown...", "Error", err.Error())
	} else {
		log.Info("Server exiting")
	}
}
