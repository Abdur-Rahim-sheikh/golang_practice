package product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type ProductManager struct {
	items []Product
}

func (m *ProductManager) Add(p Product) {
	m.items = append(m.items, p)
}

func (m *ProductManager) Items() []Product {
	return m.items
}
