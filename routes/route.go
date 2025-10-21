package routes

import (
	"log"

	"github.com/Fadell-Karlsefni/project-management/config"
	"github.com/Fadell-Karlsefni/project-management/controllers"
	"github.com/Fadell-Karlsefni/project-management/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App,
	uc *controllers.UserController,bc *controllers.BoardController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)

	// JWT Protect routes
	api := app.Group("/api/v1", jwtware.New(jwtware.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
		ContextKey: "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.Unauthorized(c, "Erorr Unauthorized", err.Error())
		},
	}))

	userGroup := api.Group("/users")
	userGroup.Get("/page", uc.GetUserPagination)
	userGroup.Get("/:id", uc.GetUser)
	userGroup.Put("/:id", uc.UpdateUser)
	userGroup.Delete("/:id", uc.DeleteUser)
	
	boardGrup := api.Group("/boards")
	boardGrup.Post("/", bc.CreateBoard)
	boardGrup.Put("/:id", bc.UpdateBoard)
	boardGrup.Post("/:id/members", bc.AddBoardMember)
	boardGrup.Delete("/:id/members", bc.RemoveBoardMembers)
	boardGrup.Get("/my", bc.GettMyBoardPaginate)
}
