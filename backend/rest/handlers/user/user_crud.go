package user

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, http.StatusOK,  h.userRepo.List())
}
func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	// r.Body => description, imageUrl, price, title =>
	var newUser repo.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid json", http.StatusBadRequest)
		return
	}
	created_user,err := h.userRepo.Add(newUser)
	utils.SendData(w, http.StatusCreated, created_user)
}
