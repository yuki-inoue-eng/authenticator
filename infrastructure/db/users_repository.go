package db

import (
	"errors"

	"github.com/yuki-inoue-eng/authenticator/internal/models"
	"github.com/yuki-inoue-eng/authenticator/internal/models/repositories"
	"gorm.io/gorm"
)

type usersRepository struct {
	ConnPool *MySQL
}

func newUsersRepository() repositories.UserRepository{
	return usersRepository{
		ConnPool: ConPool,
	}
}

func (r usersRepository) Register(user models.User) error {
	result := r.ConnPool.Con.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r usersRepository) User(loginID string) (*models.User, error) {
	u := models.User{}
	result := r.ConnPool.Con.Model(u).Where("login_id = ?", loginID).First(u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &u, nil
}
