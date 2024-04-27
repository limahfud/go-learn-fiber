package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	ngrambe "github.com/limahfud/learn-go-fiber/ngawi/sine"
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
		greeting := ngrambe.ExampleFunction("ALI")
		return c.SendString("Good Morning, World! " + greeting + "Anything about this song " + os.Getenv("THE_SONG"))
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
