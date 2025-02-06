package middleware

import (
	"net/http"

	"github.com/pabloantipan/go-api-gateway-poc/config"
	"github.com/pabloantipan/go-api-gateway-poc/internal/infrastructure/ratelimit"
)

type RateLimitMiddleware struct {
	limiter *ratelimit.RateLimiter
}

func NewRateLImitMiddleware(cfg *config.Config) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		limiter: ratelimit.NewRateLimiter(cfg.RateLimitRequestsPerSecond, cfg.RateLimitBurstSize),
	}
}

func (m *RateLimitMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := getClientID(r)

		if !m.limiter.Allow(clientID) {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getClientID(r *http.Request) string {
	// Prefere athorization user if avaliaable
	if userID := r.Header.Get("X-User-ID"); userID != "" {
		return userID
	}

	// Fallback to IP address
	ip := r.Header.Get("X-Real-IP")

	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}

	if ip == "" {
		ip = r.RemoteAddr
	}

	return ip
}
