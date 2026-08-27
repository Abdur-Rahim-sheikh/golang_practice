package handlers

import (
	"ecommerce/product"
	"ecommerce/utils"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, product.GetProducts(), 200)
}
