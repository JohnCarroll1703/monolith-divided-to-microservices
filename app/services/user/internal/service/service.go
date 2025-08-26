package service

import (
	"context"
	"monolith-divided-to-microservices/app/services/user/internal/config"
	"monolith-divided-to-microservices/app/services/user/internal/repository"
	"monolith-divided-to-microservices/app/services/user/internal/schema"
)

type Services struct {
	UserService *UserService
}

type User interface {
	GetAllUsers(ctx context.Context) ([]schema.UserResponse, error)
	CreateUser(ctx context.Context, req schema.CreateUserRequest) error
}

func NewServices(repo *repository.Repositories, cfg *config.Config) *Services {
	return &Services{
		UserService: NewUserService(repo.UserRepo, cfg),
	}
}
