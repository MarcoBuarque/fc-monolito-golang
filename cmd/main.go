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

	"log"

	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	// config.InitLog()
}

func main() {
	response := config.GetDB()
	fmt.Println("RESPONSE", response)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	done := make(chan os.Signal)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.SetupServer().ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// log.Errorf("error: %s\n", err)
			// TODO: ADD LOGS
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

	if err := server.Shutdown(ctx); err != nil {
		// logger.Fatal("Error while shutting down Server. Initiating force shutdown...", logger.String("Error", err.Error()))
		// TODO: ADD LOGS
	} else {
		// logger.Info("Server exiting")
		// TODO: ADD LOGS
	}
}
