package service

import (
	airticket "github.com/SokoloSHA/AirTicket"
	"github.com/SokoloSHA/AirTicket/pkg/repository"
)

type UserService struct {
	repo repository.TodoUser
}

func NewUserService(repo repository.TodoUser) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll(ticketId int) ([]airticket.User, error) {
	return s.repo.GetAll(ticketId)
}

func (s *UserService) UpdateUser(user airticket.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
