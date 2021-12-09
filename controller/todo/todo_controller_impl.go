package controller

import (
	"net/http"
	"strconv"

	"github.com/nunonano/hacktiv-final-project/model/api"
	tApi "github.com/nunonano/hacktiv-final-project/model/api/todo"
	service "github.com/nunonano/hacktiv-final-project/service/todo"

	"github.com/gin-gonic/gin"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func InitTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

func (controller *TodoControllerImpl) GetById(c *gin.Context) {
	todoId := c.Param("id")
	id, err := strconv.Atoi(todoId)

	var apiResponse api.ApiResponse
	if err != nil {
		apiResponse = api.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		}
	} else {
		todo := controller.TodoService.GetById(c, id)
		apiResponse = api.ApiResponse{
			Code:    http.StatusOK,
			Message: "OK",
			Data:    todo,
		}
	}

	c.JSON(apiResponse.Code, apiResponse)

}

func (controller *TodoControllerImpl) GetAll(c *gin.Context) {
	todos := controller.TodoService.GetAll(c)

	apiResponse := api.ApiResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    todos,
	}

	c.JSON(apiResponse.Code, apiResponse)
}

func (controller *TodoControllerImpl) Create(c *gin.Context) {
	var TodoRequest tApi.TodoRequest

	c.ShouldBindJSON(&TodoRequest)

	todoResponse := controller.TodoService.Create(c, TodoRequest)

	apiResponse := api.ApiResponse{
		Code:    http.StatusCreated,
		Message: "Created",
		Data:    todoResponse,
	}

	c.JSON(apiResponse.Code, apiResponse)
}

func (controller *TodoControllerImpl) Update(c *gin.Context) {
	var TodoRequest tApi.TodoRequest
	var apiResponse api.ApiResponse

	todoId := c.Param("id")

	id, err := strconv.Atoi(todoId)
	if err != nil {
		apiResponse = api.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		}
		c.JSON(apiResponse.Code, apiResponse)
		return
	}

	c.ShouldBindJSON(&TodoRequest)

	todoResponse := controller.TodoService.Update(c, id, TodoRequest)

	apiResponse = api.ApiResponse{
		Code:    http.StatusAccepted,
		Message: "Updated",
		Data:    todoResponse,
	}

	c.JSON(apiResponse.Code, apiResponse)
}

func (controller *TodoControllerImpl) Delete(c *gin.Context) {
	todoId := c.Param("id")
	id, err := strconv.Atoi(todoId)

	var apiResponse api.ApiResponse
	if err != nil {
		apiResponse = api.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		}
	} else {
		todo := controller.TodoService.Delete(c, id)
		apiResponse = api.ApiResponse{
			Code:    http.StatusAccepted,
			Message: "Deleted",
			Data:    todo,
		}
	}

	c.JSON(apiResponse.Code, apiResponse)
}
