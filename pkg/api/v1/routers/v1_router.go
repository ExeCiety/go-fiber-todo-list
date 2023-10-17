package routers

import (
	apiV1AuthRouters "github.com/ExeCiety/go-fiber-todo-list/pkg/api/v1/auth/routers"
	apiV1TodoListRouters "github.com/ExeCiety/go-fiber-todo-list/pkg/api/v1/todo_list/routers"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router) {
	v1Router := router.Group("/v1")

	apiV1AuthRouters.SetRouter(v1Router)
	apiV1TodoListRouters.SetRouter(v1Router)
}
