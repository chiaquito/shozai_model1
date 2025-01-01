package request

import "shozai_model1/internal/domain/table_model"

type BulkCreateProduct struct {
	Products []CreateProduct `json:"companies" validate:"required"`
}

func (req *BulkCreateProduct) ToProducts() ([]*table_model.Product, error) {
	result := make([]*table_model.Product, len(req.Products))

	for i, v := range req.Products {
		m := v.ToProduct()
		result[i] = m
	}
	return result, nil
}
