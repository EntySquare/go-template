package routing

import (
	"github.com/gofiber/fiber/v2"
	"go-template/service/app"
)

func Setup(f *fiber.App) {
	appApi := f.Group("/api")

	appApi.Post("/example", app.Example)
}
