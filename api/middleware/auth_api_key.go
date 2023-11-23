package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)




func CheckApiKey(c *fiber.Ctx) error{
	err:= godotenv.Load(".env")
	if err!= nil{
		log.Fatalf("Error loading .env file: %s", err)
	}
		apiKeyHeader := c.Get("X-API-Key")
	
		apiKeyEnv :=os.Getenv("API_KEY")
		if apiKeyHeader != apiKeyEnv {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: API key required",
			})
		}
		return c.Next();
}