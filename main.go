package main

import (
	"log"

	pkgRouters "github.com/ExeCiety/go-fiber-todo-list/pkg/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	pkgRouters.SetRouter(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error when listen to port 3000: %v", err)
		return
	}
}
