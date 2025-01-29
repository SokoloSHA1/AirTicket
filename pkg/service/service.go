package service

import (
	airticket "github.com/SokoloSHA/AirTicket"

	"github.com/SokoloSHA/AirTicket/pkg/repository"
)

type TodoUser interface {
	UpdateUser(user airticket.User) error
	DeleteUser(id int) error
}

type TodoTicket interface {
	UpdateTicket(ticket airticket.Ticket) error
	DeleteTicket(id int) error
	GetAll() ([]airticket.Ticket, error)
}

type TodoDocument interface {
	UpdateDocument(document airticket.Document) error
	DeleteDocument(id int) error
}

type Service struct {
	TodoUser
	TodoTicket
	TodoDocument
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoUser:     NewUserService(repos.TodoUser),
		TodoTicket:   NewTicketService(repos.TodoTicket),
		TodoDocument: NewDocumentService(repos.TodoDocument),
	}
}
