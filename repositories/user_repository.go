package repositories

import (
	"github.com/Fadell-Karlsefni/project-management/config"
	"github.com/Fadell-Karlsefni/project-management/models"
)

type UserRepositroy interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindPublicID(publicID string) (*models.User, error)

}

type userRepositroy struct{}

func NewUserRepository() UserRepositroy {
	return &userRepositroy{}
}

func (r *userRepositroy) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepositroy) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepositroy) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func (r *userRepositroy) FindPublicID(publicID string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("public_id = ?", publicID).First(&user).Error
	return &user, err
}
