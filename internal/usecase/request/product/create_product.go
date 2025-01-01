package request

import (
	"shozai_model1/internal/domain/table_model"
)

type CreateProduct struct {
	Name        *string `json:"name" validate:"required"`
	Code        *string `json:"code" validate:"required"`
	Description *string `json:"description" validate:"required"`
	CompanyID   *int64  `json:"companyId" validate:"required"`
}

func (r *CreateProduct) ToProduct() *table_model.Product {

	return &table_model.Product{
		Name:        *r.Name,
		Code:        *r.Code,
		Description: *r.Description,
		CompanyID:   int(*r.CompanyID),
	}
}
