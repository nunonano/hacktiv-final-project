package service

import (
	"context"

	api "github.com/nunonano/hacktiv-final-project/model/api/todo"
)

type TodoService interface {
	GetAll(ctx context.Context) []api.TodoResponse
	GetById(ctx context.Context, id int) api.TodoResponse
	Create(ctx context.Context, request api.TodoRequest) api.TodoResponse
	Update(ctx context.Context, id int, request api.TodoRequest) api.TodoResponse
	Delete(ctx context.Context, id int) api.TodoResponse
}
