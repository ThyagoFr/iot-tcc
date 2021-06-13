package user

import (
	"errors"
	"fmt"
	"github.com/thyagofr/tcc/api/domain/entity"
	domainErrors "github.com/thyagofr/tcc/api/domain/errors"
	"github.com/thyagofr/tcc/api/domain/repository"
	"github.com/thyagofr/tcc/api/domain/service"
)

type User struct {
	repository repository.UserRepository
	crypt      service.Crypt
}

func (u *User) Register(user *entity.User) error {
	_, err := u.repository.FindByEmail(user.Email)
	if err == nil {
		return domainErrors.ErrEmailAlreadyUsed
	}
	if errors.Is(err, domainErrors.ErrInternal) {
		return err
	}
	hash, err := u.crypt.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	err = u.repository.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Login(email, password string) (*entity.User, error) {
	user, err := u.repository.FindByEmail( email)
	if err != nil {
		return nil, err
	}
	ok, err := u.crypt.Check(user.Password, password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("password doesn't match")
	}
	return user, nil
}
