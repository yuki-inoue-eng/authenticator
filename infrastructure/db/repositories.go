package db

import (
	"context"
	"fmt"

	"github.com/yuki-inoue-eng/authenticator/configs"
	"github.com/yuki-inoue-eng/authenticator/internal/models/repositories"
)

func NewRepositories(ctx context.Context, configs configs.Configs) (*repositories.Repositories, error) {
	if err := NewDB(ctx, configs); err != nil {
		return nil, fmt.Errorf("failed to establish db connection: %v", err)
	}
	return &repositories.Repositories{
		UserRepository: newUsersRepository(),
	}, nil
}
