package presenter

import (
	"mobile-legend-api/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type AllDetailResult struct {
	Id       int
	ResultId int
	Hero     string
	User     string
	Grade    string
}

type DetailResult struct {
	Id         int
	ResultId   int
	Hero       string
	User       string
	LevelUser  int
	KillUser   int
	DeathUser  int
	AssistUser int
	Gold       float32
	Skor       string
	Grade      string
}

func GetAllDetailResultResponses(detailresult []AllDetailResult) *fiber.Map {
	return &fiber.Map{
		"message": "Success Get All Data Detail Result",
		"data":    detailresult,
	}
}


func GetDetailResultByIdResponses(detailresult []DetailResult) *fiber.Map {
	return &fiber.Map{
		"message": "Success Get All Data Detail Result By Result Id",
		"data":    detailresult,
	}
}

func GetDetailByIdResponse(detailresult DetailResult) *fiber.Map {
	return &fiber.Map{
		"message": "Success Get All Data Detail Result By  Id",
		"data":    detailresult,
	}
}

func AddNewDetailResponse(detailResult entities.DetailResult) *fiber.Map {
	return &fiber.Map{
		"message": "Success Add New Data Detail Result ",
		"data":    detailResult,
	}
}

func UpdatedDetailResponse(detailResult entities.DetailResult) *fiber.Map {
	return &fiber.Map{
		"message": "Success Add New Data Detail Result ",
		"data":    detailResult,
	}
}


func DeleteDetailResponse() *fiber.Map {
	return &fiber.Map{
		"message": "Success Delete Data Detail Result ",
	}
}
