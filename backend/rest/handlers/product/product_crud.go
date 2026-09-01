package product

import (
	"ecommerce/utils"
	"ecommerce/repo"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w,http.StatusOK, h.productRepo.List())
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// r.Body => description, imageUrl, price, title =>

	var newProduct repo.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid json", 400)
		return
	}
	newProduct.ID = len(h.productRepo.List()) + 1
	h.productRepo.Add(newProduct)
	utils.SendData(w, http.StatusAccepted, h.productRepo.List())
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	err = h.productRepo.Delete(id)
	if err != nil {
		http.Error(w, "Product with this product id not found", 400)

	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	item := h.productRepo.Get(id)
	if item == nil {
		http.Error(w, "Product with this product id not found", 400)

	}
	utils.SendData(w, http.StatusOK, item)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// productId := r.PathValue("productId")
	// id, err := strconv.Atoi(productId)
	// if err != nil {
	// 	http.Error(w, "Please give me a valid product id", 400)
	// 	return
	// }

	var newProduct repo.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid json", 400)
		return
	}
	item, err := h.productRepo.Update(newProduct)
	if err != nil {
		http.Error(w, "Product with this product id not found", 400)

	}
	utils.SendData(w, http.StatusOK, item)
}
