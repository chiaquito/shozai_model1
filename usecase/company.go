package usecase

import (
	"github.com/labstack/echo/v4"

	"shozai_model1/domain/table_model"
	"shozai_model1/repository"
	"shozai_model1/usecase/request"
	"shozai_model1/usecase/response"
)

type Company interface {
	GetCompanyByID(c echo.Context) (*response.Company, error)
	ListCompany(c echo.Context) ([]*response.Company, error)
}

type company struct {
	CompanyRepo repository.Company
}

func NewCompany(companyRepo repository.Company) *company {
	return &company{
		CompanyRepo: companyRepo,
	}
}

func (u *company) GetCompanyByID(c echo.Context) (*response.Company, error) {

	var req request.GetCompanyByID
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	cond := req.ToCompany()

	company, err := u.CompanyRepo.Get(cond)
	if err != nil {
		return nil, err
	}

	return response.ToCompany(company), nil
}

func (u *company) ListCompany(c echo.Context) ([]*response.Company, error) {

	cond := &table_model.Company{}
	comapnies, err := u.CompanyRepo.List(cond)
	if err != nil {
		return nil, err
	}
	return response.ToCompanies(comapnies), nil
}
