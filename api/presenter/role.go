package presenter

import (
	"mobile-legend-api/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type RoleResponse struct {
	Status string        `json:"status"`
	Message string `json:"message"`
	Data   []fiber.Map `json:"data"`
}

func GetAllRoleResponse(roles []entities.Role) *RoleResponse {
	var roleData []fiber.Map

	for _, role := range roles {
		roleMap := fiber.Map{
			"id":   role.Id,
			"name": role.Name,

		}
		roleData = append(roleData, roleMap)
	}

	return &RoleResponse{
		Status: "OK",
		Message: "Get All Role Success",
		Data:   roleData,
	}
}

func DeleteRoleResponse() *fiber.Map{
	return &fiber.Map{
		"status": "OK",
		"message":"Delete Role Success",
	}
}

func UpdateRoleResponse(role entities.Role)  *fiber.Map{
	return &fiber.Map{
		"status": "OK",
		"message":"Update Role Success",
		"data": role,
	}
}


func CreateRoleResponse(role entities.Role)  *fiber.Map{
	return &fiber.Map{
		"status": "OK",
		"message":"Add New Role Success",
		"data": role,
	}
}