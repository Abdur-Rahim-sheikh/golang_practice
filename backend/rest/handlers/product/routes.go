package product

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	authRequired := manager.With(h.middlewares.Auth)

	mux.HandleFunc("GET /api/products", h.GetProducts)
	mux.HandleFunc("GET /api/products/{productId}", h.GetProductById)

	// auth required
	mux.Handle(
		"POST /api/products",
		authRequired(http.HandlerFunc(h.CreateProduct)),
	)

	mux.Handle(
		"PUT /api/products/{productId}",
		authRequired(http.HandlerFunc(h.UpdateProduct)),
	)

	mux.Handle(
		"DELETE /api/products/{productId}",
		authRequired(http.HandlerFunc(h.DeleteProduct)),
	)
}
