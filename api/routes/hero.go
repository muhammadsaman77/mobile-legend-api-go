package routes

import (
	"mobile-legend-api/api/handlers"
	"mobile-legend-api/api/middleware"
	"mobile-legend-api/pkg/hero"

	"github.com/gofiber/fiber/v2"
)

func HeroRouter(app fiber.Router, service hero.HeroService) {
	app.Get("/", middleware.CheckApiKey, handlers.GetAllHero(service))
	app.Get("/:id", middleware.CheckApiKey, handlers.GetDetailHero(service))
	app.Post("/", middleware.CheckApiKey, handlers.AddNewHero(service))
	app.Put("/:id", middleware.CheckApiKey, handlers.UpdateHero(service))
	app.Delete("/:id", middleware.CheckApiKey, handlers.DeleteHero(service))
}