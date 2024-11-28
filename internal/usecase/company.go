package usecase

import (
	"shozai_model1/internal/pkg/context"
	"shozai_model1/internal/repository"
	request "shozai_model1/internal/usecase/request/company"
	"shozai_model1/internal/usecase/response"
)

type Company interface {
	GetCompanyByID(c *context.Context) (*response.Company, error)
	ListCompany(c *context.Context) ([]*response.Company, error)
	CreateCompany(c *context.Context) (*response.Company, error)
	BulkCreateCompany(c *context.Context) ([]*response.Company, error)
	UpdateCompany(c *context.Context) (*response.Company, error)
}

type company struct {
	CompanyRepo repository.Company
}

func NewCompany(companyRepo repository.Company) *company {
	return &company{
		CompanyRepo: companyRepo,
	}
}

func (u *company) GetCompanyByID(c *context.Context) (*response.Company, error) {

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

func (u *company) ListCompany(c *context.Context) ([]*response.Company, error) {

	var req request.ListCompany
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	cond := req.ToCompany()

	comapnies, err := u.CompanyRepo.List(cond)
	if err != nil {
		return nil, err
	}
	return response.ToCompanies(comapnies), nil
}

func (u *company) CreateCompany(c *context.Context) (*response.Company, error) {

	var req request.CreateCompany
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	m, err := req.ToCompany()
	if err != nil {
		return nil, err
	}

	company, err := u.CompanyRepo.Create(m)
	if err != nil {
		return nil, err
	}
	return response.ToCompany(company), nil
}

func (u *company) BulkCreateCompany(c *context.Context) ([]*response.Company, error) {

	var req request.BulkCreateCompany
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	m, err := req.ToCompanies()
	if err != nil {
		return nil, err
	}
	companies, err := u.CompanyRepo.BulkCreate(m)
	if err != nil {
		return nil, err
	}
	return response.ToCompanies(companies), nil
}

func (u *company) UpdateCompany(c *context.Context) (*response.Company, error) {

	var req request.UpdateCompany
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	m, err := req.ToCompany()
	if err != nil {
		return nil, err
	}

	company, err := u.CompanyRepo.Update(m)
	if err != nil {
		return nil, err
	}
	return response.ToCompany(company), nil
}
