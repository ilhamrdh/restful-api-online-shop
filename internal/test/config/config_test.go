package config_test

import (
	"log"
	"online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		filename := "../../../cmd/api/config.yaml"
		err := config.LoadConfig(filename)
		require.Nil(t, err)
		log.Printf("%+v\n", config.Cfg)
	})
}
