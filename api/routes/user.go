package routes

import (
	"mobile-legend-api/api/handlers"
	"mobile-legend-api/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.UserService){
	app.Post("/register",handlers.Register(service))
	app.Post("/login",handlers.Login(service))
	app.Get("/searchByName",handlers.SearchByName(service))
	app.Get("/findByEmail",handlers.FindByEmail(service))
	// app.Put("/")
	// app.Delete("")
}