package handlers

import (
	"mobile-legend-api/api/presenter"
	"mobile-legend-api/pkg/entities"
	"mobile-legend-api/pkg/hero"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllHero(service hero.HeroService)fiber.Handler{
	return func(c *fiber.Ctx) error {
		heroes,err:= service.GetAllHero(c.Context())
		if err!=nil{
			panic(err)
		}
		return c.Status(fiber.StatusOK).JSON(presenter.GetAllHeroResponses(*heroes))
	}
}

func GetDetailHero(service hero.HeroService) fiber.Handler{
	return func(c *fiber.Ctx) error {
		idParams := c.Params("id")
		id, err := strconv.Atoi(idParams)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		hero,err:= service.GetDetailHero(c.Context(),id)
		if err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
		}
		return c.Status(fiber.StatusOK).JSON(presenter.GetHeroResponse(*hero))
	}}
func AddNewHero (service hero.HeroService)fiber.Handler{
		return func(c *fiber.Ctx) error {
			var requestBody entities.Hero
			err:= c.BodyParser(&requestBody)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
			}
			hero, err:=service.AddNewHero(c.Context(),&requestBody)
			if err!=nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
			}
			return c.Status(fiber.StatusOK).JSON(presenter.CreateHeroResponse(*hero))
		}
	}
	func UpdateHero(service hero.HeroService)fiber.Handler{
		return func(c *fiber.Ctx) error {
			idParams := c.Params("id")
			id, err := strconv.Atoi(idParams)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
			}
			var requestBody entities.Hero
			err= c.BodyParser(&requestBody)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
			}
			requestBody.Id = id
			hero,err:= service.UpdateHero(c.Context(),&requestBody)
			if err!=nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
			}
			return c.Status(fiber.StatusOK).JSON(presenter.UpdateHeroResponse(*hero))
		}
	}
	func DeleteHero(service hero.HeroService) fiber.Handler{
		return func(c *fiber.Ctx) error {
			idParams := c.Params("id")
			if len(idParams)<=0 {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
			}
			id,_ := strconv.Atoi(idParams)
			err := service.DeleteHero(c.Context(),id)
			if err!=nil{
				panic(err)
			}
			return c.Status(fiber.StatusOK).JSON(presenter.DeleteHeroResponse())
		}
	}