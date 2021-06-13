package errors

import "github.com/pkg/errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrEmailAlreadyUsed = errors.New("email already used by other user")
	ErrInternal         = errors.New("an error occurred")
)
