package handlers

import (
	"mobile-legend-api/api/presenter"
	"mobile-legend-api/pkg/entities"
	"mobile-legend-api/pkg/role"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllRole(service role.RoleService)fiber.Handler{
	return func(c *fiber.Ctx) error {
		roles,err:= service.GetAllRole(c.Context())
		if err!=nil{
			panic(err)
		}
		return c.Status(fiber.StatusOK).JSON(presenter.GetAllRoleResponse(*roles))
	}
}

func AddNewRole(service role.RoleService)fiber.Handler{
	return func(c *fiber.Ctx) error {
		var requestBody entities.Role
		err:= c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		role, err:=service.AddNewRole(c.Context(),&requestBody)
		if err!=nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
		}
		return c.Status(fiber.StatusOK).JSON(presenter.CreateRoleResponse(*role))
	}
}
func UpdateRole(service role.RoleService)fiber.Handler{
	return func(c *fiber.Ctx) error {
		idParams := c.Params("id")
		id, err := strconv.Atoi(idParams)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		var requestBody entities.Role
		err= c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		requestBody.Id = id
		role,err:= service.UpdateRole(c.Context(),&requestBody)
		if err!=nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
		}
		return c.Status(fiber.StatusOK).JSON(presenter.UpdateRoleResponse(*role))
	}
}
func DeleteRole(service role.RoleService) fiber.Handler{
	return func(c *fiber.Ctx) error {
		idParams := c.Params("id")
		if len(idParams)<=0 {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		id,_ := strconv.Atoi(idParams)
		err := service.DeleteRole(c.Context(),id)
		if err!=nil{
			panic(err)
		}
		return c.Status(fiber.StatusOK).JSON(presenter.DeleteRoleResponse())
	}
}