package products

import (
	"online-shop/infra/response"
	"time"
)

type Product struct {
	Id        int       `db:"id"`
	SKU       string    `db:"sku"`
	Name      string    `db:"name"`
	Stock     int16     `db:"stock"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ProductPagination struct {
	Cursor int `json:"cursor"`
	Size   int `json:"size"`
}

func (p Product) Validate() error {
	if err := p.ValidateName(); err != nil {
		return err
	}
	if err := p.ValidateStock(); err != nil {
		return err
	}
	if err := p.ValidatePrice(); err != nil {
		return err
	}
	return nil
}

func (p Product) ValidateName() error {
	if p.Name == "" {
		return response.ErrProductRequired
	}
	if len(p.Name) < 4 {
		return response.ErrProductInvalid
	}
	return nil
}

func (p Product) ValidatePrice() error {
	if p.Price <= 0 {
		return response.ErrPriceInvalid
	}
	return nil
}
func (p Product) ValidateStock() error {
	if p.Stock <= 0 {
		return response.ErrStockInvalid
	}
	return nil
}
