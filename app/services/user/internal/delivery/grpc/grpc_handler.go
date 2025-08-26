package grpc

import (
	"context"
	userpb "monolith-divided-to-microservices/app/sdk/proto/user/v1"
	"monolith-divided-to-microservices/app/services/user/internal/service"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	u, err := h.svc.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    req.Id,
			Name:  u.Name,
			Email: u.Email,
		},
	}, nil
}
