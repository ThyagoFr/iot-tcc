package repositories

import (
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

func (u *UserPostgresql) Insert(user *entity.User) error {
	return nil
}

func (u *UserPostgresql) FindByID(userID string) (*entity.User, error) {
	return nil, nil
}

func (u *UserPostgresql) FindByEmail(email string) (*entity.User, error) {
	return nil, nil
}

func (u *UserPostgresql) FindAll() ([]*entity.User, error) {
	return nil, nil
}

func (u *UserPostgresql) Remove(userID string) error {
	return nil
}

func (u *UserPostgresql) Update(user *entity.User) error {
	return nil
}
