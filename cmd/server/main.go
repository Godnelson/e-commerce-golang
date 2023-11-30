package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	_ = godotenv.Load()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"Hello": "World",
		})
	})

	_ = app.Listen(":" + os.Getenv("SERVER_PORT"))
}
