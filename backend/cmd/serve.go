package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/database"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
)

func Serve() {
	conf := config.GetConfig()
	dbconf := config.GetDBConfig()
	db, err := database.NewConnection(*dbconf)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = database.MigrateDB(db, "./migrations")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	middlewares := middlewares.NewMiddlewares(conf)
	productRepo := repo.NewProductRepo(db)
	userRepo := repo.NewUserRepo(db)
	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(middlewares, userRepo)
	server := rest.NewServer(productHandler, userHandler)
	server.Start(conf)
}
