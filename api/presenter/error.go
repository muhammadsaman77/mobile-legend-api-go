package presenter

import "github.com/gofiber/fiber/v2"

func BadRequest() *fiber.Map{
	return &fiber.Map{
		"status":"error",
		"message":"Bad Request",
	}
}

func InternalServer() *fiber.Map{
	return &fiber.Map{
		"status":"error",
		"message":"Internal Server Error",
		
	}
}