package request

import (
	"shozai_model1/internal/domain/table_model"
)

type UpdateProduct struct {
	ID      *int64  `path:"id" validate:"required"`
	Name    *string `json:"name" validate:"required,max=255"`
	Version *int64  `json:"version" validate:"required,gt=0"`
}

func (r *UpdateProduct) ToProduct() (*table_model.Product, error) {
	m := &table_model.Product{
		BaseModel: table_model.BaseModel{
			ID:      int(*r.ID),
			Version: int(*r.Version),
		},
		Name: *r.Name,
	}

	return m, nil
}
