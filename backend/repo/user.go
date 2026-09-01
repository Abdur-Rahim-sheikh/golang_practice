package repo

import ("fmt")

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Add(user User) (*User, error)
	Get(id int) *User
	GetByMail(email string) (*User, error)
	List() []*User
	Delete(id int) error
	Update(user User) (*User, error)
}

type userRepo struct {
	items []*User
}

func NewUserRepo() UserRepo {
	repo := &userRepo{}
	generateDefaultUsers(repo)
	return repo
}

func (self userRepo) Add(user User) (*User, error) {
	user.ID = len(self.items) + 1
	self.items = append(self.items, &user)
	return &user, nil
}
func (self userRepo) Get(id int) *User {
	for idx := range self.items {
		if self.items[idx].ID == id {
			return self.items[idx]
		}
	}
	return nil
}
func (self userRepo) GetByMail(email string) (*User, error) {
	for idx := range self.items {
		if self.items[idx].Email == email {
			return self.items[idx], nil
		}
	}
	return nil, fmt.Errorf("user with email %s not found", email)
}
func (self userRepo) List() []*User {
	return self.items
}
func (self userRepo) Delete(id int) error {
	for idx := range self.items {
		if self.items[idx].ID == id {
			lastIdx := len(self.items) - 1
			self.items[idx] = self.items[lastIdx]
			self.items = self.items[:lastIdx]
			return nil
		}
	}
	return fmt.Errorf("user with id %d not found", id)
}
func (self userRepo) Update(user User) (*User, error) {
	for idx := range self.items {
		if self.items[idx].ID == user.ID {
			self.items[idx] = &user
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with id %d not found", user.ID)
}

func generateDefaultUsers(repo *userRepo) {
	default_user1 := User{
		FirstName:   "Pansy",
		LastName:    "Schaefer",
		Email:       "test@gmai.com",
		Password:    "test123",
		IsShopOwner: true,
	}

	repo.items = append(repo.items, &default_user1)
}