package service

import "net/http"

type GatewayService interface {
	ProxyRequest(w http.ResponseWriter, r *http.Request) error
}
