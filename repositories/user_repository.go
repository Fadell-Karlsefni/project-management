package repositories

import (
	"strings"

	"github.com/Fadell-Karlsefni/project-management/config"
	"github.com/Fadell-Karlsefni/project-management/models"
)

type UserRepositroy interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindPublicID(publicID string) (*models.User, error)
	FindAllPagination(filter, sort string, limit, offset int) ([]models.User, int64, error)
	Update(user *models.User) error
	Delete(id uint) error
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

func (r *userRepositroy) FindAllPagination(filter, sort string, limit, offset int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	db := config.DB.Model(&models.User{})

	// filltering
	if filter != "" {
		filterPattern := "%" + filter + "%"
		db = db.Where("name Ilike ? OR email Ilike ?", filterPattern, filterPattern)
	}

	// count data
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// sorting
	if sort != "" {
		switch sort {
		case "-id":
			sort = "-internal_id"
		case "id":
			sort = "internal_id"
		}

		if strings.HasPrefix(sort, "-") {
			sort = strings.TrimPrefix(sort, "-") + " DESC"
		} else {
			sort += " ASC"
		}

		db = db.Order(sort)
	}

	// pagination
	err := db.Limit(limit).Offset(offset).Find(&users).Error
	return users, total, err

}

func (r *userRepositroy) Update(user *models.User) error {
	return config.DB.Model(&models.User{}).Where("public_id = ?",user.PublicID).Updates(map[string]interface{}{
		"name": user.Name,
	}).Error
}

func (r *userRepositroy) Delete(id uint) error {
	return config.DB.Delete(&models.User{},id).Error
}
