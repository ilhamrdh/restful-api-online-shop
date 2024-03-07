package database

import (
	"fmt"
	"online-shop/internal/config"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectionPostgres(cfg config.DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.Name,
	)

	db, err := sqlx.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxIdletimeConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPool.MaxLifetimeConnection) * time.Second)
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxOpenConnection))
	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxIdleConnection))

	return db, nil
}
