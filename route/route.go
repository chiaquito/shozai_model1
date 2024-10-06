package route

import (
	"shozai_model1/handler"

	echo "github.com/labstack/echo/v4"
)

func NewRouter(h *handler.Handler) *echo.Echo {

	e := echo.New()

	api := e.Group("/api/")
	{
		api.GET("", h.Hello.MyMethod)

		company := api.Group("companies/")
		{
			company.GET("", h.Company.ListCompany)
			company.GET(":id", h.Company.GetCompanyByID)
		}
	}
	return e
}
