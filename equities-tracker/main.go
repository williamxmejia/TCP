package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Tracker struct {
	Count int
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		var tracker Tracker
		tracker.Count = 0

		for i := 0; i < 5; i++ {
			tracker.Count += 1
		}

		return ctx.Render("index", fiber.Map{
			"Title": "Fiber",
			"Message": "Dynamic view",
			"Count": tracker.Count,
		})
	})
	
	app.Listen(":3000")

}
