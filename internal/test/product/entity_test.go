package product_test

import (
	"online-shop/apps/products"
	"online-shop/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateProduct(t *testing.T) {
	t.Run("Product Success", func(t *testing.T) {
		product := products.Product{
			Name:  "Baju Baru",
			Stock: 10,
			Price: 10.000,
		}
		err := product.Validate()
		require.Nil(t, err)
	})
	t.Run("Product Require", func(t *testing.T) {
		product := products.Product{
			Name:  "",
			Stock: 10,
			Price: 10.000,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})
	t.Run("Product Invalid", func(t *testing.T) {
		product := products.Product{
			Name:  "ba",
			Stock: 10,
			Price: 10.000,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductInvalid, err)
	})
	t.Run("Stock Invalid", func(t *testing.T) {
		product := products.Product{
			Name:  "Baju Baru",
			Stock: 0,
			Price: 10.000,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrStockInvalid, err)
	})
	t.Run("Price Invalid", func(t *testing.T) {
		product := products.Product{
			Name:  "Baju Baru",
			Stock: 10,
			Price: 0,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})
}
