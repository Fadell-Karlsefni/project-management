package seed

import (
	"log"

	"github.com/Fadell-Karlsefni/project-management/config"
	"github.com/Fadell-Karlsefni/project-management/models"
	"github.com/Fadell-Karlsefni/project-management/utils"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin1123")
	admin := models.User{
		Name:     "Super admin",
		Email:    "Admin@example.com",
		Password: password,
		Role:     "admin",
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed to seed admin : ", err)
	} else {
		log.Println("Admin user seeded")
	}
}
