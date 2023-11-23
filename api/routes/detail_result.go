package routes

import (
	"mobile-legend-api/api/handlers"
	"mobile-legend-api/api/middleware"
	detailresult "mobile-legend-api/pkg/detail_result"

	"github.com/gofiber/fiber/v2"
)

func DetailResultRouter(app fiber.Router, service detailresult.DetailResultService) {
	app.Get("/", middleware.CheckApiKey, handlers.GetAllDetailResult(service))
	app.Get("/result/:id", middleware.CheckApiKey, handlers.GetDetailByResultId(service))
	app.Get("/:id", middleware.CheckApiKey, handlers.GetDetailById(service))
	app.Post("/", middleware.CheckApiKey, handlers.AddNewDetail(service))

	app.Delete("/:id", middleware.CheckApiKey, handlers.DeleteDetail(service))
}