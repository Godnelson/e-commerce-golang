package main

import (
	"e-commerce-golang/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()
	app := fiber.New()

	connection := db.New()

	err := connection.Client.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"Hello": "World",
		})
	})

	_ = app.Listen(":" + os.Getenv("SERVER_PORT"))
}
