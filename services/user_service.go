package services

import (
	"errors"

	"github.com/Fadell-Karlsefni/project-management/models"
	"github.com/Fadell-Karlsefni/project-management/repositories"
	"github.com/Fadell-Karlsefni/project-management/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *models.User) error
	Login(email,password string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepositroy
}

func NewUserService(repo repositories.UserRepositroy) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {

	// Cek alamat email sudah terdaftar atau belum
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("email already registered")
	}

	// hash	password
	hased, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hased

	// set default role & publicID
	user.Role = "user"
	user.PublicID = uuid.New()

	// simpan ke DB lewat Repository
	return s.repo.Create(user)
}

func (s *userService) Login( email,password string) (*models.User,error) {
	user,err := s.repo.FindByEmail(email)
	if err != nil {
		return nil,errors.New("invalid credential")
	}
	if !utils.CheckPasswordHash(password,user.Password) {
		return nil,errors.New("invalid credential")
	}
	return user, nil
}
