package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/williamxmejia/TCP/chat/handlers"
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

	appHandler := NewAppHandler()

	app.Get("/", appHandler.HandleGetIndex)


	app.Listen(":3000")
}
