package auth

import (
	"online-shop/infra/response"
	"online-shop/utility"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	ROLE_Admin Role = "admin"
	ROLE_User  Role = "user"
)

type AuthEntity struct {
	Id        int       `db:"id"`
	PublicId  uuid.UUID `db:"public_id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      Role      `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (a AuthEntity) Validate() error {
	if err := a.ValidateEmail(); err != nil {
		return err
	}
	if err := a.ValidatePassword(); err != nil {
		return err
	}
	return nil
}

func (a AuthEntity) ValidateEmail() error {
	if a.Email == "" {
		return response.ErrEmailRequired
	}
	emails := strings.Split(a.Email, "@")
	if len(emails) != 2 {
		return response.ErrEmailInvalid
	}
	return nil
}

func (a AuthEntity) ValidatePassword() error {
	if a.Password == "" {
		return response.ErrPasswordRequired
	}

	if len(a.Password) < 6 {
		return response.ErrPasswordInvalidLength
	}
	return nil
}

func (a AuthEntity) IsExists() bool {
	return a.Id != 0
}

func (a *AuthEntity) EncryptPassword(salt int) error {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(encryptedPass)
	return nil
}

func (a AuthEntity) VerifyPassword(encrypted string) error {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(a.Password))
}

func (a AuthEntity) GenerateToken(secret string) (string, error) {
	return utility.GenerateToken(a.PublicId.String(), string(a.Role), secret)
}
