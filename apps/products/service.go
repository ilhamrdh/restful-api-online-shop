package products

import (
	"context"
	"online-shop/infra/response"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	CreateProduct(ctx context.Context, model Product) error
	GetAllProducts(ctx context.Context, model ProductPagination) ([]Product, error)
	GetProductDetail(ctx context.Context, sku string) (Product, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) CreateProduct(ctx context.Context, req CreateProductRequest) error {
	product := Product{
		SKU:       uuid.NewString(),
		Name:      req.Name,
		Stock:     req.Stock,
		Price:     req.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := product.Validate(); err != nil {
		return err
	}
	if err := s.repository.CreateProduct(ctx, product); err != nil {
		return err
	}
	return nil
}

func (s Service) GetAllProducts(ctx context.Context, req ListProductRequest) ([]Product, error) {
	req = req.GenerateDefaultValue()
	pagination := ProductPagination{
		Cursor: req.Cursor,
		Size:   req.Size,
	}
	products, err := s.repository.GetAllProducts(ctx, pagination)
	if err != nil {
		if err == response.ErrNotFound {
			return []Product{}, nil
		}
		return nil, err
	}
	if len(products) == 0 {
		return []Product{}, nil
	}

	return products, nil
}

func (s Service) ProductDetail(ctx context.Context, sku string) (Product, error) {
	var product Product
	product, err := s.repository.GetProductDetail(ctx, sku)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}
