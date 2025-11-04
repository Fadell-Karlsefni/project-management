package controllers

import (
	"github.com/Fadell-Karlsefni/project-management/models"
	"github.com/Fadell-Karlsefni/project-management/services"
	"github.com/Fadell-Karlsefni/project-management/utils"
	"github.com/gofiber/fiber/v2"
)

type ListController struct {
	services services.ListService
}

func NewListController(s services.ListService) *ListController {
	return &ListController{services: s}
}

func (c *ListController) CreateList(ctx *fiber.Ctx) error {
	list := new(models.List)
	if err := ctx.BodyParser(list); err  != nil {
		return utils.BadRequest(ctx,"Gagal Membaca Request",err.Error())
	}
	if err := c.services.Create(list); err != nil {
		return utils.BadRequest(ctx,"Gagal Membuat List", err.Error())
	}
	return utils.Success(ctx,"List Berhasil Di buat",list)
}
