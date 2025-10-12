package controllers

import (
	"github.com/Fadell-Karlsefni/project-management/models"
	"github.com/Fadell-Karlsefni/project-management/services"
	"github.com/Fadell-Karlsefni/project-management/utils"
	"github.com/gofiber/fiber/v2"
)

type BoardController struct {
	services services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController {
	return &BoardController{services: s}
}

func (c *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	board := new(models.Board)

	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx,"Gagal Membaca Request",err.Error())
	}
	if err := c.services.Create(board); err != nil {
		return utils.BadRequest(ctx,"Gagal Menyimpan Data",err.Error())
	}
	return utils.Success(ctx,"Board Berhasil Di buat",board)
}