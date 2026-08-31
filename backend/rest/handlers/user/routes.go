package user

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.HandleFunc("POST /api/users/login", h.Login)
	mux.HandleFunc("POST /api/users", h.AddUser)
	mux.HandleFunc("GET /api/users", h.GetUsers)
}
