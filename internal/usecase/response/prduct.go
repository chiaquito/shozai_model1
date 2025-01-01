package response

import (
	"shozai_model1/internal/domain/integrated_model"
	"shozai_model1/internal/domain/table_model"
	"shozai_model1/internal/pkg/ptr"
)

type Product struct {
	ID            int64   `json:"id"`
	Version       int64   `json:"version"`
	Name          string  `json:"name"`
	Code          string  `json:"code"`
	Description   string  `json:"description"`
	CreatedUserID int64   `json:"createdUserId"`
	CompanyID     int64   `json:"companyId"`
	Company       Company `json:"company"`
}

func ToProduct(m *table_model.Product) *Product {
	result := &Product{
		ID:            int64(m.ID),
		Version:       int64(m.Version),
		Name:          m.Name,
		Code:          m.Code,
		Description:   m.Description,
		CreatedUserID: int64(m.CreatedUserID),
		CompanyID:     int64(m.CompanyID),
		Company:       *ToCompany(&m.Company),
	}

	return result
}

func ToProducts(m []*table_model.Product) []*Product {
	res := make([]*Product, len(m))
	for i, v := range m {
		res[i] = ToProduct(v)
	}
	return res
}

func ToProductRaw(m *integrated_model.Product) *Product {
	result := &Product{
		ID:          int64(m.ProductID),
		Name:        m.ProductName,
		Code:        m.ProductCode,
		Description: m.ProductDescription,
		CompanyID:   int64(m.CompanyID),
	}
	result.Company.ID = int64(m.CompanyID)
	result.Company.Name = m.CompanyName
	result.Company.EstablishedDate = ptr.FromString(m.CompanyEstablishedDate)
	return result
}

func ToProductsRaw(m []*integrated_model.Product) []*Product {
	res := make([]*Product, len(m))
	for i, v := range m {
		res[i] = ToProductRaw(v)
	}
	return res

}
