package middlewares

import (
	"ecommerce/config"
	"ecommerce/utils"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	temp := func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		auth_arr := strings.Split(header, " ")
		if len(auth_arr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := utils.VerifyJwt(config.GetConfig().JwtSecret, auth_arr[1])
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(temp)
}
