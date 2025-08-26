package server

import "github.com/gofiber/fiber/v3"

func Run(port string) {
	var app = fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Index route")
	})

	app.Listen(":"+port, fiber.ListenConfig{EnablePrefork: true})
}
