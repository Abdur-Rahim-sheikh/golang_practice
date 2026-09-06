package user

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, http.StatusOK, h.svc.List())
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
	created_user, err := h.svc.Add(domain.User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		IsShopOwner: newUser.IsShopOwner,
	})
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendData(w, http.StatusCreated, created_user)
}
