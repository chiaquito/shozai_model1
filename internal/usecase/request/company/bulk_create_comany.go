package request

import "shozai_model1/internal/domain/table_model"

type BulkCreateCompany struct {
	Companies []CreateCompany `json:"companies" validate:"required"`
}

func (req *BulkCreateCompany) ToCompanies() []*table_model.Company {
	m := make([]*table_model.Company, len(req.Companies))

	for i, v := range req.Companies {
		m[i] = v.ToCompany()
	}
	return m
}
