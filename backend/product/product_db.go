package product

var productManager ProductManager

func GetProducts() []Product {
	return productManager.items
}

func AddProducts(products ...Product) {
	for _, product := range products {
		productManager.Add(product)
	}
}

func GetProduct(id int) (*Product, error) {
	return productManager.GetById(id)
}

func DeleteProduct(id int) error {
	return productManager.Delete(id)
}

func UpdateProduct(id int, product Product) (Product, error) {
	return productManager.UpdateProduct(id, product)
}
