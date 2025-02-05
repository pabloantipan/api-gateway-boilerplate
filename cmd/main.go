package main

import (
	"context"
	"log"
	"net/http"

	"github.com/pabloantipan/go-api-gateway-poc/config"
	"github.com/pabloantipan/go-api-gateway-poc/internal/data/repository"
	"github.com/pabloantipan/go-api-gateway-poc/internal/infrastructure/cloud"
	"github.com/pabloantipan/go-api-gateway-poc/internal/presentation/handler"
	"github.com/pabloantipan/go-api-gateway-poc/internal/presentation/middleware"
	"github.com/pabloantipan/go-api-gateway-poc/internal/service"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()

	// Initialize Cloud dependenicies
	firebaseClient, err := cloud.NewFirebaseClient(ctx, cfg.FirebaseCredentialsFile)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	// Initialize Repositories
	routeRepo := repository.NewRouteRepository(cfg)

	// Initialize Services
	authService := service.NewAuthService(firebaseClient)
	gatewayService := service.NewGatewayService(routeRepo)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(authService, cfg.AuthWhitelistedPaths)

	// Initialize Handlers
	gatewayHandler := handler.NewGatewayHandler(gatewayService)

	// Setup middleware
	finalHandler := authMiddleware.Handle(gatewayHandler)

	// Setup router
	http.Handle("/", finalHandler)
	// http.Handle("/", finalHandler)

	// Start server
	log.Printf("Starting API Gateway on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
