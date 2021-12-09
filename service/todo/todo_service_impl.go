package service

import (
	"context"

	api "github.com/nunonano/hacktiv-final-project/model/api/todo"
	"github.com/nunonano/hacktiv-final-project/model/entity"
	repository "github.com/nunonano/hacktiv-final-project/repository/todo"

	"gorm.io/gorm"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	DB             *gorm.DB
}

func InitTodoService(todoRepository repository.TodoRepository, DB *gorm.DB) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		DB:             DB,
	}
}

func (service *TodoServiceImpl) Delete(ctx context.Context, todoId int) api.TodoResponse {
	todoId, err := service.TodoRepository.Delete(ctx, service.DB, todoId)
	if err != nil {
		return api.TodoResponse{}
	} else {
		return api.TodoResponse{
			Id: todoId,
		}
	}
}

func (service *TodoServiceImpl) Update(ctx context.Context, todoId int, request api.TodoRequest) api.TodoResponse {

	todo := entity.Todo{
		Name: request.Name,
	}

	todo = service.TodoRepository.Update(ctx, service.DB, todoId, todo)

	return api.TodoResponse{
		Name: todo.Name,
		Id:   todo.Id,
	}
}

func (service *TodoServiceImpl) Create(ctx context.Context, request api.TodoRequest) api.TodoResponse {

	todo := entity.Todo{
		Name: request.Name,
	}

	todo = service.TodoRepository.Create(ctx, service.DB, todo)

	return api.TodoResponse{
		Name: todo.Name,
		Id:   todo.Id,
	}
}

func (service *TodoServiceImpl) GetAll(ctx context.Context) []api.TodoResponse {

	var todoResponses []api.TodoResponse

	todos := service.TodoRepository.GetAll(ctx, service.DB)

	for _, todo := range todos {
		todoResponses = append(todoResponses, api.TodoResponse{Name: todo.Name,
			Id: todo.Id})
	}

	return todoResponses
}

func (service *TodoServiceImpl) GetById(ctx context.Context, todoId int) api.TodoResponse {
	todo, err := service.TodoRepository.GetById(ctx, service.DB, todoId)
	if err != nil {
		return api.TodoResponse{}
	} else {
		return api.TodoResponse{
			Name: todo.Name,
			Id:   todo.Id,
		}
	}
}
