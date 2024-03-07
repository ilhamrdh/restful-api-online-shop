package auth_test

import (
	"context"
	"fmt"
	"log"
	"online-shop/apps/auth"
	"online-shop/external/database"
	"online-shop/infra/response"
	"online-shop/internal/config"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var svc auth.Service

func init() {
	filename := "../../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectionPostgres(config.Cfg.DBConfig)
	if err != nil {
		panic(err)
	}

	repository := auth.NewRepository(db)
	svc = auth.NewService(repository)
}

func TestRegisterSuccess(t *testing.T) {
	req := auth.RegisterRequest{
		Email:    fmt.Sprintf("%v@example.ir", uuid.NewString()),
		Password: "ilham12345",
	}
	err := svc.Register(context.Background(), req)
	require.Nil(t, err)
}
func TestRegisterFail(t *testing.T) {
	email := fmt.Sprintf("%v@example.ir", uuid.NewString())
	req := auth.RegisterRequest{
		Email:    email,
		Password: "ilham12345",
	}
	err := svc.Register(context.Background(), req)
	require.Nil(t, err)

	err = svc.Register(context.Background(), req)
	require.NotNil(t, err)
	require.Equal(t, response.ErrEmailAlreadyUsed, err)
}

func TestLogin(t *testing.T) {
	email := fmt.Sprintf("%v@example.ir", uuid.NewString())
	pass := "Ilham12345"
	reqRegister := auth.RegisterRequest{
		Email:    email,
		Password: pass,
	}
	err := svc.Register(context.Background(), reqRegister)
	require.Nil(t, err)

	reqLogin := auth.LoginRequest{
		Email:    email,
		Password: pass,
	}

	token, err := svc.Login(context.Background(), reqLogin)
	require.Nil(t, err)
	require.NotEmpty(t, token)
	log.Println(token)
}
