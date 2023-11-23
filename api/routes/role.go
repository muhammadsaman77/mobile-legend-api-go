package routes

import (
	"mobile-legend-api/api/handlers"
	"mobile-legend-api/api/middleware"
	"mobile-legend-api/pkg/role"

	"github.com/gofiber/fiber/v2"
)

func RoleRouter(app fiber.Router, service role.RoleService){
	app.Get("/",middleware.CheckApiKey, handlers.GetAllRole(service))
	app.Post("/",middleware.CheckApiKey, handlers.AddNewRole(service))
	app.Put("/:id",middleware.CheckApiKey, handlers.UpdateRole(service))
	app.Delete("/:id",middleware.CheckApiKey, handlers.DeleteRole(service))
}