package integrated_model

type Product struct {
	ProductID              int    `gorm:"column:product_id"`
	ProductName            string `gorm:"column:product_name"`
	ProductCode            string `gorm:"column:product_code"`
	ProductDescription     string `gorm:"column:product_description"`
	CompanyID              int    `gorm:"column:product_company_id"`
	CompanyName            string `gorm:"column:company_name"`
	CompanyEstablishedDate string `gorm:"column:company_established_date"`
}

func (t Product) TableName() string {
	return "product"
}
