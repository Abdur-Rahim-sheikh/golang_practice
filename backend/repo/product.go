package repo

import (
	"fmt"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}
type ProductRepo interface {
	Add(product Product) (*Product, error)
	Get(id int) *Product
	List() []*Product
	Delete(id int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	items []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateDefaultProducts(repo)
	return repo
}

func (self *productRepo) Add(product Product) (*Product, error) {
	product.ID = len(self.items) + 1
	self.items = append(self.items, &product)
	return &product, nil
}
func (self *productRepo) Get(id int) *Product {
	for idx := range self.items {
		if self.items[idx].ID == id {
			return self.items[idx]
		}
	}
	return nil

}
func (self *productRepo) List() []*Product {
	return self.items
}
func (self *productRepo) Delete(id int) error {
	for idx := range self.items {
		if self.items[idx].ID == id {
			lastIdx := len(self.items) - 1
			self.items[idx] = self.items[lastIdx]
			self.items = self.items[:lastIdx]
			return nil
		}
	}
	return fmt.Errorf("product with id %d not found", id)

}
func (self *productRepo) Update(product Product) (*Product, error) {
	for idx := range self.items {
		if self.items[idx].ID == product.ID {
			self.items[idx] = &product
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product with id %d not found", product.ID)

}

func generateDefaultProducts(pr *productRepo) {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is orange, I love orange",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/pteG9T1-Pgd47tAC6az-lh1OjLB1aoxqpWUhlU37z38/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly93d3cu/cmQuY29tL3dwLWNv/bnRlbnQvdXBsb2Fk/cy8yMDE3LzEyLzAx/X29yYW5nZXNfRmlu/YWxseSVFMiU4MCU5/NEhlcmUlRTIlODAl/OTlzLVdoaWNoLSVF/MiU4MCU5Q09yYW5n/ZSVFMiU4MCU5RC1D/YW1lLUZpcnN0LXRo/ZS1Db2xvci1vci10/aGUtRnJ1aXRfNjkx/MDY0MzUzX0x1Y2t5/LUJ1c2luZXNzLmpw/Zz9maXQ9NjQwLDQy/Nw",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/k64M2mdcB9vnknWz3p_Ovw1zKux-E81TzEBt-UVDJXs/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9tZWRp/YS5pc3RvY2twaG90/by5jb20vaWQvMTE1/MjA2Nzc3Mi9waG90/by9kZWxpY2lvdXMt/cmVkLWFwcGxlcy1v/bi1yZXRhaWwtZGlz/cGxheS1hdC1zdXBl/cm1hcmtldC5qcGc_/cz02MTJ4NjEyJnc9/MCZrPTIwJmM9enZB/ckJKVmZtM1lyQlhO/ZHN4YWFXU2VMVEJU/RmVhM0VTOTg1bVhR/QXFtaz0",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is Green",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/Bh203dGWHpOJQYtYaG77ohn8ZCpH8xt0veS9QVN1FMg/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9tZWRp/YS5nZXR0eWltYWdl/cy5jb20vaWQvMjE2/ODc1MjgxNi9waG90/by9jbG9zZS11cC1v/Zi1iYW5hbmEtdHJl/ZS5qcGc_cz02MTJ4/NjEyJnc9MCZrPTIw/JmM9aXBHNWRNLUxk/R0ZEY0hta1ZOSjlJ/LUp2X2ZvdDNKY2Vz/Q0V6MlZnUjVRTT0",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Guava",
		Description: "Guava is green",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/VosOYMkoCA43ivdH_eD_232M3utY3zTMhDCCUHXR-hU/rs:fit:0:180:1:0/g:ce/aHR0cHM6Ly9jZG4u/bW9zLmNtcy5mdXR1/cmVjZG4ubmV0L1FX/SlplaWo3cGlwbnNp/YTc1ZEdOdlgtMjMw/LTgwLmpwZw",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Pomegranate",
		Description: "Pomegranate is red",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/lr54-BpcmvudejK69bknqjndUkfJQ0VOImgtugxbems/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly93d3cu/bnV0cml0aW9uYWR2/YW5jZS5jb20vd3At/Y29udGVudC91cGxv/YWRzLzIwMjMvMDgv/Y3V0LXBvbWVncmFu/YXRlLXNob3dpbmct/cmVkLXNlZWRzLmpw/Zw",
	}

	// database.AddProducts(prd1, prd2, prd3, prd4, prd5)
	pr.items = append(pr.items, &prd1)
	pr.items = append(pr.items, &prd2)
	pr.items = append(pr.items, &prd3)
	pr.items = append(pr.items, &prd4)
	pr.items = append(pr.items, &prd5)
	// default_user := User{
	// 	FirstName:   "Pansy",
	// 	LastName:    "Schaefer",
	// 	Email:       "test@gmail.com",
	// 	Password:    "test",
	// 	IsShopOwner: true,
	// }
	// default_user.Add()
}
