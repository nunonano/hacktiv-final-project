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

// @Tags todo
// @Summary Delete data Todo by ID
// @Description This is information about update todo
// @Accept json
// @Produce json
// @Success 200 {array} entity.Todo
// @Param id path int true "Todo ID"
// @Router /todo/{id} [delete]
func (service *TodoServiceImpl) Delete(ctx context.Context, todoId uint) (api.TodoResponse, error) {
	todoId, err := service.TodoRepository.Delete(ctx, service.DB, todoId)

	if err != nil {
		return api.TodoResponse{}, err
	}

	return api.TodoResponse{
		ID: todoId,
	}, nil
}

// @Tags todo
// @Summary Update data Todo by ID
// @Description This is information about update todo
// @Accept json
// @Produce json
// @Success 200 {array} entity.Todo
// @Param id path int true "Todo ID"
// @Param data body entity.Todo true "Update Data Todo"
// @Router /todo/{id} [put]
func (service *TodoServiceImpl) Update(ctx context.Context, todoId uint, request api.TodoRequest) (api.TodoResponse, error) {

	todo := entity.Todo{
		Name: request.Name,
	}

	todo, err := service.TodoRepository.Update(ctx, service.DB, todoId, todo)

	if err != nil {
		return api.TodoResponse{}, err
	}

	return api.TodoResponse{
		Name: todo.Name,
		ID:   todo.ID,
	}, nil
}

// @Tags todo
// @Summary Create Todo
// @Description Create Todo Description
// @Accept json
// @Produce json
// @Success 200 {array} entity.Todo
// @Param id body entity.Todo true "Input Data"
// @Router /todo [post]
func (service *TodoServiceImpl) Create(ctx context.Context, request api.TodoRequest) api.TodoResponse {

	todo := entity.Todo{
		Name: request.Name,
	}

	todo = service.TodoRepository.Create(ctx, service.DB, todo)

	return api.TodoResponse{
		Name: todo.Name,
		ID:   todo.ID,
	}
}

// @Tags todo
// @Summary Get All Todo
// @Description This is information about all todos
// @Accept json
// @Produce json
// @Success 200 {array} entity.Todo
// @Router /todo [get]
func (service *TodoServiceImpl) GetAll(ctx context.Context) []api.TodoResponse {

	var todoResponses []api.TodoResponse

	todos := service.TodoRepository.GetAll(ctx, service.DB)

	for _, todo := range todos {
		todoResponses = append(todoResponses, api.TodoResponse{Name: todo.Name,
			ID: todo.ID})
	}

	return todoResponses
}

// @Tags todo
// @Summary Get Todo by ID
// @Description This is information about todo
// @Accept json
// @Produce json
// @Success 200 {array} entity.Todo
// @Param id path int true "Todo ID"
// @Router /todo/{id} [get]
func (service *TodoServiceImpl) GetById(ctx context.Context, todoId uint) api.TodoResponse {
	todo, err := service.TodoRepository.GetById(ctx, service.DB, todoId)
	if err != nil {
		return api.TodoResponse{}
	} else {
		return api.TodoResponse{
			Name: todo.Name,
			ID:   todo.ID,
		}
	}
}
