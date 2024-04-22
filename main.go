package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Get("/hi", func(c *fiber.Ctx) error {
		return c.SendString("Hi World!")
	})

	app.Get("/greeting", func(c *fiber.Ctx) error {
		return c.SendString("Good Morning, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	ipAddress := "0.0.0.0"

	err := app.Listen(fmt.Sprintf("%s:%s", ipAddress, port))

	if err != nil {
		log.Fatal(err)
	}
}
