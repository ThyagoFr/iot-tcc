package interfaces

import (
	"context"

	"github.com/thyago/tcc/api-service/domain/entity"
)

type UserCase interface {
	Register(ctx context.Context, user *entity.User) error 
	Login(ctx context.Context, user *entity.User
}