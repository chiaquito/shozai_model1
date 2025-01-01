package request

import "shozai_model1/internal/domain/table_model"

type ListProduct struct {
	ID        *int64  `query:"id"`
	Name      *string `query:"name"`
	Code      *string `query:"code"`
	CompanyID *int64  `query:"companyId"`
}

func (req ListProduct) ToProduct() *table_model.Product {

	m := &table_model.Product{}

	if req.ID != nil {
		m.BaseModel.ID = int(*req.ID)
	}
	if req.Name != nil {
		m.Name = *req.Name
	}
	if req.Code != nil {
		m.Code = *req.Code
	}
	if req.CompanyID != nil {
		m.CompanyID = int(*req.CompanyID)
	}

	return m
}
