package repositories

import (
	"github.com/Fadell-Karlsefni/project-management/config"
	"github.com/Fadell-Karlsefni/project-management/models"
	"github.com/google/uuid"
)

type ListRepository interface{
	Create(list *models.List) error
	Update(list *models.List) error
	Delete(id uint) error
	UpdatePosition(boardPublicID string, position []string) error
	GetCardPosition(listPublicID string) ([]uuid.UUID, error)
}

type listRepository struct{}

func NewListRepository() ListRepository {
	return &listRepository{}
}

func (r *listRepository) Create(list *models.List) error {
	return config.DB.Create(list).Error
}

func (r *listRepository) Update(list *models.List) error {
	return config.DB.Model(&models.List{}).Where("public_id = ?",list.PublicID).
	Updates(map[string]interface{}{
		"title": list.Title,
	}).Error
}

func (r *listRepository) Delete(id uint) error {
	return config.DB.Delete(&models.List{},id).Error
}

func (r *listRepository) UpdatePosition(boardPublicID string, position []string) error {
	return config.DB.Model(&models.ListPosition{}).
	Where("board_internal_id = (Select internal_id FROM boards Where public_id = ?)",boardPublicID).
	Update("list_order", position).Error
}

func (r *listRepository) GetCardPosition(listPublicID string) ([]uuid.UUID, error) {
	var position models.CardPosition
	err :=  config.DB.Joins("JOIN lists ON list.internal_id = card_positions.list_internal_id").
	Where("list.public_id = ?", listPublicID).Error
	return position.CardOrder,err
}