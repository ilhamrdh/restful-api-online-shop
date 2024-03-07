package products

import "time"

type ProductResponse struct {
	Id    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
}

type ProductDetailResponse struct {
	Id        int       `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Stock     int16     `json:"stock"`
	Price     int       `json:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewProductResponse(products []Product) []ProductResponse {
	var productList = []ProductResponse{}
	for _, product := range products {
		productList = append(productList, ProductResponse{
			Id:    product.Id,
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		})
	}
	return productList
}
