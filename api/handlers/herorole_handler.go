package handlers

import (
	"mobile-legend-api/api/presenter"
	"mobile-legend-api/pkg/entities"
	"mobile-legend-api/pkg/herorole"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllHeroRole(service herorole.HeroRoleService) fiber.Handler{
	return func(c *fiber.Ctx) error {
		heroroles,err:=service.GetAllHeroRole(c.Context())
		if err!=nil{
			panic(err)
		}
		return  c.Status(fiber.StatusOK).JSON(presenter.GetAllHeroRoleResponses(*heroroles))
	}
}

func GetHeroRoleById(service herorole.HeroRoleService) fiber.Handler{
	return func(c *fiber.Ctx) error {
		idParams := c.Params("id")
		id, err := strconv.Atoi(idParams)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		herorole,err:= service.GetHeroRoleById(c.Context(),id)
		if err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
		}
		return c.Status(fiber.StatusOK).JSON(presenter.GetHeroRoleResponse(*herorole))
	}}
func AddNewHeroRole(service herorole.HeroRoleService)fiber.Handler{
		return func(c *fiber.Ctx) error {
			var requestBody entities.HeroRole
			err:= c.BodyParser(&requestBody)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
			}
			herorole, err:=service.AddNewHeroRole(c.Context(),&requestBody)
			if err!=nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
			}
			return c.Status(fiber.StatusOK).JSON(presenter.CreateHeroRoleResponse(*herorole))
		}
	}
	func UpdateHeroRole(service herorole.HeroRoleService)fiber.Handler{
		return func(c *fiber.Ctx) error {
			idParams := c.Params("id")
			id, err := strconv.Atoi(idParams)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
			}
			var requestBody entities.HeroRole
			err= c.BodyParser(&requestBody)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
			}
			requestBody.Id = id
			herorole,err:= service.UpdateHeroRole(c.Context(),&requestBody)
			if err!=nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
			}
			return c.Status(fiber.StatusOK).JSON(presenter.UpdateHeroRoleResponse(*herorole))
		}
	}
	func DeleteHeroRole(service herorole.HeroRoleService) fiber.Handler{
		return func(c *fiber.Ctx) error {
			idParams := c.Params("id")
			if len(idParams)<=0 {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
			}
			id,_ := strconv.Atoi(idParams)
			err := service.DeleteHeroRole(c.Context(),id)
			if err!=nil{
				panic(err)
			}
			return c.Status(fiber.StatusOK).JSON(presenter.DeleteHeroRoleResponse())
		}
	}