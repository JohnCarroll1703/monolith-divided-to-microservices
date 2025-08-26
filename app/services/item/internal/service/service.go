package service

import (
	"context"
	"monolith-divided-to-microservices/app/services/item/internal/schema"
	"monolith-divided-to-microservices/app/services/item/internal/config"
	"monolith-divided-to-microservices/app/services/item/internal/repository"
)

type Services struct {
	ItemService *ItemService
}

func NewServices(repo *repository.Repositories, cfg *config.Config) *Services {
	return &Services{
		ItemService: NewItemService(repo.ItemRepo, cfg),
	}
}

type Item interface {
	GetItems(ctx context.Context) ([]schema.ItemResponse, error)
}
