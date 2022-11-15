package main

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/sauzzon/GO-REST/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
