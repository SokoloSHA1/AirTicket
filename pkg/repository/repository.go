package repository

import (
	airticket "github.com/SokoloSHA/AirTicket"

	"github.com/jmoiron/sqlx"
)

type TodoUser interface {
	GetAll(ticketId int) ([]airticket.User, error)
	UpdateUser(user airticket.User) error
	DeleteUser(id int) error
	GetReport(report airticket.ReportUser) ([]airticket.Report, error)
}

type TodoTicket interface {
	UpdateTicket(user airticket.Ticket) error
	DeleteTicket(id int) error
	GetAll() ([]airticket.Ticket, error)
	GetTicket(ticketId int) (airticket.Ticket, error)
}

type TodoDocument interface {
	GetAll(userId int) ([]airticket.Document, error)
	UpdateDocument(document airticket.Document) error
	DeleteDocument(id int) error
	GetAllDocumentsTicket(user airticket.User) ([]airticket.Document, error)
}

type Repository struct {
	TodoUser
	TodoTicket
	TodoDocument
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser:     NewUserPostgreSQL(db),
		TodoTicket:   NewTicketPostgreSQL(db),
		TodoDocument: NewDocumentPostgreSQL(db),
	}
}
