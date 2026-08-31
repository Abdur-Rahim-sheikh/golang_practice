package middlewares

import "ecommerce/config"

type Middlewares struct {
	Conf *config.Config
}

func NewMiddlewares(conf *config.Config) *Middlewares {
	return &Middlewares{Conf: conf}
}
