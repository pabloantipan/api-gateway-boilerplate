package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pabloantipan/go-api-gateway-poc/internal/presentation/dto"
)

type HealthHandler struct {
	version string
}

func NewHealthHandler(version string) *HealthHandler {
	return &HealthHandler{
		version: version,
	}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := dto.HealthResponse{
		Status:    "UP",
		Version:   h.version,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Services: map[string]string{
			"players": "UP",
			// Add more services as they come
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
