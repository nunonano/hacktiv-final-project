package service

import (
	"context"

	api "github.com/nunonano/hacktiv-final-project/model/api/todo"
)

type TodoService interface {
	GetAll(ctx context.Context) []api.TodoResponse
	GetById(ctx context.Context, id uint) api.TodoResponse
	Create(ctx context.Context, request api.TodoRequest) api.TodoResponse
	Update(ctx context.Context, id uint, request api.TodoRequest) (api.TodoResponse, error)
	Delete(ctx context.Context, id uint) (api.TodoResponse, error)
}
