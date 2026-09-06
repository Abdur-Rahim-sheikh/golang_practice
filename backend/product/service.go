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

func (self *service) Add(product domain.Product) (*domain.Product, error) {
	return self.productRepo.Add(product)
}
func (self *service) Get(id int) *domain.Product {
	return self.productRepo.Get(id)
}
func (self *service) List() []*domain.Product {
	return self.productRepo.List()
}
func (self *service) Delete(id int) error {
	return self.productRepo.Delete(id)
}
func (self *service) Update(product domain.Product) (*domain.Product, error) {
	return self.productRepo.Update(product)
}
