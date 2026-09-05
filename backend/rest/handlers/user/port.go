package user

import "ecommerce/domain"

type Service interface {
	Find(email string, pass string) (*domain.User, error)
	Add(user domain.User) (*domain.User, error)
	Get(id int) *domain.User
	List() []*domain.User
	Delete(id int) error
	Update(user domain.User) (*domain.User, error)
}
