package handler

import (
	"log"
	"net/http"

	"github.com/pabloantipan/go-api-gateway-poc/internal/service"
)

type GatewayHandler struct {
	gatewayService service.GatewayService
}

func NewGatewayHandler(svc service.GatewayService) *GatewayHandler {
	return &GatewayHandler{
		gatewayService: svc,
	}
}

func (h *GatewayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Gateway handler called for path: %s", r.URL.Path)

	err := h.gatewayService.ProxyRequest(w, r)
	if err != nil {
		http.Error(w, "Service not found", http.StatusNotFound)
	}
}
