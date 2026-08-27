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
