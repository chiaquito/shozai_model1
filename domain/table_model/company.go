package table_model

type Company struct {
	BaseModel
	Name string `json:"name"`
}

func (t Company) TableName() string {
	return "company"
}
