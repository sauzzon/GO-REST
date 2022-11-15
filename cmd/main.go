package main

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/sauzzon/GO-REST/database"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("HelloWorld from Docker Container!!")
	})

	app.Listen(":3000")
}
