package presenter

import (
	"mobile-legend-api/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func UserSuccessResponse() *fiber.Map{
	return &fiber.Map{
		"status":true,
		"message":"success add data",
	}
}

func UserSuccessLogin(token string)*fiber.Map{
	return &fiber.Map{
		"status":true,
		"message":"success login",
		"token": token,
	}
}

func FindByNameResponse(users [] entities.User) *fiber.Map{
	return &fiber.Map{
		"status":true,
		"message": "Success Get Data User By Name",
		"data": users,
	}
}

func DeleteUserResponse(users entities.User)*fiber.Map{
	return &fiber.Map{
		"status":"OK",
		"message": "Success Delete Data User",
		"data":users,
	}
}


func FindByEmailResponse(user  entities.User) *fiber.Map{
	return &fiber.Map{

		"message": "Success Get Data User By Name",
		"data": user,
	}
}
