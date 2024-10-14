package request

import "shozai_model1/internal/domain/table_model"

type ListCompany struct {
	ID   *int64  `query:"id"`
	Name *string `query:"name"`
}

func (req ListCompany) ToCompany() *table_model.Company {

	m := &table_model.Company{}

	if req.ID != nil {
		m.BaseModel.ID = int(*req.ID)
	}
	if req.Name != nil {
		m.Name = *req.Name
	}

	return m
}
