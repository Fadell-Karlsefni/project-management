package main

import (
	"log"

	"github.com/Fadell-Karlsefni/project-management/config"
	"github.com/Fadell-Karlsefni/project-management/controllers"
	"github.com/Fadell-Karlsefni/project-management/database/seed"
	"github.com/Fadell-Karlsefni/project-management/repositories"
	"github.com/Fadell-Karlsefni/project-management/routes"
	"github.com/Fadell-Karlsefni/project-management/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	// untuk user
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)


	// untuk board
	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(boardRepo,userRepo,boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	// list
	listPosRepo := repositories.NewListPositionRepository()
	listRepo := repositories.NewListRepository()
	listService := services.NewServiceList(listRepo,boardRepo,listPosRepo)
	listController := controllers.NewListController(listService)
	


	routes.Setup(app,userController,boardController,listController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port : ", port)
	log.Fatal(app.Listen(":" + port))

}