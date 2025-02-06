package middleware

import (
	"fmt"
	"net/http"

	"github.com/pabloantipan/go-api-gateway-poc/internal/infrastructure/cloud"
)

type RequestLoggerMiddleware struct {
	logger *cloud.CloudLogger
}

func NewRequestLoggerMiddleware(logger *cloud.CloudLogger) *RequestLoggerMiddleware {
	return &RequestLoggerMiddleware{
		logger: logger,
	}
}

func (m *RequestLoggerMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.LogRequest(r.Context(), r.Method, r.URL.Path, http.StatusOK)
		next.ServeHTTP(w, r)
	})
}

type ResponseLoggerMiddleware struct {
	logger *cloud.CloudLogger
}

func NewResponseLoggerMiddleware(logger *cloud.CloudLogger) *ResponseLoggerMiddleware {
	return &ResponseLoggerMiddleware{
		logger: logger,
	}
}

func (m *ResponseLoggerMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(w, r)

		if rw.status >= 400 {
			m.logger.LogError(r.Context(), fmt.Errorf("status code %d", rw.status), r.Method, r.URL.Path)
		}

		// Logging only errors
		// else {
		// 	m.logger.LogRequest(r.Context(), r.Method, r.URL.Path, http.StatusOK)
		// }
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}
