package repositories

import (
	"context"

	"github.com/thyagofr/tcc/api/domain/entity"
	"github.com/thyagofr/tcc/api/domain/repository"
	"gorm.io/gorm"
)

type UserPostgresql struct {
	db *gorm.DB
}

func NewUserPostgresql(db *gorm.DB) repository.UserRepository {
	return &UserPostgresql{
		db: db,
	}
}

func (u *UserPostgresql) Insert(ctx context.Context, user *entity.User) error {
	return nil
}

func (u *UserPostgresql) FindByID(ctx context.Context, userID string) (*entity.User, error) {
	return nil, nil
}

func (u *UserPostgresql) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

func (u *UserPostgresql) FindAll(ctx context.Context) ([]*entity.User, error) {
	return nil, nil
}

func (u *UserPostgresql) Remove(ctx context.Context, userID string) error {
	return nil
}

func (u *UserPostgresql) Update(ctx context.Context, user *entity.User) error {
	return nil
}
