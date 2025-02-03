package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Simple service configuration
type ServiceConfig struct {
	URL string
}

func main() {
	// Users service configuration
	usersService := ServiceConfig{
		URL: "http://localhost:8085",
	}

	// Create reverse proxy
	usersURL, err := url.Parse(usersService.URL)
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(usersURL)

	// Configure custom proxy handler
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy error: %v", err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Service unavailable"))
	}

	// Initialize Gin
	r := gin.Default()

	// Middleware for basic request logging
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
		)
	}))

	// Users endpoints
	players := r.Group("/api/v1/players")
	{
		players.Any("/*path", func(c *gin.Context) {
			proxy.ServeHTTP(c.Writer, c.Request)
		})
	}

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().Unix(),
		})
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting gateway on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
