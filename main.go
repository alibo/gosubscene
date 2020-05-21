package main

import (
	"github.com/alibo/gosubscene/pkg/handler"
	"github.com/gofiber/fiber"
	"net/http"
)

func main() {

	app := fiber.New()

	app.Get("/search", handler.Search)
	app.Get("/subtitles/:name", handler.List)
	app.Get("/subtitles/:name/farsi_persian/:id", handler.Details)
	app.Get("/download", handler.Download)
	app.Get("/healthz", func(c *fiber.Ctx) {
		c.Status(http.StatusOK).JSON(fiber.Map{
			"status": http.StatusOK,
		})
	})

	app.Listen(3000)
}
