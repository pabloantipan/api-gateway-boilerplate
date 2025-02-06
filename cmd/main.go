package main

import (
	"context"
	"log"
	"net/http"

	"github.com/pabloantipan/go-api-gateway-poc/config"
	"github.com/pabloantipan/go-api-gateway-poc/internal/infrastructure/cloud"
	"github.com/pabloantipan/go-api-gateway-poc/internal/presentation/handler"
	"github.com/pabloantipan/go-api-gateway-poc/internal/presentation/middleware"
	"github.com/pabloantipan/go-api-gateway-poc/internal/service"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Cloud dependenicies
	firebaseClient, err := cloud.NewFirebaseClient(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	logger, err := cloud.NewCloudLogger(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Cloud Logger: %v", err)
	}

	// Initialize Repositories

	// Initialize Services
	authService := service.NewAuthService(firebaseClient)
	gatewayService := service.NewGatewayService()

	// Initialize middleware
	rateLimiter := middleware.NewRateLImitMiddleware(cfg)
	requestLoggerMiddleware := middleware.NewRequestLoggerMiddleware(logger)
	responseLoggerMiddleware := middleware.NewResponseLoggerMiddleware(logger)
	authMiddleware := middleware.NewAuthMiddleware(authService, cfg.AuthWhitelistedPaths)

	// Initialize Handlers
	authHandler := handler.NewAuthHandler(authService)
	gatewayHandler := handler.NewGatewayHandler(gatewayService)

	healthChain := middleware.NewChain().Add(rateLimiter.Handle)
	healthHandler := healthChain.Then(handler.NewHealthHandler("1.0.0"))

	protectedChain := middleware.NewChain().
		Add(requestLoggerMiddleware.Handle).
		Add(responseLoggerMiddleware.Handle).
		Add(rateLimiter.Handle).
		Add(authMiddleware.Handle)

	protectedHandler := protectedChain.
		Then(gatewayHandler)

	// Setup router
	mux := http.NewServeMux()

	// Health check
	mux.Handle("/health", healthHandler)

	// Auth endpoints
	mux.HandleFunc("/login/v1/auth/login", authHandler.Login)

	// Protected endpoints
	mux.Handle("/api/v1/", protectedHandler)

	// Start server
	log.Printf("Starting API Gateway on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
