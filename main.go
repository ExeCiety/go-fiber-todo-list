package main

import (
	"fmt"
	"github.com/ExeCiety/go-fiber-todo-list/utils/enums"
	"log"
	"os"

	"github.com/ExeCiety/go-fiber-todo-list/cmd"
	"github.com/ExeCiety/go-fiber-todo-list/db"
	pkgRouters "github.com/ExeCiety/go-fiber-todo-list/pkg/routers"
	"github.com/ExeCiety/go-fiber-todo-list/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Set root path
	cwd, _ := os.Getwd()
	utils.SetRootPath(cwd + "/")

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// Execute cmd tool
	cmd.Execute()

	// Connect DB
	db.Connect()

	// Create Fiber app
	app := fiber.New()

	// Set Routers
	pkgRouters.SetRouter(app)

	// Listen to port
	port := utils.GetEnvValue("APP_PORT", enums.DefaultAppPort)
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Error when listen to port 3000: %v", err)
		return
	}
}
