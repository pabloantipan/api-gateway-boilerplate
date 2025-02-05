package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pabloantipan/go-api-gateway-poc/config/constants"
)

type Config struct {
	Port                    string
	FirebaseCredentialsFile string
	FirebaseWebAPIKey       string
	Services                map[string]constants.ServiceConfig
	AuthWhitelistedPaths    []string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	return &Config{
		Port:                    os.Getenv("PORT"),
		FirebaseCredentialsFile: os.Getenv("FIREBASE_SERVICE_ACCOUNT_FILE"),
		FirebaseWebAPIKey:       os.Getenv("FIREBASE_WEB_API_KEY"),
		Services:                constants.SERVICE_PATHS,
		AuthWhitelistedPaths:    constants.AUTH_WHITELIST_PATHS,
	}, nil
}
