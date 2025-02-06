package config

import (
	"os"

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
		AuthWhitelistedPaths:        constants.AUTH_WHITELIST_PATHS,
	}, nil
}
