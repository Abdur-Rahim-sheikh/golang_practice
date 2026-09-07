package product

import (
	"ecommerce/domain"
	productHandler "ecommerce/rest/handlers/product"
)

type Service interface {
	productHandler.Service
}

type ProductRepo interface {
	Add(product domain.Product) (*domain.Product, error)
	Get(id int) *domain.Product
	List(page, limit int64) []*domain.Product
	Count() (int64, error)
	Delete(id int) error
	Update(product domain.Product) (*domain.Product, error)
}
