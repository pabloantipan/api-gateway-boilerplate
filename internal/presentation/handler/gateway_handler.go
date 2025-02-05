package handler

import (
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
	err := h.gatewayService.ProxyRequest(w, r)
	if err != nil {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}
}
