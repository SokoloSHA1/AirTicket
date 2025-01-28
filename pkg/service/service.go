package service

import (
	airticket "test-module/src/AirTicket"
	"test-module/src/AirTicket/pkg/repository"
)

type TodoUser interface {
	UpdateUser(user airticket.User) error
	DeleteUser(id int) error
}

type TodoTicket interface {
	UpdateTicket(user airticket.Ticket) error
	deleteTicket(id int) error
}

type TodoDocument interface {
}

type Service struct {
	TodoUser
	TodoTicket
	TodoDocument
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoUser:   NewUserService(repos.TodoUser),
		TodoTicket: NewTicketService(repos.TodoTicket),
	}
}
