package main

import (
	"os"

	"github.com/nunonano/hacktiv-final-project/app"

	oController "github.com/nunonano/hacktiv-final-project/controller/order"
	oRepository "github.com/nunonano/hacktiv-final-project/repository/order"
	oService "github.com/nunonano/hacktiv-final-project/service/order"

	tController "github.com/nunonano/hacktiv-final-project/controller/todo"
	tRepository "github.com/nunonano/hacktiv-final-project/repository/todo"
	tService "github.com/nunonano/hacktiv-final-project/service/todo"
)

// @title Hacktiv8 golang final project API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
// @schemes http
func main() {
	db := app.InitDB()

	orderRepository := oRepository.InitCategoryRepository()
	orderService := oService.InitOrderService(orderRepository, db)
	orderController := oController.InitOrderController(orderService)

	todoRepository := tRepository.InitCategoryRepository()
	todoService := tService.InitTodoService(todoRepository, db)
	todoController := tController.InitTodoController(todoService)

	r := app.InitRouter(orderController, todoController)

	PORT := os.Getenv("PORT")
	r.Run(":" + PORT)
}
