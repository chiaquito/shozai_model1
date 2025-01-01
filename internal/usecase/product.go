package usecase

import (
	echo "github.com/labstack/echo/v4"

	"shozai_model1/internal/repository"
	request "shozai_model1/internal/usecase/request/product"
	"shozai_model1/internal/usecase/response"
)

type Product interface {
	GetProductByID(c echo.Context) (*response.Product, error)
	ListProduct(c echo.Context) ([]*response.Product, error)
	CreateProduct(c echo.Context) (*response.Product, error)
	BulkCreateProduct(c echo.Context) ([]*response.Product, error)
	UpdateProduct(c echo.Context) (*response.Product, error)
	ListProductRaw(c echo.Context) ([]*response.Product, error)
}

type product struct {
	ProductRepo repository.Product
}

func NewProduct(productRepo repository.Product) *product {
	return &product{
		ProductRepo: productRepo,
	}
}

func (u *product) GetProductByID(c echo.Context) (*response.Product, error) {

	var req request.GetProductByID
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	cond := req.ToProduct()

	product, err := u.ProductRepo.Get(cond)
	if err != nil {
		return nil, err
	}

	return response.ToProduct(product), nil
}

func (u *product) ListProduct(c echo.Context) ([]*response.Product, error) {

	var req request.ListProduct
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	cond := req.ToProduct()

	products, err := u.ProductRepo.List(cond)
	if err != nil {
		return nil, err
	}
	return response.ToProducts(products), nil
}

func (u *product) CreateProduct(c echo.Context) (*response.Product, error) {

	var req request.CreateProduct
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	m := req.ToProduct()

	product, err := u.ProductRepo.Create(m)
	if err != nil {
		return nil, err
	}
	return response.ToProduct(product), nil
}

func (u *product) BulkCreateProduct(c echo.Context) ([]*response.Product, error) {

	var req request.BulkCreateProduct
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	m, err := req.ToProducts()
	if err != nil {
		return nil, err
	}
	products, err := u.ProductRepo.BulkCreate(m)
	if err != nil {
		return nil, err
	}
	return response.ToProducts(products), nil
}

func (u *product) UpdateProduct(c echo.Context) (*response.Product, error) {

	var req request.UpdateProduct
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	m, err := req.ToProduct()
	if err != nil {
		return nil, err
	}

	product, err := u.ProductRepo.Update(m)
	if err != nil {
		return nil, err
	}
	return response.ToProduct(product), nil
}

func (u *product) ListProductRaw(c echo.Context) ([]*response.Product, error) {

	var req request.ListProductRaw
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	if err := c.Validate(req); err != nil {
		return nil, err
	}

	cond := req.ToProduct()

	products, err := u.ProductRepo.ListRaw(cond)
	if err != nil {
		return nil, err
	}
	return response.ToProductsRaw(products), nil
}
