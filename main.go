package main

import (
	"github.com/nunonano/hacktiv-final-project/app"

	oController "github.com/nunonano/hacktiv-final-project/controller/order"
	oRepository "github.com/nunonano/hacktiv-final-project/repository/order"
	oService "github.com/nunonano/hacktiv-final-project/service/order"

	tController "github.com/nunonano/hacktiv-final-project/controller/todo"
	tRepository "github.com/nunonano/hacktiv-final-project/repository/todo"
	tService "github.com/nunonano/hacktiv-final-project/service/todo"
)

func main() {
	db := app.InitDB()

	orderRepository := oRepository.InitCategoryRepository()
	orderService := oService.InitOrderService(orderRepository, db)
	orderController := oController.InitOrderController(orderService)

	todoRepository := tRepository.InitCategoryRepository()
	todoService := tService.InitTodoService(todoRepository, db)
	todoController := tController.InitTodoController(todoService)

	r := app.InitRouter(orderController, todoController)

	r.Run(":8080")
}
