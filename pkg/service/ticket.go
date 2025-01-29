package service

import (
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
	return s.repo.GetAll()
}

func (s *TicketService) UpdateTicket(ticket airticket.Ticket) error {
	return s.repo.UpdateTicket(ticket)
}

func (s *TicketService) DeleteTicket(id int) error {
	return s.repo.DeleteTicket(id)
}
