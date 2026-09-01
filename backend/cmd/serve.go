package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Serve() {
	conf := config.GetConfig()
	middlewares := middlewares.NewMiddlewares(conf)
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()
	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(middlewares, userRepo)
	server := rest.NewServer(productHandler, userHandler)
	server.Start(conf)
}
