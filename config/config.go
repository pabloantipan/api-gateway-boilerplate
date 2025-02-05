package config

import (
	"github.com/pabloantipan/go-api-gateway-poc/config/constants"
)

type Config struct {
	Port                    string
	FirebaseCredentialsFile string
	Services                map[string]constants.ServiceConfig
	AuthWhitelistedPaths    []string
}

func NewConfig() *Config {
	return &Config{
		Port:                    "8080",
		FirebaseCredentialsFile: "./service-accounts/firebase-service_account.json",
		Services:                constants.SERVICE_PATHS,
		AuthWhitelistedPaths:    constants.AUTH_WHITELIST_PATHS,
	}
}
