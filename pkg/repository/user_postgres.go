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

	query := fmt.Sprintf("SELECT us.id, us.first_name, us.last_name, us.middle_name FROM %s us INNER JOIN %s tu on us.id = tu.user_id WHERE tu.ticket_id = %d",
		usersTable, ticketUsersTable, ticketId)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UserPostgreSQL) UpdateUser(user airticket.User) error {
	setValues := make([]string, 0)
	if user.FirstName != "" {
		setValues = append(setValues, fmt.Sprintf("first_name = '%s'", user.FirstName))
	}

	if user.LastName != "" {
		setValues = append(setValues, fmt.Sprintf("last_name = '%s'", user.LastName))
	}

	if user.MiddleName != "" {
		setValues = append(setValues, fmt.Sprintf("middle_name = '%s'", user.MiddleName))
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

func (r *UserPostgreSQL) GetReport(report airticket.ReportUser) ([]airticket.Report, error) {
	var allTickets, tickets []airticket.Report

	query := fmt.Sprintf("SELECT ti.created_at, ti.date_start, ti.order_number, ti.point_start, ti.point_end, tu.done FROM %s ti INNER JOIN %s tu on ti.id = tu.ticket_id WHERE tu.user_id = %d AND ti.date_start >= '%s' AND ti.date_finish <= '%s' AND ti.created_at < '%s'",
		ticketsTable, ticketUsersTable, report.UserId, report.DateFrom, report.DateBefore, report.DateFrom)
	err := r.db.Select(&tickets, query)
	fmt.Println(query)
	if err != nil {
		return nil, err
	}

	allTickets = append(allTickets, tickets...)

	query = fmt.Sprintf("SELECT ti.created_at, ti.date_start, ti.order_number, ti.point_start, ti.point_end, tu.done FROM %s ti INNER JOIN %s tu on ti.id = tu.ticket_id WHERE tu.user_id = %d AND ti.created_at >= '%s' AND ti.created_at <= '%s'",
		ticketsTable, ticketUsersTable, report.UserId, report.DateFrom, report.DateBefore)
	err = r.db.Select(&tickets, query)
	fmt.Println(query)
	if err != nil {
		return nil, err
	}

	allTickets = append(allTickets, tickets...)

	return allTickets, err
}
