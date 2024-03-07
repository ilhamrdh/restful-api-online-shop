package utility_test

import (
	"log"
	"online-shop/utility"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		publicId := uuid.NewString()
		tokenSetring, err := utility.GenerateToken(publicId, "user", "IniScret")
		require.Nil(t, err)
		require.NotEmpty(t, tokenSetring)
		log.Println(tokenSetring)
	})
}

func TestVerifyToken(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		publicId := uuid.NewString()
		role := "user"
		tokenString, err := utility.GenerateToken(publicId, role, "IniSecret")
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)

		jwtId, jwtRole, err := utility.ValidateToken(tokenString, "IniSecret")
		require.Nil(t, err)
		require.NotEmpty(t, jwtId)
		require.NotEmpty(t, jwtRole)

		require.Equal(t, publicId, jwtId)
		require.Equal(t, role, jwtRole)
	})
}
