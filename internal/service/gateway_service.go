package service

import (
	"fmt"
	"net/http"

	"github.com/pabloantipan/go-api-gateway-poc/config"
	"github.com/pabloantipan/go-api-gateway-poc/internal/data/repository"
	"github.com/pabloantipan/go-api-gateway-poc/internal/infrastructure/proxy"
)

type gatewayService struct {
	routeRepo    repository.RouteRepository
	routerConfig *config.RouterConfig
	proxyFactory *proxy.ProxyFactory
}

func NewGatewayService(repo repository.RouteRepository) GatewayService {
	return &gatewayService{
		routeRepo:    repo,
		proxyFactory: proxy.NewProxyFactory(),
	}
}

func (s *gatewayService) ProxyRequest(w http.ResponseWriter, r *http.Request) error {
	route, exists := s.routerConfig.GetRoute(r.URL.Path)
	if !exists {
		return fmt.Errorf("no route found")
	}

	proxy, err := s.proxyFactory.GetProxy(route.TargetURL)
	if err != nil {
		return err
	}

	proxy.ServeHTTP(w, r)
	return nil
}

func (s *gatewayService) ProxyRequestFromRepo(w http.ResponseWriter, r *http.Request) error {
	route, err := s.routeRepo.GetRoute(r.URL.Path)
	if err != nil {
		return err
	}

	proxy, err := s.proxyFactory.GetProxy(route.TargetURL)
	if err != nil {
		return err
	}

	proxy.ServeHTTP(w, r)
	return nil
}
