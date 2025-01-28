package service

import (
	airticket "test-module/src/AirTicket"
	"test-module/src/AirTicket/pkg/repository"
)

type UserService struct {
	repo repository.TodoUser
}

func NewUserService(repo repository.TodoUser) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UpdateUser(user airticket.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
