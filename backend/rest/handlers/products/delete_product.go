package handlers

import (
	"ecommerce/database"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	err = database.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Product with this product id not found", 400)

	}

	w.WriteHeader(http.StatusNoContent)
}
