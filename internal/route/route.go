package route

import (
	"shozai_model1/internal/handler"

	echo "github.com/labstack/echo/v4"
)

func NewRouter(h *handler.Handler) *echo.Echo {

	e := echo.New()

	api := e.Group("/api")
	{
		api.GET("/health", h.Health.HealthCheck)
		company := api.Group("/companies")
		{
			company.GET("", h.Company.ListCompany)
			company.GET("/:id", h.Company.GetCompanyByID)
			company.POST("", h.Company.CreateCompany)
			company.POST("/bulk", h.Company.BulkCreateCompany)
			company.PUT("/:id", h.Company.UpdateCompany)
		}
		products := api.Group("/products")
		{
			products.GET("", h.Product.ListProduct)
			products.GET("/:id", h.Product.GetProductByID)
			products.POST("", h.Product.CreateProduct)
			products.POST("/bulk", h.Product.BulkCreateProduct)
			products.PUT("/:id", h.Product.UpdateProduct)
			products.GET("/raw", h.Product.ListProductRaw)
		}
	}
	return e
}
