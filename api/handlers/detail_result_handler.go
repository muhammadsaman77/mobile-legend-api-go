package handlers

import (
	"mobile-legend-api/api/presenter"
	detailresult "mobile-legend-api/pkg/detail_result"
	"mobile-legend-api/pkg/entities"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllDetailResult(service detailresult.DetailResultService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		detailresults,err:= service.GetAllDetailResult(c.Context())
		if err!=nil{
			panic(err)
		}
		return c.Status(fiber.StatusOK).JSON(presenter.GetAllDetailResultResponses(*detailresults))
	}
}

func GetDetailByResultId(service detailresult.DetailResultService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParams := c.Params("id")
		id, err := strconv.Atoi(idParams)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		detailresult,err:= service.GetDetailByResultId(c.Context(),id)
		if err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
		}
		return c.Status(fiber.StatusOK).JSON(presenter.GetDetailResultByIdResponses(*detailresult))
}}

func GetDetailById(service detailresult.DetailResultService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParams := c.Params("id")
		id, err := strconv.Atoi(idParams)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		detailresult,err:= service.GetDetailById(c.Context(),id)
		if err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
		}
		return c.Status(fiber.StatusOK).JSON(presenter.GetDetailByIdResponse(*detailresult))
}}

func AddNewDetail(service detailresult.DetailResultService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DetailResult
		err:= c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		newDetail, err:=service.AddNewDetail(c.Context(),&requestBody)
		if err!=nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.InternalServer())
		}
		return c.Status(fiber.StatusOK).JSON(presenter.AddNewDetailResponse(*newDetail))
	}
}


func DeleteDetail(service detailresult.DetailResultService)fiber.Handler{
	return func(c *fiber.Ctx) error {
		idParams := c.Params("id")
		if len(idParams)<=0 {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BadRequest())
		}
		id,_ := strconv.Atoi(idParams)
		err := service.DeleteDetail(c.Context(),id)
		if err!=nil{
			panic(err)
		}
		return c.Status(fiber.StatusOK).JSON(presenter.DeleteDetailResponse())
	}
}