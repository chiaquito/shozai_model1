package request

import "shozai_model1/internal/domain/table_model"

type CreateCompany struct {
	Name string `json:"name" validate:"required"`
}

func (req CreateCompany) ToCompany() *table_model.Company {
	return &table_model.Company{
		Name: req.Name,
	}
}
