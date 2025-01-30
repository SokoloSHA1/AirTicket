package repository

import (
	"fmt"
	"strings"

	airticket "github.com/SokoloSHA/AirTicket"
	"github.com/jmoiron/sqlx"
)

type DocumentPostgreSQL struct {
	db *sqlx.DB
}

func NewDocumentPostgreSQL(db *sqlx.DB) *UserPostgreSQL {
	return &UserPostgreSQL{db: db}
}

func (r *DocumentPostgreSQL) GetAll(userId int) ([]airticket.Document, error) {
	var documents []airticket.Document

	query := fmt.Sprintf("SELECT * FROM %s WHERE userId = %d",
		documentTable, userId)
	err := r.db.Select(&documents, query)

	return documents, err
}

func (r *DocumentPostgreSQL) UpdateDocument(document airticket.Document) error {
	setValues := make([]string, 0)
	if document.TypeDocument != "" {
		setValues = append(setValues, fmt.Sprintf("TypeDocument = %s", document.TypeDocument))
	}

	if document.Number != 0 {
		setValues = append(setValues, fmt.Sprintf("Number = %d", document.Number))
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE Id = %d",
		documentTable, setQuery, document.Id)

	_, err := r.db.Exec(query)

	return err
}

func (r *DocumentPostgreSQL) DeleteDocument(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE Id = %d", documentTable, id)
	_, err := r.db.Exec(query)

	return err
}

func (r *DocumentPostgreSQL) GetAllDocumentsTicket(users []airticket.User) ([]airticket.Document, error) {
	var allDocuments []airticket.Document

	for _, user := range users {
		var documents []airticket.Document

		query := fmt.Sprintf("SELECT * FROM %s WHERE UserId = %d", documentTable, user.Id)
		err := r.db.Select(&documents, query)
		if err != nil {
			return nil, err
		}

		allDocuments = append(allDocuments, documents...)
	}

	return allDocuments, nil
}
