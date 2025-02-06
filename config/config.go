package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pabloantipan/go-api-gateway-poc/config/constants"
	"github.com/pabloantipan/go-api-gateway-poc/pkg/utils"
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
		RateLimitRequestsPerSecond:  utils.ParseEnvFloat64("RATE_LIMIT_REQUESTS_PER_SECOND"),
		RateLimitBurstSize:          utils.ParseEnvFloat64("RATE_LIMIT_BURST_SIZE"),
		AuthWhitelistedPaths:        constants.AUTH_WHITELIST_PATHS,
	}, nil
}
