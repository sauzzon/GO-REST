package main

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/sauzzon/GO-REST/database"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
