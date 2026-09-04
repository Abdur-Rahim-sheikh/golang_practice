package user

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, http.StatusOK, h.userRepo.List())
}
func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	// r.Body => description, imageUrl, price, title =>
	var newUser RequestUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid json", http.StatusBadRequest)
		return
	}
	created_user, err := h.userRepo.Add(repo.User{
		ID:          len(h.userRepo.List()) + 1,
		FirstName:   newUser.Name,
		Email:       newUser.Email,
		Password:    newUser.Password,
		IsShopOwner: newUser.IsShopOwner,
	})
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
	}
	fmt.Println("hi", created_user, h.userRepo.List())
	utils.SendData(w, http.StatusCreated, created_user)
}
