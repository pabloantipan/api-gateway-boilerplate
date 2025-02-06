package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pabloantipan/go-api-gateway-poc/config"
	"github.com/pabloantipan/go-api-gateway-poc/internal/infrastructure/proxy"
)

type gatewayService struct {
	routerConfig *config.RouterConfig
	proxyFactory *proxy.ProxyFactory
}

func NewGatewayService() GatewayService {
	return &gatewayService{
		routerConfig: config.NewRouterConfig(),
		proxyFactory: proxy.NewProxyFactory(),
	}
}

func (s *gatewayService) ProxyRequest(w http.ResponseWriter, r *http.Request) error {
	path := strings.TrimSuffix(r.URL.Path, "/")
	route, exists := s.routerConfig.GetRoute(path)
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
