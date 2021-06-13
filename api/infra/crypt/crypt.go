package crypt

import "golang.org/x/crypto/bcrypt"

type Crypt struct{}

// Hash - HashPassword
func (c *Crypt) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Check - CheckPasswordHash
func (c *Crypt) Check(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
