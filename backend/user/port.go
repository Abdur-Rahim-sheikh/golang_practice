package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}
type UserRepo interface {
	Add(user domain.User) (*domain.User, error)
	Get(id int) *domain.User
	GetByMail(email string) (*domain.User, error)
	List() []*domain.User
	Delete(id int) error
	Update(user domain.User) (*domain.User, error)
}
