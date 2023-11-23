package presenter

import (
	"mobile-legend-api/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func GetAllHeroRoleResponses(heroroles []entities.HeroRole) *fiber.Map{
	return &fiber.Map{
		"message": "Success Get All Data Herorole",
		"data": heroroles,
	}
}

func GetHeroRoleResponse(herorole entities.HeroRole) *fiber.Map{
	return &fiber.Map{
		"message": "Success Get All Data Herorole",
		"data": herorole,
	}
}

func CreateHeroRoleResponse(herorole entities.HeroRole)*fiber.Map{
	return &fiber.Map{
		"message": "Success Add New Data Herorole",
		"data": herorole,
	}
}

func UpdateHeroRoleResponse(herorole entities.HeroRole)*fiber.Map{
	return &fiber.Map{
		"message": "Success Update Data Herorole",
		"data": herorole,
	}
}

func DeleteHeroRoleResponse()*fiber.Map{
	return &fiber.Map{
		"message": "Success Delete Data Herorole",
	}}