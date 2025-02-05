package repository

import (
	"fmt"
	"strings"

	"github.com/pabloantipan/go-api-gateway-poc/config"
	"github.com/pabloantipan/go-api-gateway-poc/internal/data/model"
)

type RouteRepository interface {
	GetRoute(path string) (*model.Route, error)
	AddRoute(route *model.Route) error
}

type routeRepository struct {
	routes map[string]*model.Route
}

func NewRouteRepository(config *config.Config) RouteRepository {
	repo := &routeRepository{
		routes: make(map[string]*model.Route),
	}

	// Initialize routes from config
	for serviceName, svc := range config.Services {
		for _, path := range svc.Paths {
			repo.routes[path] = &model.Route{
				ServiceName: serviceName,
				Path:        path,
				TargetURL:   svc.BaseURL,
			}
		}
	}

	return repo
}

func (r *routeRepository) GetRoute(path string) (*model.Route, error) {
	for routePath, route := range r.routes {
		if strings.HasPrefix(path, routePath) {
			return route, nil
		}
	}
	return nil, fmt.Errorf("no route found for path: %s", path)
}

func (r *routeRepository) AddRoute(route *model.Route) error {
	if _, exists := r.routes[route.Path]; exists {
		return fmt.Errorf("route already exists for path: %s", route.Path)
	}
	r.routes[route.Path] = route
	return nil
}
