package table_model

type Product struct {
	BaseModel
	Name        string
	Code        string
	Description string
	CompanyID   int
	Company     Company
}

func (t Product) TableName() string {
	return "product"
}
