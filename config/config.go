package config

import (
	"os"

	"strconv"

	"github.com/joho/godotenv"
	"github.com/pabloantipan/go-api-gateway-poc/config/constants"
)

type Config struct {
	Port                        string
	ProjectID                   string
	CloudLoggingCredentialsFile string
	FirebaseCredentialsFile     string
	FirebaseWebAPIKey           string
	AuthWhitelistedPaths        []string
	RateLimitRequestsPerSecond  float64
	RateLimitBurstSize          float64
}

func ParseEnvFloat64(key string) float64 {
	value, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return 0.0
	}
	return value
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	return &Config{
		Port:                        os.Getenv("PORT"),
		ProjectID:                   os.Getenv("PROJECT_ID"),
		CloudLoggingCredentialsFile: os.Getenv("LOGGING_SERVICE_ACCOUNT_FILE"),
		FirebaseCredentialsFile:     os.Getenv("FIREBASE_SERVICE_ACCOUNT_FILE"),
		FirebaseWebAPIKey:           os.Getenv("FIREBASE_WEB_API_KEY"),
		RateLimitRequestsPerSecond:  ParseEnvFloat64("RATE_LIMIT_REQUESTS_PER_SECOND"),
		RateLimitBurstSize:          ParseEnvFloat64("RATE_LIMIT_BURST_SIZE"),
		AuthWhitelistedPaths:        constants.AUTH_WHITELIST_PATHS,
	}, nil
}
