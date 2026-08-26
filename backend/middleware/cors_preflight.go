package middleware

import (
	"net/http"
)

func allowCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

}

func CorsPreflight(next http.Handler) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		allowCors(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleAllReq)
}
