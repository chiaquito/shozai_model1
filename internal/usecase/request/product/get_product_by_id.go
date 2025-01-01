package request

import "shozai_model1/internal/domain/table_model"

type GetProductByID struct {
	ID int64 `param:"id" validate:"required"`
}

func (req GetProductByID) ToProduct() *table_model.Product {
	return &table_model.Product{
		BaseModel: table_model.BaseModel{
			ID: int(req.ID),
		},
	}
}
