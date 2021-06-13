package user

import "github.com/thyagofr/tcc/api/domain/entity"

type Service interface {
	Register(user *entity.User) error
	Login(email, password string) (user *entity.User,err error)
}



