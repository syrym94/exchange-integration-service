package main

import (
	grpc_server "github.com/syrym94/exchange-integration-service/internal/api/grpc"
	handlers "github.com/syrym94/exchange-integration-service/internal/api/http"
	"github.com/syrym94/exchange-integration-service/internal/config"

	"github.com/gorilla/mux"

	"log"
	"net/http"
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
	grpc_server.StartGRPCServer(cfg.GrpcAddress, cfg)
}
