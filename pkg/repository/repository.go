package repository

import (
	airticket "test-module/src/AirTicket"

	"github.com/jmoiron/sqlx"
)

type TodoUser interface {
	UpdateUser(user airticket.User) error
	DeleteUser(id int) error
}

type TodoTicket interface {
}

type TodoDocument interface {
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
