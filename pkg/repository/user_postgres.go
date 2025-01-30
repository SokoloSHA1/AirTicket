package repository

import (
	"fmt"
	"strings"

	airticket "github.com/SokoloSHA/AirTicket"
	"github.com/jmoiron/sqlx"
)

type UserPostgreSQL struct {
	db *sqlx.DB
}

func NewUserPostgreSQL(db *sqlx.DB) *UserPostgreSQL {
	return &UserPostgreSQL{db: db}
}

func (r *UserPostgreSQL) GetAll(ticketId int) ([]airticket.User, error) {
	var users []airticket.User

	query := fmt.Sprintf("SELECT Id, FirstName, LastName, MiddleName FROM %s us INNER JOIN %s tu WHERE tu.ticketId = %d",
		usersTable, ticketId)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UserPostgreSQL) UpdateUser(user airticket.User) error {
	setValues := make([]string, 0)
	if user.FirstName != "" {
		setValues = append(setValues, fmt.Sprintf("FirstName = %s", user.FirstName))
	}

	if user.LastName != "" {
		setValues = append(setValues, fmt.Sprintf("LastName = %s", user.LastName))
	}

	if user.MiddleName != "" {
		setValues = append(setValues, fmt.Sprintf("MiddleName = %s", user.MiddleName))
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE Id = %d",
		usersTable, setQuery, user.Id)

	_, err := r.db.Exec(query)

	return err
}

func (r *UserPostgreSQL) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE Id = %d", usersTable, id)
	_, err := r.db.Exec(query)

	return err
}

func (r *UserPostgreSQL) GetReport(report airticket.ReportUser) ([]airticket.Ticket, error) {
	//TODO: Получение данный для отчета
	return nil, nil
}
