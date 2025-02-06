package middleware

import "net/http"

type MiddlewareChain struct {
	middlewares []func(http.Handler) http.Handler
}

func NewChain() *MiddlewareChain {
	return &MiddlewareChain{}
}

func (c *MiddlewareChain) Add(middleware func(http.Handler) http.Handler) *MiddlewareChain {
	c.middlewares = append(c.middlewares, middleware)
	return c
}

func (c *MiddlewareChain) Then(handler http.Handler) http.Handler {
	for i := range c.middlewares {
		handler = c.middlewares[len(c.middlewares)-1-i](handler)
	}
	return handler
}
