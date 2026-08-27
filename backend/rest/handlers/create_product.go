package handlers

import (
	"ecommerce/product"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// r.Body => description, imageUrl, price, title =>
	var newProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid json", 400)
		return
	}
	newProduct.ID = len(product.GetProducts()) + 1
	product.AddProducts(newProduct)
	utils.SendData(w, product.GetProducts(), 201)
}
