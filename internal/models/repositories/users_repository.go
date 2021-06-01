package repositories

import "github.com/yuki-inoue-eng/authenticator/internal/models"

type UserRepository interface {
	Register(models.User) error
	User(loginID string) (*models.User, error)
}