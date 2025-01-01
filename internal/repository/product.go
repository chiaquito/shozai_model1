package repository

import (
	"strings"

	"gorm.io/gorm"

	"shozai_model1/internal/domain/integrated_model"
	"shozai_model1/internal/domain/table_model"
	"shozai_model1/internal/pkg/errors"
)

type Product interface {
	Get(cond *table_model.Product) (*table_model.Product, error)
	List(cond *table_model.Product) ([]*table_model.Product, error)
	Create(m *table_model.Product) (*table_model.Product, error)
	BulkCreate(m []*table_model.Product) ([]*table_model.Product, error)
	Update(m *table_model.Product) (*table_model.Product, error)
	ListRaw(cond *table_model.Product) ([]*integrated_model.Product, error)
}

type product struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) *product {
	return &product{db: db}
}

func (r *product) Get(cond *table_model.Product) (*table_model.Product, error) {
	var product table_model.Product
	if err := r.db.Preload("Company").Take(&product, cond).Error; err != nil {
		if errors.IsRecordNotFoundError(err) {
			return nil, errors.MakeRecordNotFoundError(err)
		}
		return nil, errors.MakeSystemError(err)
	}
	return &product, nil
}

func (r *product) List(cond *table_model.Product) ([]*table_model.Product, error) {
	var products []*table_model.Product
	if err := r.db.Preload("Company").Find(&products, cond).Error; err != nil {
		return nil, errors.MakeSystemError(err)
	}
	return products, nil
}

func (r *product) ListRaw(cond *table_model.Product) ([]*integrated_model.Product, error) {
	var products []*integrated_model.Product
	var sb strings.Builder
	var args []interface{}

	sb.WriteString(`
	SELECT 
	product.id as product_id, 
	product.name as product_name,
	product.code as product_code,
	product.description as product_description,
	product.company_id as product_company_id,
	company.name as company_name,
	company.established_date as company_established_date
	FROM product 
	INNER JOIN company ON company.id = product.company_id
	WHERE 1=1
	`)

	if cond.BaseModel.ID != 0 {
		sb.WriteString(" AND product.id = ?")
		args = append(args, cond.BaseModel.ID)
	}
	if cond.Name != "" {
		sb.WriteString(" AND product.name = ?")
		args = append(args, cond.Name)
	}
	if cond.Code != "" {
		sb.WriteString(" AND product.code = ?")
		args = append(args, cond.Code)
	}
	if cond.CompanyID != 0 {
		sb.WriteString(" AND product.company_id = ?")
		args = append(args, cond.CompanyID)
	}

	if err := r.db.Raw(sb.String(), args...).Scan(&products).Error; err != nil {
		return nil, errors.MakeSystemError(err)
	}
	return products, nil
}

func (r *product) Create(m *table_model.Product) (*table_model.Product, error) {
	if err := r.db.Create(m).Error; err != nil {
		return nil, errors.MakeSystemError(err)
	}
	return m, nil
}

func (r *product) BulkCreate(m []*table_model.Product) ([]*table_model.Product, error) {
	if err := r.db.Create(m).Error; err != nil {
		return nil, errors.MakeSystemError(err)
	}
	return m, nil
}

func (r *product) Update(m *table_model.Product) (*table_model.Product, error) {
	cond := &table_model.Product{BaseModel: table_model.BaseModel{ID: m.ID}}
	if err := r.db.Take(&table_model.Product{}, cond).Error; err != nil {
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
