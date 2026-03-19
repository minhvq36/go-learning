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

	"github.com/minhvq36/go-learning/internal/api"
)

func main() {
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	port := 8080
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		fmt.Printf("[Main] Starting server on :%d\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[Main] Server error on :%d: %v\n", port, err)
		}
	}()

	<-ctx.Done()
	fmt.Println("[Main] Shutting down server...")

	const shutdownTimeout = 5 * time.Second

	shutdownContext, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	fmt.Printf("[Main] Shutting down server in %d seconds...\n", shutdownTimeout)

	if err := server.Shutdown(shutdownContext); err != nil {
		log.Fatalf("[Main] Server forced to shutdown: %v", err)
	}

	fmt.Println("[Main] Server exited gracefully.")
}
