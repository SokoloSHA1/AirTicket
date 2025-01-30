package airticket

type Document struct {
	Id           int    `db:"id"`
	TypeDocument string `json:"typeDocument" db:"typeDocument"`
	Number       int    `json:"number" db:"number"`
	UserId       int    `db:"userId"`
}
