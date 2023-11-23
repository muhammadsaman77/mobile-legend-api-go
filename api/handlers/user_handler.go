package handlers

import (
	"mobile-legend-api/api/presenter"
	"mobile-legend-api/pkg/entities"
	"mobile-legend-api/pkg/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Register(service user.UserService) fiber.Handler{
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err:= c.BodyParser(&requestBody)
		if err!=nil{
			panic(err)
		}
		 service.Register(c.Context(),&requestBody)
		return c.Status(fiber.StatusCreated).JSON(presenter.UserSuccessResponse())
	}
}

func Login (service user.UserService) fiber.Handler{
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err:= c.BodyParser(&requestBody)
		if err!=nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		token,err:= service.Login(c.Context(),&requestBody)
		if err!=nil {
			
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
		}
		return c.Status(fiber.StatusOK).JSON(presenter.UserSuccessLogin(token))
	}
}

func FindByEmail (service user.UserService) fiber.Handler{
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err:= c.BodyParser(&requestBody)
		if err!=nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		user,err:= service.FindByEmail(c.Context(),&requestBody)
		
		if err!=nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
		}
		return c.Status(fiber.StatusOK).JSON(presenter.FindByEmailResponse(*user))
	}
}


func SearchByName(service user.UserService) fiber.Handler{
	return func(c *fiber.Ctx) error {
		queryParam :=  c.Query("name")
		
		users,err:= service.SearchByName(c.Context(),queryParam)
		if err!= nil{
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
		}
		return c.Status(fiber.StatusOK).JSON(presenter.FindByNameResponse(*users))
	}
}

func DeleteUser(service user.UserService) fiber.Handler{
	return func(c *fiber.Ctx) error {
		idParams := c.Params("id")
		if len(idParams)<=0 {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		id,_ := strconv.Atoi(idParams)
		deletedUser, err:= service.DeleteUser(c.Context(),id)
		if err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())

		}
		return c.Status(fiber.StatusOK).JSON(deletedUser)
	}

}
