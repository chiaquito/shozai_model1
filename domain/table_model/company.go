package table_model

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (t Company) TableName() string {
	return "company"
}
