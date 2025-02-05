package main

import (
	"log"
	"net/http"

	"github.com/pabloantipan/go-api-gateway-poc/config"
	"github.com/pabloantipan/go-api-gateway-poc/internal/data/repository"
	"github.com/pabloantipan/go-api-gateway-poc/internal/presentation/handler"
	"github.com/pabloantipan/go-api-gateway-poc/internal/service"
)

func main() {
	// Initialize components
	cfg := config.NewConfig()
	routeRepo := repository.NewRouteRepository(cfg)
	gatewayService := service.NewGatewayService(routeRepo)
	gatewayHandler := handler.NewGatewayHandler(gatewayService)

	// Setup router
	http.Handle("/", gatewayHandler)

	// Start server
	log.Printf("Starting API Gateway on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
