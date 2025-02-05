package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type ProxyFactory struct {
	proxies map[string]*httputil.ReverseProxy
	mu      sync.RWMutex
}

func NewProxyFactory() *ProxyFactory {
	return &ProxyFactory{
		proxies: make(map[string]*httputil.ReverseProxy),
	}
}

func (f *ProxyFactory) GetProxy(targetURL string) (*httputil.ReverseProxy, error) {
	f.mu.RLock()
	proxy, exists := f.proxies[targetURL]
	f.mu.RUnlock()

	if exists {
		return proxy, nil
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	target, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	proxy = httputil.NewSingleHostReverseProxy(target)

	// Customize director
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Header.Add("X-Gateway", "api-gateway")
		req.Host = target.Host
	}

	// Add error handling
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
	}

	f.proxies[targetURL] = proxy
	return proxy, nil
}
