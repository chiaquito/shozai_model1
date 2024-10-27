package request

import (
	"shozai_model1/internal/domain/table_model"
	"shozai_model1/internal/pkg/time"
)

type UpdateCompany struct {
	ID              *int64  `path:"id" validate:"required"`
	Name            *string `json:"name" validate:"required,max=255"`
	EstablishedDate *string `json:"established_date" validate:"required"`
	Version         *int64  `json:"version" validate:"required,gt=0"`
}

func (r *UpdateCompany) ToCompany() (*table_model.Company, error) {
	m := &table_model.Company{
		BaseModel: table_model.BaseModel{
			ID:      int(*r.ID),
			Version: int(*r.Version),
		},
		Name: *r.Name,
	}
	if r.EstablishedDate == nil {
		return m, nil
	}

	t, err := time.Str2time(*r.EstablishedDate)
	if err != nil {
		return nil, err
	}

	m.EstablishedDate = t

	return m, nil
}
