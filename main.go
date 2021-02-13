package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/justpoypoy/go-base/infrastructure"
)

var timeout = 30 * time.Second

func main() {
	logger := infrastructure.NewLogger()
	infrastructure.Load(logger)

	sqlHandler, err := infrastructure.NewSQL()
	if err != nil {
		logger.LogError("%s", err)
	}

	var port = os.Getenv("APP_PORT")
	if port == "" {
		port = "8088"
	}

	server := infrastructure.Dispatch(
		fmt.Sprintf(":%s", port),
		logger,
		sqlHandler,
	)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed listen server on: %s\n", err)
		}
	}()
	graceShutdown(server)
}

func graceShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
	os.Exit(0)
}
