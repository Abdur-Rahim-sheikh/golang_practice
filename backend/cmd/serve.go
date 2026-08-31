package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Serve() {
	conf := config.GetConfig()
	middlewares := middlewares.NewMiddlewares(conf)
	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler(middlewares)
	server := rest.NewServer(productHandler, userHandler)
	server.Start(conf)
}
