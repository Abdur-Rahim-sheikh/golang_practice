package middlewares

import (
	"log"
	"net/http"
)

func Hudai(next http.Handler) http.Handler {
	temp := func(w http.ResponseWriter, r *http.Request) {
		log.Println("ami huda middleware")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(temp)
}
