package response

import "shozai_model1/domain/table_model"

type Company struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func ToCompany(m *table_model.Company) *Company {
	return &Company{
		ID:   int64(m.ID),
		Name: m.Name,
	}
}

func ToCompanies(m []*table_model.Company) []*Company {
	res := make([]*Company, len(m))
	for i, v := range m {
		res[i] = ToCompany(v)
	}
	return res
}
