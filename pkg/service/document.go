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

func (s *DocumentService) GetAll(userId int) ([]airticket.Document, error) {
	return s.repo.GetAll(userId)
}

func (s *DocumentService) UpdateDocument(document airticket.Document) error {
	return s.repo.UpdateDocument(document)
}

func (s *DocumentService) DeleteDocument(id int) error {
	return s.repo.DeleteDocument(id)
}

type UserDocuments struct {
	User      airticket.User
	Documents []airticket.Document
}

func (s *DocumentService) GetAllDocumentsTicket(users []airticket.User) ([]UserDocuments, error) {
	var result []UserDocuments

	for _, user := range users {
		var documentsForUser UserDocuments

		documents, err := s.repo.GetAllDocumentsTicket(user)
		if err != nil {
			return nil, err
		}

		documentsForUser.User = user
		documentsForUser.Documents = documents

		result = append(result, documentsForUser)

	}

	return result, nil
}
