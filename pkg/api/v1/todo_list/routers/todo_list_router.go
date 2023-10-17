package routers

import (
	apiV1TodoListControllers "github.com/ExeCiety/go-fiber-todo-list/pkg/api/v1/todo_list/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router) {
	todoListRouter := router.Group("/todo-list")

	todoListRouter.Get("/", apiV1TodoListControllers.Get)
	todoListRouter.Get("/:id", apiV1TodoListControllers.GetByParam)
	todoListRouter.Post("/", apiV1TodoListControllers.Create)
	todoListRouter.Put("/:id", apiV1TodoListControllers.UpdateByParam)
	todoListRouter.Delete("/", apiV1TodoListControllers.BulkDelete)
}
