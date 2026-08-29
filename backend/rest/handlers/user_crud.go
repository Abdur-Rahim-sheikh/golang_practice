package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.GetUsers(), http.StatusOK)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// r.Body => description, imageUrl, price, title =>
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid json", http.StatusBadRequest)
		return
	}
	created_user := newUser.Add()
	utils.SendData(w, created_user, http.StatusCreated)
}
