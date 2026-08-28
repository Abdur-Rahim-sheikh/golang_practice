package product

import "fmt"

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

func (m *ProductManager) GetById(id int) (*Product, error) {
	for idx := range m.items {
		if m.items[idx].ID == id {
			return &m.items[idx], nil
		}
	}
	return nil, fmt.Errorf("product with id %d not found", id)
}
func (m *ProductManager) UpdateProduct(id int, product Product) (Product, error) {
	for idx := range m.items {
		if m.items[idx].ID == id {
			m.items[idx] = product
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("product with id %d not found", id)
}

func (m *ProductManager) Delete(id int) error {
	for idx := range m.items {
		if m.items[idx].ID == id {
			lastIdx := len(m.items) - 1
			m.items[idx] = m.items[lastIdx]
			m.items = m.items[:lastIdx]
			return nil
		}
	}
	return fmt.Errorf("product with id %d not found", id)
}
