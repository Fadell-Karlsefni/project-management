package services

import (
	"errors"

	"github.com/Fadell-Karlsefni/project-management/models"
	"github.com/Fadell-Karlsefni/project-management/repositories"
	"github.com/Fadell-Karlsefni/project-management/utils"
	"github.com/google/uuid"
)

type listService struct {
	listRepo repositories.ListRepository
	boardRepo repositories.BoardRepository
	listPosRepo repositories.ListPositionRepository
}

type ListWithOrder struct {
	Position []uuid.UUID
	Lists []models.List
}

type ListService interface {
	GetByBoardID(boardPublicID string) (*ListWithOrder,error)
	GetByID(id uint) (*models.List, error)
	GetByPublicID(publicID string) (*models.List,error)
	Create(list *models.List) error
	Update(list *models.List) error
	Delete(id uint) error
	UpdatePositions(boardPublicID string, position []uuid.UUID) error 
}

func NewServiceList(listRepo repositories.ListRepository, boardRepo repositories.BoardRepository, listPosRepo repositories.ListPositionRepository) ListService {
	return &listService{listRepo,boardRepo,listPosRepo,}
}

func (s *listService) GetByBoardID(boardPublicID string) (*ListWithOrder, error) {
	// verifikasi board

	_, err :=  s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return nil, errors.New("board not found")
	}

	positions, err := s.listPosRepo.GetListOrder(boardPublicID)
	if err != nil {
		return nil, errors.New("failed to get list order : " + err.Error())
	}

	lists, err := s.listRepo.FindByBoardID(boardPublicID)
	if err != nil {
		return nil, errors.New("failed to get list : " + err.Error())
	}

	//sorting by position
	orderedList :=  utils.SortListByPosition(lists,positions)
	return &ListWithOrder{
		Position: positions,
		Lists: orderedList,
	},nil
}

func (s *listService) GetByID(id uint) (*models.List, error) {
	return nil, nil
}

func (s *listService) GetByPublicID(publicID string) (*models.List, error) {
	return nil, nil
}

func (s *listService) Create(list *models.List) error {
	return nil
}

func (s *listService) Update(list *models.List) error {
	return nil
}

func (s *listService) Delete(id uint) error {
	return nil
}

func (s *listService) UpdatePositions(boardPublicID string, position []uuid.UUID) error {
	return nil
}
