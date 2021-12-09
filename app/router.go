package app

import (
	"net/http"

	oController "github.com/nunonano/hacktiv-final-project/controller/order"
	tController "github.com/nunonano/hacktiv-final-project/controller/todo"
	"github.com/nunonano/hacktiv-final-project/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func InitRouter(orderController oController.OrderController, todoController tController.TodoController) *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(g *gin.Context) { g.JSON(http.StatusOK, "pong") })

		orderRouter := v1.Group("/order")
		{
			orderRouter.GET("/", orderController.GetAll)
			orderRouter.GET("/:id", orderController.GetById)
			orderRouter.POST("/", orderController.Create)
		}

		todoRouter := v1.Group("/todo")
		{
			todoRouter.GET("/", todoController.GetAll)
			todoRouter.GET("/:id", todoController.GetById)
			todoRouter.POST("/", todoController.Create)
			todoRouter.PUT("/:id", todoController.Update)
			todoRouter.DELETE("/:id", todoController.Delete)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
