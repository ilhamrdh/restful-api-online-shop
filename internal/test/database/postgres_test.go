package database_test

import (
	"online-shop/external/database"
	"online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	filename := "../../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
}

func TestConnectionPostgres(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, err := database.ConnectionPostgres(config.Cfg.DBConfig)

		require.Nil(t, err)
		require.NotNil(t, db)
	})
}
