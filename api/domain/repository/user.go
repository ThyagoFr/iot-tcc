package repository

import (
	"github.com/thyagofr/tcc/api/domain/entity"
)

type UserRepository interface {
	Insert(user *entity.User) error
	FindByID(userID string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Remove(userID string) error
	Update(user *entity.User) error
}
