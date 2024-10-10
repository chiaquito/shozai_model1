package request

import "shozai_model1/domain/table_model"

type GetCompanyByID struct {
	ID int64 `param:"id" validate:"required"`
}

func (req GetCompanyByID) ToCompany() *table_model.Company {
	return &table_model.Company{
		BaseModel: table_model.BaseModel{
			ID: int(req.ID),
		},
	}
}
