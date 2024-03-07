package products

import (
	"context"
	"database/sql"
	"online-shop/infra/response"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateProduct(ctx context.Context, model Product) error {
	query := `
		INSERT INTO products(
			sku, name, stock, price, created_at, updated_at
		) VALUES (
			:sku, :name, :stock, :price, :created_at, :updated_at
		)
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.ExecContext(ctx, model); err != nil {
		return err
	}
	return nil
}

func (r repository) GetAllProducts(ctx context.Context, model ProductPagination) ([]Product, error) {
	var products []Product
	query := `
		SELECT id, sku, name, stock, price, created_at, updated_at
		FROM products
		WHERE id > $1
		ORDER BY id ASC
		LIMIT $2
	`
	if err := r.db.SelectContext(ctx, &products, query, model.Cursor, model.Size); err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	return products, nil
}

func (r repository) GetProductDetail(ctx context.Context, sku string) (Product, error) {
	var product Product
	query := `
		SELECT id, sku, name, stock, price, created_at, updated_at
		FROM products
		WHERE sku = $1
	`
	if err := r.db.GetContext(ctx, &product, query, sku); err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return Product{}, err
		}
		return Product{}, err
	}
	return product, nil
}
