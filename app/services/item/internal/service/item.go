package service

import (
	"context"
	"monolith-divided-to-microservices/app/services/item/internal/config"
	"monolith-divided-to-microservices/app/services/item/internal/repository/postgres"
	"monolith-divided-to-microservices/app/services/item/internal/schema"
)

type ItemService struct {
	repo *postgres.ItemRepository
	cfg  *config.Config
}

func NewItemService(repo *postgres.ItemRepository,
	cfg *config.Config) *ItemService {
	return &ItemService{
		repo: repo,
		cfg:  cfg,
	}
}

func (i *ItemService) GetAllItems(ctx context.Context) (
	[]schema.ItemResponse, error) {
	res, err := i.repo.GetItems(ctx)
	if err != nil {
		return nil, err
	}

	var response []schema.ItemResponse
	for _, item := range res {
		response = append(response, schema.ItemResponse{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Stock:       item.Stock,
			CategoryID:  item.CategoryID,
			CountryID:   item.CountryID,
			AddedAt:     item.AddedAt,
		})
	}
	return response, nil
}

func (i *ItemService) GetItemByID(ctx context.Context, id string) (*schema.ItemResponse, error) {
	item, err := i.repo.GetItemByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil
	}

	response := &schema.ItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Stock:       item.Stock,
		CategoryID:  item.CategoryID,
		CountryID:   item.CountryID,
		AddedAt:     item.AddedAt,
	}
	return response, nil
}
