package repository

import "gorm.io/gorm"

type Repository struct {
	Company *company
	Product *product
}

func New(db *gorm.DB) *Repository {

	company := NewCompany(db)
	product := NewProduct(db)

	return &Repository{
		Company: company,
		Product: product,
	}
}
