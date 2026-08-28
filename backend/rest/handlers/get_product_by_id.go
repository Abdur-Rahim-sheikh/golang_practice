package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	item, err := database.GetProduct(id)
	if err != nil {
		http.Error(w, "Product with this product id not found", 400)

	}
	utils.SendData(w, item, 200)
}
