package controllers

import (
	"math"
	"strconv"

	"github.com/Fadell-Karlsefni/project-management/models"
	"github.com/Fadell-Karlsefni/project-management/services"
	"github.com/Fadell-Karlsefni/project-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type BoardController struct {
	service services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController {
	return &BoardController{service: s}
}

func (c *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	var userID uuid.UUID
	var err error

	board := new(models.Board)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx,"Gagal Membaca Request",err.Error())
	}

	userID,err = uuid.Parse(claims["pub_id"].(string))
	if err != nil {
		return utils.BadRequest(ctx,"Gagal Membaca request",err.Error())
	}
	board.OwnerPublicID = userID

	if err := c.service.Create(board); err != nil {
		return utils.BadRequest(ctx,"Gagal Menyimpan Data",err.Error())
	}
	return utils.Success(ctx,"Board Berhasil Di buat",board)
}

func (c *BoardController) UpdateBoard(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	board := new(models.Board)

	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx,"Gagal Parsing Data",err.Error())
	}

	if _,err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx,"ID Tidak Valid",err.Error())
	}

	existingBoard, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx,"Board Tidak DiTemukan",err.Error())
	}
	board.InternalID = existingBoard.InternalID
	board.PublicID = existingBoard.PublicID
	board.OwnerID = existingBoard.OwnerID
	board.OwnerPublicID = existingBoard.OwnerPublicID
	board.CreatedAt = existingBoard.CreatedAt


	if err := c.service.Update(board);err != nil {
		return utils.BadRequest(ctx,"Gagal Update Board",err.Error())
	}

	return  utils.Success(ctx,"Board Berhasil DiPerbarui",board)
	// 0dc6c520-8a39-4b81-a213-36a37f638834 id public board
}

func (c *BoardController) AddBoardMember(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")

	var userIDs []string
	if err := ctx.BodyParser(&userIDs); err != nil {
		return utils.BadRequest(ctx,"Gagal Pasing Data",err.Error())
	}

	if err := c.service.AddMember(publicID,userIDs); err != nil {
		return utils.BadRequest(ctx,"Gagal Menambahkan Members",err.Error())
	}

	return utils.Success(ctx,"Members Berhasil Di tambahkan", nil)
}

func (c *BoardController) RemoveBoardMembers(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")

	var userIDs []string
	if err := ctx.BodyParser(&userIDs); err != nil {
		return utils.BadRequest(ctx,"Gagal Pasing Data",err.Error())
	}

	if err := c.service.RemoveMembers(publicID,userIDs); err != nil {
		return utils.BadRequest(ctx,"Gagal Menghapus Members",err.Error())
	}

	return utils.Success(ctx,"Members Berhasil DiHapus", nil)
}

func (c *BoardController) GettMyBoardPaginate(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["pub_id"].(string)

	page, _ := strconv.Atoi(ctx.Query("page","1"))
	limit, _ := strconv.Atoi(ctx.Query("limit","10"))
	offset := (page - 1) * limit

	filter := ctx.Query("filter","")
	sort := ctx.Query("sort","")

	boards,total,err := c.service.GetAllByUserPaginate(userID,filter,sort,limit,offset)
	if err != nil {
		return utils.InternalServerError(ctx, "Gagal Mengambil Data Board",err.Error())
	}

	meta := utils.PaginationMeta{
		Page: page,
		Limit: limit,
		Total: int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
		Filter: filter,
		Sort: sort,
	}

	return utils.SuccessPagination(ctx, "Data Board Berhasil DiAmbil",boards,meta)
}