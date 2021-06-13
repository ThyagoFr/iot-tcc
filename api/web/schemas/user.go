package schemas

import (
	"github.com/thyagofr/tcc/api/domain/entity"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (create *CreateUserRequest) ToModel() *entity.User {
	return &entity.User{
		Name: create.Name,
		Email: create.Email,
		Password: create.Password,
	}
}

type UserResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

func ToUserResponse(model *entity.User) *UserResponse {
	return &UserResponse{
		ID: model.ID,
		Email: model.Email,
		CreatedAt: model.CreatedAt,
		Name: model.Name,
	}
}

type UpdateUserRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
