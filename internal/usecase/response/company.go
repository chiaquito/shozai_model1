package response

import (
	"shozai_model1/internal/domain/table_model"
	"shozai_model1/internal/pkg/ptr"
)

type Company struct {
	ID              int64   `json:"id"`
	Version         int64   `json:"version"`
	Name            string  `json:"name"`
	EstablishedDate *string `json:"established_date"`
	CreatedUserID   int64   `json:"created_user_id"`
}

func ToCompany(m *table_model.Company) *Company {
	result := &Company{
		ID:            int64(m.ID),
		Version:       int64(m.Version),
		CreatedUserID: int64(m.CreatedUserID),
		Name:          m.Name,
	}
	if m.EstablishedDate != nil {
		result.EstablishedDate = ptr.FromString(ptr.ToTime(m.EstablishedDate).Local().String())
	}

	return result
}

func ToCompanies(m []*table_model.Company) []*Company {
	res := make([]*Company, len(m))
	for i, v := range m {
		res[i] = ToCompany(v)
	}
	return res
}
