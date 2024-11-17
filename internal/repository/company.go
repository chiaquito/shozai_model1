package repository

import (
	// "errors"

	"gorm.io/gorm"

	"shozai_model1/internal/domain/table_model"
	"shozai_model1/internal/pkg/errors"
)

type Company interface {
	Get(cond *table_model.Company) (*table_model.Company, error)
	List(cond *table_model.Company) ([]*table_model.Company, error)
	Create(m *table_model.Company) (*table_model.Company, error)
	BulkCreate(m []*table_model.Company) ([]*table_model.Company, error)
	Update(m *table_model.Company) (*table_model.Company, error)
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
		if errors.IsRecordNotFoundError(err) {
			return nil, errors.MakeRecordNotFoundError(err)
		}
		return nil, errors.MakeSystemError(err)
	}
	return &company, nil
}

func (r *company) List(cond *table_model.Company) ([]*table_model.Company, error) {
	var companies []*table_model.Company
	if err := r.db.Find(&companies, cond).Error; err != nil {
		return nil, errors.MakeSystemError(err)
	}
	return companies, nil
}

func (r *company) Create(m *table_model.Company) (*table_model.Company, error) {
	if err := r.db.Create(m).Error; err != nil {
		return nil, errors.MakeSystemError(err)
	}
	return m, nil
}

func (r *company) BulkCreate(m []*table_model.Company) ([]*table_model.Company, error) {
	if err := r.db.Create(m).Error; err != nil {
		return nil, errors.MakeSystemError(err)
	}
	return m, nil
}

func (r *company) Update(m *table_model.Company) (*table_model.Company, error) {
	cond := &table_model.Company{BaseModel: table_model.BaseModel{ID: m.ID}}
	if err := r.db.Take(&table_model.Company{}, cond).Error; err != nil {
		if errors.IsRecordNotFoundError(err) {
			return nil, errors.MakeRecordNotFoundError(err)
		}
		return nil, errors.MakeSystemError(err)
	}
	if cond.Version != m.Version {
		return nil, errors.MakeOptimisticLockingError(m.ID)
	}
	m.Version++

	if err := r.db.Debug().Model(cond).Updates(m).Error; err != nil {
		return nil, errors.MakeSystemError(err)
	}

	return m, nil
}
