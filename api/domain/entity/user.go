package entity

import (
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("password must have unless 8 characters")
	ErrEmptyPassword   = errors.New("password can't be empty")
	ErrInvalidName     = errors.New("name must have unless 2 characters")
	ErrEmptyName       = errors.New("name can't be empty")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrEmptyEmail      = errors.New("email can't me empty")
	ErrCryptPassword   = errors.New("can't crypt password")
	ErrWrongPassword   = errors.New("can't match password")
)

type User struct {
	Entity
	Email    string
	Name     string
	Password string
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrCryptPassword
	}
	u.Password = string(hash)
	return nil
}

func (u *User) VerifyPassword(password string) error {
	pass := []byte(password)
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), pass); err != nil {
		return ErrWrongPassword
	}
	return nil
}

func (u *User) Validate() error {
	err := u.validateName()
	if err != nil {
		return err
	}
	err = u.validateEmail()
	if err != nil {
		return err
	}
	err = u.validatePassword()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) validatePassword() error {
	if len(u.Password) < 8 {
		if len(u.Password) == 0 {
			return ErrEmptyPassword
		}
		return ErrInvalidPassword
	}
	return nil
}

func (u *User) validateName() error {
	if len(u.Name) < 2 {
		if len(u.Name) == 0 {
			return ErrEmptyName
		}
		return ErrInvalidName
	}
	return nil
}

func (u *User) validateEmail() error {
	if len(u.Email) == 0 {
		return ErrEmptyEmail
	}
	if !strings.Contains(u.Email, "@") {
		return ErrInvalidEmail
	}
	return nil
}
