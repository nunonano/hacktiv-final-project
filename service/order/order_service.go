package service

import (
	"context"

	api "github.com/nunonano/hacktiv-final-project/model/api/order"
)

type OrderService interface {
	GetAll(ctx context.Context) []api.OrderResponse
	GetById(ctx context.Context, id int) api.OrderResponse
	Create(ctx context.Context, request api.OrderCreateRequest) api.OrderResponse
}
