package user

import (
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResLogin struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid data", http.StatusBadRequest)
		return
	}
	user, err := h.userRepo.GetByMail(reqLogin.Email)
	passwordMatched := user.Password == reqLogin.Password
	if err != nil || !passwordMatched {
		http.Error(w, "user mail or password not matched", http.StatusBadRequest)
		return
	}
	payload := utils.Claims{Sub: user.ID, Email: user.Email, IsShopOwner: user.IsShopOwner}
	token, err := utils.CreateJwt(h.middlewares.Conf.JwtSecret, payload)

	if err != nil {
		http.Error(w, "token generation failed "+err.Error(), http.StatusConflict)
		return
	}
	response := ResLogin{
		AccessToken: token,
		TokenType:   "Bearer",
	}
	utils.SendData(w, http.StatusCreated, response)
}
