package grpc

import (
	"context"
	itempb "monolith-divided-to-microservices/app/sdk/proto/item/v1"
	"monolith-divided-to-microservices/app/services/item/internal/service"
)

type ItemHandler struct {
	itempb.UnimplementedItemServiceServer
	svc *service.ItemService
}

func NewItemHandler(svc *service.ItemService) *ItemHandler {
	return &ItemHandler{svc: svc}
}

func (h *ItemHandler) GetItem(ctx context.Context, req *itempb.GetItemRequest) (*itempb.GetItemResponse, error) {
	item, err := h.svc.GetItemByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &itempb.GetItemResponse{
		Item: &itempb.Item{
			Id:          req.Id,
			Name:        item.Name,
			Price:       item.Price,
			Stock:       int32(item.Stock),
			Description: item.Description,
		},
	}, nil
}
