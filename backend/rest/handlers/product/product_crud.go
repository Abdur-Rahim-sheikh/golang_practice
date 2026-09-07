package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type RequestProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"imageUrl"`
	Price       float64 `json:"price"`
}

type Pagination struct {
	Data       []*domain.Product `json:"data"`
	Limit      int64             `json:"limit"`
	Page       int64             `json:"page"`
	TotalItems int64             `json:"total_items"`
	TotalPagse int64             `json:"total_pages"`
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	page, _ := strconv.ParseInt(pageStr, 10, 32)
	limit, _ := strconv.ParseInt(limitStr, 10, 32)

	// if error at page, it will have a default zero value,
	// and this is what we want at default case
	if limit == 0 {
		limit = 10
	}
	product_list := h.svc.List(page, limit)
	cnt, err := h.svc.Count()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	paginated_data := Pagination{
		Data:       product_list,
		Page:       page,
		Limit:      limit,
		TotalItems: cnt,
		TotalPagse: cnt / limit,
	}
	utils.SendData(w, http.StatusOK, paginated_data)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// r.Body => description, imageUrl, price, title =>

	var newProduct RequestProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid json", 400)
		return
	}

	product, err := h.svc.Add(domain.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		ImgUrl:      newProduct.ImageUrl,
		Price:       newProduct.Price,
	})
	utils.SendData(w, http.StatusCreated, product)
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	err = h.svc.Delete(id)
	if err != nil {
		http.Error(w, "Product with this product id not found", 400)
		return
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

	item := h.svc.Get(id)
	if item == nil {
		http.Error(w, "Product with this product id not found", 400)
		return
	}
	utils.SendData(w, http.StatusOK, item)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	var newProduct RequestProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid json", 400)
		return
	}
	item, err := h.svc.Update(domain.Product{
		ID:          id,
		Title:       newProduct.Title,
		Description: newProduct.Description,
		ImgUrl:      newProduct.ImageUrl,
		Price:       newProduct.Price,
	})
	if err != nil {
		http.Error(w, "Product with this product id not found", 400)
		return
	}
	utils.SendData(w, http.StatusOK, item)
}
