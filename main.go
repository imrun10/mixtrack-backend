// main.go
package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to MixTrack Backend!")
	})
	app.Listen(":8080")
}
