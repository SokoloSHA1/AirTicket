package service

import (
	airticket "github.com/SokoloSHA/AirTicket"
	"github.com/SokoloSHA/AirTicket/pkg/repository"
)

type DocumentService struct {
	repo repository.TodoDocument
}

func NewDocumentService(repo repository.TodoDocument) *DocumentService {
	return &DocumentService{repo: repo}
}

func (s *DocumentService) UpdateDocument(document airticket.Document) error {
	return s.repo.UpdateDocument(document)
}

func (s *DocumentService) DeleteDocument(id int) error {
	return s.repo.DeleteDocument(id)
}
