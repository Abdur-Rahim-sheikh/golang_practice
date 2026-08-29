package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// r.Body => description, imageUrl, price, title =>
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid json", 400)
		return
	}
	newProduct.ID = len(database.GetProducts()) + 1
	database.AddProducts(newProduct)
	utils.SendData(w, database.GetProducts(), 201)
}
