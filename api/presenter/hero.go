package presenter

import (
	"mobile-legend-api/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type AllHero struct {
	Id int
	Name string
}

func GetAllHeroResponses(hero []AllHero) *fiber.Map{
	return &fiber.Map{
		"message": "Success Get All Data hero",
		"data": hero,
	}
}

func GetHeroResponse(hero entities.Hero) *fiber.Map{
	return &fiber.Map{
		"message": "Success Get  Data hero",
		"data": hero,
	}
}

func CreateHeroResponse(hero entities.Hero)*fiber.Map{
	return &fiber.Map{
		"message": "Success Add New Data hero",
		"data": hero,
	}
}

func UpdateHeroResponse(hero entities.Hero)*fiber.Map{
	return &fiber.Map{
		"message": "Success Update Data Hero",
		"data": hero,
	}
}

func DeleteHeroResponse()*fiber.Map{
	return &fiber.Map{
		"message": "Success Delete Data Hero",
	}}