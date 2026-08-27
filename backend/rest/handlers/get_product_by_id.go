package handlers

import (
	"ecommerce/product"
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

	for _, product := range product.GetProducts() {
		if product.ID == id {
			utils.SendData(w, product, 200)
			return
		}
	}

	http.Error(w, "Product with this product id not found", 400)
}
