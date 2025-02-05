package config

type ServiceConfig struct {
	BaseURL string
	Paths   []string
}

type Config struct {
	Port     string
	Services map[string]ServiceConfig
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",
		Services: map[string]ServiceConfig{
			"players": {
				BaseURL: "http://localhost:8085",
				Paths:   []string{"/api/v1/players"},
			},
			// Add new services:
			/*
			   "users": {
			       BaseURL: "http://localhost:8086",
			       Paths:   []string{"/api/v1/users"},
			   },
			   "games": {
			       BaseURL: "http://localhost:8087",
			       Paths:   []string{"/api/v1/games"},
			   },
			   "teams": {
			       BaseURL: "http://localhost:8088",
			       Paths:   []string{"/api/v1/teams"},
			   },
			*/
		},
	}
}
