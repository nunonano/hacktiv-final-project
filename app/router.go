package app

import (
	oController "github.com/nunonano/hacktiv-final-project/controller/order"
	tController "github.com/nunonano/hacktiv-final-project/controller/todo"

	"github.com/gin-gonic/gin"
)

func InitRouter(orderController oController.OrderController, todoController tController.TodoController) *gin.Engine {
	r := gin.Default()

	orderRouter := r.Group("/order")
	{
		orderRouter.GET("/", orderController.GetAll)
		orderRouter.GET("/:id", orderController.GetById)
		orderRouter.POST("/", orderController.Create)
	}

	todoRouter := r.Group("/todo")
	{
		todoRouter.GET("/", todoController.GetAll)
		todoRouter.GET("/:id", todoController.GetById)
		todoRouter.POST("/", todoController.Create)
		todoRouter.PUT("/:id", todoController.Update)
		todoRouter.DELETE("/:id", todoController.Delete)
	}

	return r
}
