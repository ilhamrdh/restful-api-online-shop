package product_test

import (
	"context"
	"log"
	"online-shop/apps/products"
	"online-shop/external/database"
	"online-shop/infra/response"
	"online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc products.Service

func init() {
	filename := "../../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectionPostgres(config.Cfg.DBConfig)
	if err != nil {
		panic(err)
	}

	repository := products.NewRepository(db)
	svc = products.NewService(repository)
}

func TestCreateProductSuccess(t *testing.T) {
	req := products.CreateProductRequest{
		Name:  "Baju Barong",
		Stock: 10,
		Price: 10000,
	}
	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}
func TestCreateProductFail(t *testing.T) {
	req := products.CreateProductRequest{
		Name:  "",
		Stock: 10,
		Price: 10000,
	}
	err := svc.CreateProduct(context.Background(), req)
	require.NotNil(t, err)
	require.Equal(t, response.ErrProductRequired, err)
}

func TestListProductSuccess(t *testing.T) {
	pagination := products.ListProductRequest{
		Cursor: 0,
		Size:   10,
	}
	products, err := svc.GetAllProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Println(products)
}
func TestProductDetailSuccess(t *testing.T) {
	ctx := context.Background()
	req := products.CreateProductRequest{
		Name:  "Baju Pantai",
		Stock: 10,
		Price: 10000,
	}

	err := svc.CreateProduct(ctx, req)
	require.Nil(t, err)

	products, err := svc.GetAllProducts(ctx, products.ListProductRequest{Cursor: 0, Size: 10})
	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)

	product, err := svc.ProductDetail(ctx, products[0].SKU)
	require.Nil(t, err)
	require.NotEmpty(t, product)
	log.Println(product)

}
