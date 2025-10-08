package controllers

import (
	

	"github.com/Fadell-Karlsefni/project-management/models"
	"github.com/Fadell-Karlsefni/project-management/services"
	"github.com/Fadell-Karlsefni/project-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)
	
	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data",err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Registrasi Gagal",err.Error())
	}

	var userResp models.UserResponse
	_ =  copier.Copy(&userResp,&user)
	return utils.Success(ctx, "Register Sukses",userResp)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var body struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx,"Invalid Request",err.Error())
	}
	user,err := c.service.Login(body.Email,body.Password)
	if err != nil {
		return utils.Unauthorized(ctx,"Login Fail",err.Error())
	}
	token, _ := utils.GenerateToken(user.InternalID,user.Role,user.Email,user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)

	var userResp models.UserResponse
	_ =  copier.Copy(&userResp,&user)
	return utils.Success(ctx,"Login Succesful",fiber.Map{
		"access_token" : token,
		"refresh_token" : refreshToken,
		"user" : userResp,

	})
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user,err := c.service.GetByPublicID(id)
	if err != nil {
		return utils.NotFound(ctx,"Data Not Found",err.Error())
	}
	var userResp models.UserResponse
	err = copier.Copy(&userResp,&user)
	if err != nil {
		return utils.BadRequest(ctx,"Internal Server Error",err.Error())
	}
	return utils.Success(ctx,"Data Berhasil Di temmukan",userResp)
}