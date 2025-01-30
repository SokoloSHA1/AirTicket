package airticket

type Document struct {
	Id           int    `db:"id"`
	TypeDocument string `json:"typeDocument" db:"type_document"`
	Number       string `json:"number" db:"number"`
	UserId       int    `db:"user_id"`
}
