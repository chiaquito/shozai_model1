package table_model

import "time"

type Company struct {
	BaseModel
	Name            string
	EstablishedDate *time.Time
}

func (t Company) TableName() string {
	return "company"
}
