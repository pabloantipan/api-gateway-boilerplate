package handler

import (
	"net/http"

	"github.com/pabloantipan/go-api-gateway-poc/internal/service"
)

type GatewayHandler struct {
	service service.GatewayService
}

func NewGatewayHandler(svc service.GatewayService) *GatewayHandler {
	return &GatewayHandler{
		service: svc,
	}
}

func (h *GatewayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.service.ProxyRequest(w, r)
	if err != nil {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}
}
