package repository

import (
	"fmt"
	"strings"

	airticket "github.com/SokoloSHA/AirTicket"
	"github.com/jmoiron/sqlx"
)

type TicketPostgreSQL struct {
	db *sqlx.DB
}

func NewTicketPostgreSQL(db *sqlx.DB) *TicketPostgreSQL {
	return &TicketPostgreSQL{db: db}
}

func (r *TicketPostgreSQL) GetAll() ([]airticket.Ticket, error) {
	var tickets []airticket.Ticket
	query := fmt.Sprintf("SELECT * FROM %s", ticketsTable)

	err := r.db.Select(&tickets, query)

	return tickets, err
}

func (r *TicketPostgreSQL) UpdateTicket(ticket airticket.Ticket) error {
	setValues := make([]string, 0)
	if ticket.PointStart != "" {
		setValues = append(setValues, fmt.Sprintf("point_start='%s'", ticket.PointStart))
	}

	if ticket.PointEnd != "" {
		setValues = append(setValues, fmt.Sprintf("point_end='%s'", ticket.PointEnd))
	}

	if ticket.OrderNumber != "" {
		setValues = append(setValues, fmt.Sprintf("order_number='%s'", ticket.OrderNumber))
	}

	if ticket.ServiceProvider != "" {
		setValues = append(setValues, fmt.Sprintf("service_provider='%s'", ticket.ServiceProvider))
	}

	if ticket.DateStart != "" {
		setValues = append(setValues, fmt.Sprintf("date_start='%s'", ticket.DateStart))
	}

	if ticket.DateFinish != "" {
		setValues = append(setValues, fmt.Sprintf("date_finish='%s'", ticket.DateFinish))
	}

	if ticket.CreatedAt != "" {
		setValues = append(setValues, fmt.Sprintf("created_at='%s'", ticket.CreatedAt))
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE Id = %d",
		ticketsTable, setQuery, ticket.Id)

	_, err := r.db.Exec(query)

	return err
}

func (r *TicketPostgreSQL) DeleteTicket(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE Id = %d", ticketsTable, id)
	_, err := r.db.Exec(query)

	return err
}

func (r *TicketPostgreSQL) GetTicket(ticketId int) (airticket.Ticket, error) {
	var ticket airticket.Ticket

	query := fmt.Sprintf("SELECT * FROM %s WHERE Id = %d", ticketsTable, ticketId)

	err := r.db.Get(&ticket, query)

	return ticket, err
}
