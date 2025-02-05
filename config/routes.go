package config

type Route struct {
	ServiceName string
	Path        string
	TargetURL   string
}

type RouterConfig struct {
	routes map[string]Route
}

func NewRouterConfig() *RouterConfig {
	return &RouterConfig{
		routes: map[string]Route{
			"/api/v1/players": {
				ServiceName: "players",
				Path:        "/api/v1/players",
				TargetURL:   "http://localhost:8085",
			},
			// Add new routes
			"/api/v1/users": {
				ServiceName: "users",
				Path:        "/api/v1/users",
				TargetURL:   "http://localhost:8086",
			},
		},
	}
}

func (rc *RouterConfig) GetRoute(path string) (Route, bool) {
	route, exists := rc.routes[path]
	return route, exists
}
