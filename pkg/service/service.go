package service

import (
	airticket "github.com/SokoloSHA/AirTicket"

	"github.com/SokoloSHA/AirTicket/pkg/repository"
)

type TodoUser interface {
	GetAll(ticketId int) ([]airticket.User, error)
	UpdateUser(user airticket.User) error
	DeleteUser(id int) error
	GetReport(report airticket.ReportUser) ([]airticket.Ticket, error)
}

type TodoTicket interface {
	UpdateTicket(ticket airticket.Ticket) error
	DeleteTicket(id int) error
	GetAll() ([]airticket.Ticket, error)
	GetTicket(ticketId int) (airticket.Ticket, error)
}

type TodoDocument interface {
	GetAll(userId int) ([]airticket.Document, error)
	UpdateDocument(document airticket.Document) error
	DeleteDocument(id int) error
	GetAllDocumentsTicket(users []airticket.User) ([]airticket.Document, error)
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
