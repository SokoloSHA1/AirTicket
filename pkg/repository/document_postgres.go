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

func NewDocumentPostgreSQL(db *sqlx.DB) *DocumentPostgreSQL {
	return &DocumentPostgreSQL{db: db}
}

func (r *DocumentPostgreSQL) GetAll(userId int) ([]airticket.Document, error) {
	var documents []airticket.Document

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = %d",
		documentTable, userId)
	err := r.db.Select(&documents, query)

	return documents, err
}

func (r *DocumentPostgreSQL) UpdateDocument(document airticket.Document) error {
	setValues := make([]string, 0)
	if document.TypeDocument != "" {
		setValues = append(setValues, fmt.Sprintf("type_document = '%s'", document.TypeDocument))
	}

	if document.Number != "" {
		setValues = append(setValues, fmt.Sprintf("number = '%s'", document.Number))
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE Id = %d",
		documentTable, setQuery, document.Id)
	fmt.Println(query)
	_, err := r.db.Exec(query)

	return err
}

func (r *DocumentPostgreSQL) DeleteDocument(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE Id = %d", documentTable, id)
	_, err := r.db.Exec(query)

	return err
}

func (r *DocumentPostgreSQL) GetAllDocumentsTicket(user airticket.User) ([]airticket.Document, error) {

	var documents []airticket.Document

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = %d", documentTable, user.Id)
	err := r.db.Select(&documents, query)
	if err != nil {
		return nil, err
	}

	return documents, nil
}
