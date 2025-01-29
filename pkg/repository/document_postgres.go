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
