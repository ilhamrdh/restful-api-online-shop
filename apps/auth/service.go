package auth

import (
	"context"
	"online-shop/infra/response"
	"online-shop/internal/config"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	CreateAuth(ctx context.Context, model AuthEntity) error
	GetAuthByEmail(ctx context.Context, email string) (AuthEntity, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) Register(ctx context.Context, req RegisterRequest) error {
	authEntity := AuthEntity{
		PublicId:  uuid.New(),
		Email:     req.Email,
		Password:  req.Password,
		Role:      ROLE_User,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := authEntity.Validate(); err != nil {
		return err
	}
	if err := authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
		return err
	}

	model, err := s.repository.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		if err != response.ErrNotFound {
			return err
		}
	}

	if model.IsExists() {
		return response.ErrEmailAlreadyUsed
	}
	return s.repository.CreateAuth(ctx, authEntity)
}

func (s Service) Login(ctx context.Context, req LoginRequest) (string, error) {
	authEntity := AuthEntity{
		Email:    req.Email,
		Password: req.Password,
	}
	if err := authEntity.ValidateEmail(); err != nil {
		return "", err
	}
	if err := authEntity.ValidatePassword(); err != nil {
		return "", err
	}

	model, err := s.repository.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		return "", err
	}

	if err := authEntity.VerifyPassword(model.Password); err != nil {
		return "", response.ErrPasswordNotMatch
	}
	token, err := model.GenerateToken(config.Cfg.App.Encryption.JWTSecret)
	if err != nil {
		return "", nil
	}
	return token, nil
}
