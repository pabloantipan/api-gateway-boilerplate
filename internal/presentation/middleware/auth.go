package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gobwas/glob"
	"github.com/pabloantipan/go-api-gateway-poc/internal/service"
)

type AuthMiddleware struct {
	authService    service.AuthService
	whitelistPaths []glob.Glob
}

func NewAuthMiddleware(authService service.AuthService, whitelistedPaths []string) *AuthMiddleware {
	// Compile whitelist paths
	patterns := make([]glob.Glob, 0, len(whitelistedPaths))
	for _, path := range whitelistedPaths {
		if g, err := glob.Compile(path); err != nil {
			patterns = append(patterns, g)
		}
	}

	return &AuthMiddleware{
		authService:    authService,
		whitelistPaths: patterns,
	}
}

func (m *AuthMiddleware) isWhitelisted(path string) bool {
	// Clean and normalize
	cleanPath := strings.TrimSpace(path)
	cleanPath = strings.Trim(cleanPath, "/")

	// Check if path is whitelisted
	for _, pattern := range m.whitelistPaths {
		if pattern.Match(cleanPath) {
			return true
		}
	}

	return false
}

func (m *AuthMiddleware) Handle(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if path is whitelisted
		if m.isWhitelisted(r.URL.Path) {
			r.Header.Set("X-User-ID", "whitelisted")
			return
		}

		token := extractToken(r)
		if token == "" {
			http.Error(w, "Unauthorized - No token provided", http.StatusUnauthorized)
			return
		}

		userID, err := m.authService.ValidateToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Unauthorized - Invalid token", http.StatusUnauthorized)
			return
		}

		// Add user info to context
		ctx := context.WithValue(r.Context(), "userID", userID)
		r.Header.Set("X-User-ID", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return ""
	}

	parts := strings.Split(auth, "Bearer ")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}
