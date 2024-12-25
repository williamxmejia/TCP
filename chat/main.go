package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	viewsEngine := html.New("./views", ".html")

	app := fiber.New(fiber.Config {
		Views: viewsEngine,

	})

	app.Static("/static/", "./static")

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to fiber")
	})

	app.Listen(":3000")
}
