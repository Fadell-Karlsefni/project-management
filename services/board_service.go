package services

import (
	"errors"

	"github.com/Fadell-Karlsefni/project-management/models"
	"github.com/Fadell-Karlsefni/project-management/repositories"
	"github.com/google/uuid"
)

type BoardService interface {
	Create(board *models.Board) error
	Update(board *models.Board) error
	GetByPublicID(publicID string) (*models.Board,error)
	AddMember(boardPublicID string, userPublicID []string) error
	RemoveMembers(boardPublicID string, userPublicIDs []string) error
	GetAllByUserPaginate(userID, filter, sort string, limit,offset int) ([]models.Board,int64,error)
}

type boardService struct {
	boardRepo repositories.BoardRepository
	userRepo repositories.UserRepositroy
	boardMemberRepo repositories.BoardMemberRepository
}

func NewBoardService(
	boardRepo repositories.BoardRepository,
	userRepo repositories.UserRepositroy,
	boardMemberRepo repositories.BoardMemberRepository,
	) BoardService {
	return &boardService{boardRepo,userRepo, boardMemberRepo}
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

func (s *boardService) Update(board *models.Board) error {
	return s.boardRepo.Update(board)
}

func (s *boardService) GetByPublicID(publicID string) (*models.Board,error) {
	return  s.boardRepo.FindByPublicID(publicID)
}

func (s *boardService) AddMember(boardPublicID string, userPublicID []string) error {
	board, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("board not found")
	}

	var userInternalIDs []uint
	for _, userPublicID := range userPublicID {
		user , err := s.userRepo.FindPublicID(userPublicID)
		if err != nil {
			return errors.New("user not found : " + userPublicID)
		}
		userInternalIDs = append(userInternalIDs, uint(user.InternalID))
	}
	// cek anggota
	existingMembers, err := s.boardMemberRepo.GetMembers(string(board.PublicID.String()))
	if err != nil{
		return err
	}

	// cek cepat menggunakan map
	memberMap := make(map[uint]bool)
	for _, member := range existingMembers {
		memberMap[uint(member.InternalID)] = true // member map[1] = true
	}

	var newMemberID []uint
	for _, userID := range userInternalIDs {
		if !memberMap[userID] {
			newMemberID = append(newMemberID, userID)
		}
	}
	if len(newMemberID) == 0 {
		return nil
	}

	return s.boardRepo.AddMember(uint(board.InternalID),newMemberID)
}

func (s *boardService) RemoveMembers(boardPublicID string, userPublicIDs []string) error {
	board, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("board not found")
	}

	// validasi user
	var userInternalIDs []uint
	for _, userPublicID := range userPublicIDs {
		user , err := s.userRepo.FindPublicID(userPublicID)
		if err != nil {
			return errors.New("user not found : " + userPublicID)
		}
		userInternalIDs = append(userInternalIDs, uint(user.InternalID))
	}

	// cek keanggotaan
	existingMembers, err := s.boardMemberRepo.GetMembers(string(board.PublicID.String()))
	if err != nil{
		return err
	}

	// cek cepat menggunakan map
	memberMap := make(map[uint]bool)
	for _, member := range existingMembers {
		memberMap[uint(member.InternalID)] = true // member map[1] = true
	}

	var membersToRemove []uint
	for _,userID := range userInternalIDs {
		if memberMap[userID] {
			membersToRemove = append(membersToRemove, userID)
		}
	}

	return s.boardRepo.RemoveMembers(uint(board.InternalID), membersToRemove)

}

func (s *boardService) GetAllByUserPaginate(userID, filter, sort string, limit, offset int) ([]models.Board,int64,error) {
	return s.boardRepo.FindAllByUserPaginate(userID,filter,sort,limit,offset)
}