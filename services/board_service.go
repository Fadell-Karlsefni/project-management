package services

import (
	"errors"

	"github.com/Fadell-Karlsefni/project-management/models"
	"github.com/Fadell-Karlsefni/project-management/repositories"
	"github.com/google/uuid"
)

type BoardService interface {
	Create(board *models.Board) error
}

type boardService struct {
	boardRepo repositories.BoardRepository
	userRepo repositories.UserRepositroy
}

func NewBoardService(
	boardRepo repositories.BoardRepository,
	userRepo repositories.UserRepositroy,
	) BoardService {
	return &boardService{boardRepo,userRepo }
}

func (s *boardService) Create(board *models.Board) error {
	user,err :=  s.userRepo.FindPublicID(board.OwnerPublicID.String())
	if err != nil {
		return errors.New("owner not found")
	}
	board.PublicID = uuid.New()
	board.OwnerID = user.InternalID
	return s.boardRepo.Create(board)
}