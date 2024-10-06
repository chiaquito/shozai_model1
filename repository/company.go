package repository

import (
	"gorm.io/gorm"

	"shozai_model1/domain/table_model"
)

type Company interface {
	Get(cond *table_model.Company) (*table_model.Company, error)
	List(cond *table_model.Company) ([]*table_model.Company, error)
}

type company struct {
	db *gorm.DB
}

func NewCompany(db *gorm.DB) *company {
	return &company{db: db}
}

func (r *company) Get(cond *table_model.Company) (*table_model.Company, error) {
	var company table_model.Company
	if err := r.db.Take(&company, cond).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *company) List(cond *table_model.Company) ([]*table_model.Company, error) {
	var companies []*table_model.Company
	if err := r.db.Find(&companies, cond).Error; err != nil {
		return nil, err
	}
	return companies, nil
}
