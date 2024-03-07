package auth

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

func (r repository) CreateAuth(ctx context.Context, model AuthEntity) error {
	query := `
		INSERT INTO auth (
			public_id, email, password, role, created_at, updated_at
		) VALUES (
			:public_id, :email, :password, :role, :created_at, :updated_at
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) GetAuthByEmail(ctx context.Context, email string) (AuthEntity, error) {
	var model AuthEntity
	query := `
		SELECT
			id, public_id, email, password, role, created_at, updated_at
		FROM auth
		WHERE email=$1
	`
	err := r.db.GetContext(ctx, &model, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return AuthEntity{}, err
		}
		return AuthEntity{}, err
	}
	return model, nil
}
