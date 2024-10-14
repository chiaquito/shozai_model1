package route

import (
	"shozai_model1/internal/handler"

	echo "github.com/labstack/echo/v4"
)

func NewRouter(h *handler.Handler) *echo.Echo {

	e := echo.New()

	api := e.Group("/api")
	{
		api.GET("/health", h.Hello.MyMethod)

		company := api.Group("/companies")
		{
			company.GET("", h.Company.ListCompany)
			company.GET("/:id", h.Company.GetCompanyByID)
			company.POST("", h.Company.CreateCompany)
			company.POST("/bulk", h.Company.BulkCreateCompany)
		}
	}
	return e
}
