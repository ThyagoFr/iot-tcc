package repository

import (
	"context"
	"github.com/thyagofr/tcc/api/domain/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, userID string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	Remove(ctx context.Context, userID string) error
	Update(ctx context.Context, user *entity.User) error
}
