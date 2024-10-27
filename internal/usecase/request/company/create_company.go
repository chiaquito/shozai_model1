package request

import (
	"shozai_model1/internal/domain/table_model"
	"shozai_model1/internal/pkg/time"
)

type CreateCompany struct {
	Name            *string `json:"name" validate:"required"`
	EstablishedDate string  `json:"established_date"`
}

func (r CreateCompany) ToCompany() (*table_model.Company, error) {

	m := &table_model.Company{
		Name: *r.Name,
	}

	if r.EstablishedDate == "" {
		return m, nil
	}

	t, err := time.Str2time(r.EstablishedDate)
	if err != nil {
		return nil, err
	}

	m.EstablishedDate = t
	return m, nil
}
