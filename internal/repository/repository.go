package repository

import "gorm.io/gorm"

type Repository struct {
	Company *company
}

func New(db *gorm.DB) *Repository {

	company := NewCompany(db)

	return &Repository{
		Company: company,
	}
}
