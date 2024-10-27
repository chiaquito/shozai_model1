package request

import "shozai_model1/internal/domain/table_model"

type BulkCreateCompany struct {
	Companies []CreateCompany `json:"companies" validate:"required"`
}

func (req *BulkCreateCompany) ToCompanies() ([]*table_model.Company, error) {
	result := make([]*table_model.Company, len(req.Companies))

	for i, v := range req.Companies {
		m, err := v.ToCompany()
		if err != nil {
			return nil, err
		}
		result[i] = m
	}
	return result, nil
}
