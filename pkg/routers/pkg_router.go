package routers

import (
	apiRouters "github.com/ExeCiety/go-fiber-todo-list/pkg/api/routers"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(app *fiber.App) {
	apiRouters.SetRouter(app)
}
