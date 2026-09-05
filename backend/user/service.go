package user

import (
	"ecommerce/domain"
	"fmt"
)

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{userRepo: userRepo}
}

func (svc *service) Add(user domain.User) (*domain.User, error) {
	usr, err := svc.userRepo.Add(user)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
func (svc *service) Get(id int) *domain.User {
	user := svc.userRepo.Get(id)
	return user
}
func (svc *service) Find(email string, pass string) (*domain.User, error) {
	user, err := svc.userRepo.GetByMail(email)
	if err != nil {
		return nil, err
	}
	if user.Password != pass {
		return nil, fmt.Errorf("Password not matched")
	}
	return user, nil
}
func (svc *service) List() []*domain.User {
	users := svc.userRepo.List()
	return users
}
func (svc *service) Delete(id int) error {
	return svc.userRepo.Delete(id)

}
func (svc *service) Update(user domain.User) (*domain.User, error) {
	return svc.userRepo.Update(user)
}
