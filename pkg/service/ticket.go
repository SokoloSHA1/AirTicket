package service

import (
	"fmt"

	airticket "github.com/SokoloSHA/AirTicket"
	"github.com/SokoloSHA/AirTicket/pkg/repository"
)

type TicketService struct {
	repo repository.TodoTicket
}

func NewTicketService(repo repository.TodoTicket) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) GetAll() ([]airticket.Ticket, error) {
	fmt.Println("1")
	return s.repo.GetAll()
}

func (s *TicketService) UpdateTicket(ticket airticket.Ticket) error {
	return s.repo.UpdateTicket(ticket)
}

func (s *TicketService) DeleteTicket(id int) error {
	return s.repo.DeleteTicket(id)
}

func (s *TicketService) GetTicket(ticketId int) (airticket.Ticket, error) {
	return s.repo.GetTicket(ticketId)
}
