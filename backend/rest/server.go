package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{productHandler: productHandler, userHandler: userHandler}
}

func (server *Server) Start(conf *config.Config) {
	manager := middlewares.NewManager()
	mux := http.NewServeMux()
	middlewares := middlewares.NewMiddlewares(conf)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	routerHandler := manager.With(
		middlewares.CorsPreflight,
		middlewares.Logger,
	)(mux)

	addr := ":" + strconv.Itoa(conf.HttpPort)
	fmt.Println("Server running on " + addr)

	err := http.ListenAndServe(addr, routerHandler)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
