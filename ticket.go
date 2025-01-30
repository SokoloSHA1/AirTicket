package airticket

type Ticket struct {
	Id              int    `json:"id" db:"id"`
	PointStart      string `json:"pointStart" db:"point_start"`
	PointEnd        string `json:"pointEnd" db:"point_end"`
	OrderNumber     string `json:"orderNumber" db:"order_number"`
	ServiceProvider string `json:"serviceProvider" db:"service_provider"`
	DateStart       string `json:"dateStart" db:"date_start"`
	DateFinish      string `json:"dateFinish" db:"date_finish"`
	CreatedAt       string `json:"createdAt" db:"created_at"`
}

type TicketUser struct {
	Id       int  `db:"Id"`
	TicketId int  `db:"ticket_id"`
	UserId   int  `db:"user_id"`
	Done     bool `db:"done"`
}

type Report struct {
	PointStart  string `db:"point_start"`
	PointEnd    string `db:"point_end"`
	OrderNumber string `db:"order_number"`
	DateStart   string `db:"date_start"`
	CreatedAt   string `db:"created_at"`
	Done        bool   `db:"done"`
}
