package service

import (
	"context"
	"errors"
	"log"
	"monolith-divided-to-microservices/app/services/user/internal/config"
	"monolith-divided-to-microservices/app/services/user/internal/model"
	"monolith-divided-to-microservices/app/services/user/internal/repository/postgres"
	"monolith-divided-to-microservices/app/services/user/internal/schema"

	"github.com/google/uuid"
)

type UserService struct {
	repo *postgres.UserRepository
	cfg  *config.Config
}

func NewUserService(repo *postgres.UserRepository,
	cfg *config.Config) *UserService {
	return &UserService{
		repo: repo,
		cfg:  cfg,
	}
}

func (u *UserService) GetAllUsers(ctx context.Context) (
	[]schema.UserResponse, error) {
	users, err := u.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var response []schema.UserResponse
	for _, user := range users {
		response = append(response, schema.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return response, nil
}

func (u *UserService) CreateUser(ctx context.Context, req schema.CreateUserRequest) error {
	user := model.User{
		ID:       uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	return u.repo.CreateUser(ctx, &user)
}

func (u *UserService) GetUser(ctx context.Context, filter schema.UserFilters) (
	[]schema.UserResponse, error) {
	users, err := u.repo.GetUser(ctx, filter)
	if err != nil {
		return nil, err
	}

	var response []schema.UserResponse
	for _, user := range users {
		response = append(response, schema.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	return response, nil
}

func (u *UserService) GetByEmailAndPassword(ctx context.Context, email, password string) (*model.User, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		log.Printf("GetByEmail error: %v", err)
		return &model.User{}, err
	}

	log.Printf("DB user found: email='%s', password='%s'", user.Email, user.Password)

	if user.Password != password {
		log.Printf("Password mismatch: db='%s' input='%s'", user.Password, password)
		return &model.User{}, errors.New("invalid password")
	}

	return user, nil
}

func (u *UserService) GetUserByID(ctx context.Context, id string) (*schema.UserResponse, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return &schema.UserResponse{}, err
	}
	return &schema.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
