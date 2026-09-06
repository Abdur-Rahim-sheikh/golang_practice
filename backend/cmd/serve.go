package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/database"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"ecommerce/user"
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
	// repos
	productRepo := repo.NewProductRepo(db)
	userRepo := repo.NewUserRepo(db)

	// domains
	userService := user.NewService(userRepo)
	productService := product.NewService(productRepo)

	// handlers
	productHandler := productHandler.NewHandler(middlewares, productService)
	userHandler := userHandler.NewHandler(middlewares, userService)
	server := rest.NewServer(productHandler, userHandler)
	server.Start(conf)
}
