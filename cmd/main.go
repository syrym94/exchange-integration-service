package main

import (
	handlers "exchange-integration-service/internal/api/http"
	"exchange-integration-service/internal/config"

	"github.com/gorilla/mux"

	"log"
	"net/http"
	grpcServer "github.com/syrym94/exchange-integration-service-client"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the router
	r := mux.NewRouter()

	// Setup routes
	handlers.RegisterRoutes(r, cfg)

	// Start the HTTP server
	go func() {
		log.Printf("Starting HTTP server on %s", cfg.ServerAddress)
		if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
			log.Fatalf("could not start HTTP server: %s", err)
		}
	}()

	// Start the gRPC server
	grpcServer.StartGRPCServer(":50051")
}
