package product

import (
	"ecommerce/domain"
)

type Service interface {
	Add(product domain.Product) (*domain.Product, error)
	Get(id int) *domain.Product
	List(page, limit int64) []*domain.Product
	Delete(id int) error
	Update(product domain.Product) (*domain.Product, error)
}
