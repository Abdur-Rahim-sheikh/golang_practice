package database

import "fmt"

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var users []User

func (u User) Add() User {
	if u.ID != 0 {
		return u
	}
	u.ID = len(users)
	users = append(users, u)
	return u
}

func GetUsers() []User {
	return users
}

func GetUserByMail(email string) (User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User with this email %s not found", email)
}
