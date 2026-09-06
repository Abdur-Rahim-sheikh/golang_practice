package repo

import (
	"fmt"
	"log"

	"ecommerce/domain"
	"ecommerce/product"
	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) product.ProductRepo {
	repo := &productRepo{db: db}
	// generateDefaultProducts(repo)
	return repo
}

func (self *productRepo) Add(product domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (title, description, price, img_url)
		VALUES (:title, :description, :price, :img_url)
		RETURNING id
	`
	rows, err := self.db.NamedQuery(query, product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var productID int
	if rows.Next() {
		rows.Scan(&productID)
	}
	product.ID = productID
	return &product, nil
}

func (self *productRepo) Get(id int) *domain.Product {
	var product domain.Product
	query := `SELECT id, title, description, price, img_url FROM products WHERE id = $1`
	err := self.db.Get(&product, query, id)
	if err != nil {
		return nil
	}
	return &product
}

func (self *productRepo) List() []*domain.Product {
	var products []*domain.Product
	query := `SELECT id, title, description, price, img_url FROM products`
	err := self.db.Select(&products, query)
	if err != nil {
		log.Println(err)
		return nil
	}
	return products
}

func (self *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	result, err := self.db.Exec(query, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("product with id %d not found", id)
	}
	return nil
}

func (self *productRepo) Update(product domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products SET
			title = :title,
			description = :description,
			price = :price,
			img_url = :img_url
		WHERE id = :id
	`
	result, err := self.db.NamedExec(query, product)
	if err != nil {
		return nil, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, fmt.Errorf("product with id %d not found", product.ID)
	}
	return &product, nil
}

// func generateDefaultProducts(repo *productRepo) {
// 	prd1 := Product{
// 		Title:       "Orange",
// 		Description: "Orange is orange, I love orange",
// 		Price:       100,
// 		ImgUrl:      "https://example.com/orange.jpg",
// 	}
// 	repo.Add(prd1)
// }
