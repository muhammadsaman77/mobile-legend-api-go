package routes

import (
	"mobile-legend-api/api/handlers"
	"mobile-legend-api/api/middleware"
	"mobile-legend-api/pkg/herorole"

	"github.com/gofiber/fiber/v2"
)

func HeroRoleRouter(app fiber.Router, service herorole.HeroRoleService) {
	app.Get("/", middleware.CheckApiKey, handlers.GetAllHeroRole(service))
	app.Get("/:id", middleware.CheckApiKey, handlers.GetHeroRoleById(service))
	app.Post("/", middleware.CheckApiKey, handlers.AddNewHeroRole(service))
	app.Put("/:id", middleware.CheckApiKey, handlers.UpdateHeroRole(service))
	app.Delete("/:id", middleware.CheckApiKey, handlers.DeleteHeroRole(service))
}