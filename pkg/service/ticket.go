package service

import (
	airticket "test-module/src/AirTicket"
	"test-module/src/AirTicket/pkg/repository"
)

type TicketService struct {
	repo repository.TodoTicket
}

func NewTicketService(repo repository.TodoTicket) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) UpdateTicket(user airticket.Ticket) error {
	return s.repo.UpdateTicket(user)
}

func (s *TicketService) deleteTicket(id int) error {
	return s.repo.deleteTicket(id)
}
