package repository

import (
	"fmt"
	"strings"
	airticket "test-module/src/AirTicket"

	"github.com/jmoiron/sqlx"
)

type UserPostgreSQL struct {
	db *sqlx.DB
}

func NewUserPostgreSQL(db *sqlx.DB) *UserPostgreSQL {
	return &UserPostgreSQL{db: db}
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
	query := fmt.Sprintf("UPDATE %s SET %s WHERE Id = %s",
		usersTable, setQuery, user.Id)

	_, err := r.db.Exec(query)

	return err
}

func (r *UserPostgreSQL) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE Id = %d", usersTable, id)
	_, err := r.db.Exec(query)

	return err
}
