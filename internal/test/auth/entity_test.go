package auth_test

import (
	"log"
	"online-shop/apps/auth"
	"online-shop/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestEntityValidate(t *testing.T) {
	t.Run("Success  Email", func(t *testing.T) {
		authEntity := auth.AuthEntity{
			Email:    "ilham@example.ir",
			Password: "1234567",
		}
		err := authEntity.Validate()
		require.Nil(t, err)
	})

	t.Run("Email is required", func(t *testing.T) {
		authEntity := auth.AuthEntity{
			Email:    "",
			Password: "1234567",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailRequired, err)
	})

	t.Run("Email is invalid", func(t *testing.T) {
		authEntity := auth.AuthEntity{
			Email:    "ilham.ir",
			Password: "1234567",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailInvalid, err)
	})

	t.Run("Success password", func(t *testing.T) {
		authEntity := auth.AuthEntity{
			Email:    "ilham@example.ir",
			Password: "1234567",
		}
		err := authEntity.Validate()
		require.Nil(t, err)
	})

	t.Run("Password is required", func(t *testing.T) {
		authEntity := auth.AuthEntity{
			Email:    "ilham@example.ir",
			Password: "",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})

	t.Run("Password must have minimum 6 character", func(t *testing.T) {
		authEntity := auth.AuthEntity{
			Email:    "ilham@example.ir",
			Password: "12345",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordInvalidLength, err)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		authEntity := auth.AuthEntity{
			Email:    "ilham@gmail.com",
			Password: "Ilham1234",
		}
		err := authEntity.EncryptPassword(bcrypt.DefaultCost)
		require.Nil(t, err)
		log.Printf("%+v\n", authEntity)
	})
}
