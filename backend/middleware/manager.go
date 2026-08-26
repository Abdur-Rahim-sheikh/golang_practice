package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) With(middlewares ...Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		h := handler
		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			h = middleware(h)
		}
		return h
	}
}
