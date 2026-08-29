package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid data", http.StatusBadRequest)
		return
	}
	user, err := database.GetUserByMail(reqLogin.Email)
	passwordMatched := user.Password == reqLogin.Password
	if err != nil || !passwordMatched {
		http.Error(w, "user mail or password not matched", http.StatusBadRequest)
		return
	}
	payload := utils.Payload{Sub: user.ID, Email: user.Email, IsShopOwner: user.IsShopOwner}
	token, err := utils.CreateJwt("my-secret", payload)

	if err != nil {
		http.Error(w, "token generation failed", http.StatusConflict)
	}
	utils.SendData(w, token, http.StatusCreated)
}
