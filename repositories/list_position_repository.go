package repositories

import (
	"github.com/Fadell-Karlsefni/project-management/config"
	"github.com/Fadell-Karlsefni/project-management/models"
)

type listPositionRepository struct{}

type ListPositionRepository interface {
	GetByBoard(boardPublicID string) (*models.ListPosition, error)
}

func NewListPositionRepository() ListPositionRepository {
	return &listPositionRepository{}
}

func (r *listPositionRepository) GetByBoard(boardPublicID string) (*models.ListPosition, error) {
	var position models.ListPosition

	err := config.DB.Joins("JOIN boards ON boards.internal_id = list_position.board_internal_id").
	Where("boards.public_id = ?", boardPublicID).Error

	return &position,err
}