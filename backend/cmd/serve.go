package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Serve() {
	conf := config.GetConfig()
	db, err := db.NewConnection(*conf)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	middlewares := middlewares.NewMiddlewares(conf)
	productRepo := repo.NewProductRepo(db)
	userRepo := repo.NewUserRepo(db)
	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(middlewares, userRepo)
	server := rest.NewServer(productHandler, userHandler)
	server.Start(conf)
}
