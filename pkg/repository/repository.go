package repository

import (
	airticket "github.com/SokoloSHA/AirTicket"

	"github.com/jmoiron/sqlx"
)

type TodoUser interface {
	UpdateUser(user airticket.User) error
	DeleteUser(id int) error
}

type TodoTicket interface {
	UpdateTicket(user airticket.Ticket) error
	DeleteTicket(id int) error
	GetAll() ([]airticket.Ticket, error)
}

type TodoDocument interface {
	UpdateDocument(document airticket.Document) error
	DeleteDocument(id int) error
}

type Repository struct {
	TodoUser
	TodoTicket
	TodoDocument
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser: NewUserPostgreSQL(db),
	}
}
