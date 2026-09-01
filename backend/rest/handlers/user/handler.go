package user

import (
	"ecommerce/repo"
	"ecommerce/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	userRepo   repo.UserRepo
}

func NewHandler(middlewares *middlewares.Middlewares, userRepo repo.UserRepo) *Handler {
	return &Handler{middlewares: middlewares, userRepo: userRepo}

}
