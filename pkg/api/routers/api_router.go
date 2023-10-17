package routers

import (
	apiMiddlewares "github.com/ExeCiety/go-fiber-todo-list/pkg/api/middlewares"
	apiV1Routers "github.com/ExeCiety/go-fiber-todo-list/pkg/api/v1/routers"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(app *fiber.App) {
	apiRouter := app.Group("/api", apiMiddlewares.ApiMiddleware)

	apiV1Routers.SetRouter(apiRouter)
}
