package repo

import (
	"fmt"
	"log"

	"ecommerce/domain"
	"ecommerce/user"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	repo := &userRepo{db: dbCon}
	// generateDefaultUsers(repo)
	return repo
}

func (self *userRepo) Add(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (first_name, last_name, email, password, is_shop_owner)
		VALUES (:first_name, :last_name, :email, :password, :is_shop_owner)
		RETURNING id
	`
	rows, err := self.db.NamedQuery(query, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var userID int
	if rows.Next() {
		rows.Scan(&userID)
	}
	user.ID = userID
	return &user, nil
}

func (self *userRepo) Get(id int) *domain.User {
	var user domain.User
	query := `SELECT first_name, last_name, email, password, is_shop_owner FROM users WHERE id = $1`
	err := self.db.Get(&user, query, id)
	if err != nil {
		return nil
	}
	return &user
}

func (self *userRepo) GetByMail(email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT first_name, last_name, email, password,is_shop_owner FROM users WHERE email = $1`
	err := self.db.Get(&user, query, email)
	if err != nil {
		return nil, fmt.Errorf("user with email %s not found", email)
	}
	return &user, nil
}

func (self *userRepo) List() []*domain.User {
	var users []*domain.User
	query := `SELECT first_name, last_name, email, password, is_shop_owner FROM users`
	err := self.db.Select(&users, query)
	if err != nil {
		log.Println(err)
		return nil
	}
	return users
}

func (self *userRepo) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := self.db.Exec(query, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}
	return nil
}

func (self *userRepo) Update(user domain.User) (*domain.User, error) {
	query := `
		UPDATE users SET
			first_name = :first_name,
			last_name = :last_name,
			email = :email,
			password = :password,
			is_shop_owner = :is_shop_owner
		WHERE id = :id
	`
	result, err := self.db.NamedExec(query, user)
	if err != nil {
		return nil, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, fmt.Errorf("user with id %d not found", user.ID)
	}
	return &user, nil
}

// func generateDefaultUsers(repo *userRepo) {
// 	default_user1 := User{
// 		FirstName:   "Pansy",
// 		LastName:    "Schaefer",
// 		Email:       "test@gmai.com",
// 		Password:    "test123",
// 		IsShopOwner: true,
// 	}

// 	repo.Add(default_user1)
// }
