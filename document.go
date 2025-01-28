package airticket

type Document struct {
	Id           int
	TypeDocument string `json:"typeDocument"`
	Number       int    `json:"number"`
}
