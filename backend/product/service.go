package product

import (
	"ecommerce/domain"
)

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{productRepo: productRepo}
}

func (svc *service) Add(product domain.Product) (*domain.Product, error) {
	return svc.productRepo.Add(product)
}
func (svc *service) Get(id int) *domain.Product {
	return svc.productRepo.Get(id)
}
func (svc *service) List() []*domain.Product {
	return svc.productRepo.List()
}
func (svc *service) Delete(id int) error {
	return svc.productRepo.Delete(id)
}
func (svc *service) Update(product domain.Product) (*domain.Product, error) {
	return svc.productRepo.Update(product)
}
