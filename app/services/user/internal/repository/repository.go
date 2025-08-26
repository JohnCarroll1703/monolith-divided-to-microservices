package repository

import (
	"context"
	"monolith-divided-to-microservices/app/services/user/internal/repository/postgres"
	"monolith-divided-to-microservices/app/services/user/internal/schema"
	"monolith-divided-to-microservices/app/services/user/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	UserRepo *postgres.UserRepository
}

type User interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, u *model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*schema.UserResponse, error)
}


func NewRepository(db *pgxpool.Pool) *Repositories {
	userRepo := postgres.NewUser(db)
	return &Repositories{
		UserRepo: userRepo,
	}
}
