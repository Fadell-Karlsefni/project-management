package services

import "github.com/Fadell-Karlsefni/project-management/repositories"

type ListService struct {
	listRepo repositories.ListRepository
	boardRepo repositories.BoardRepository
	
}