package airticket

type Ticket struct {
	Id              int    `db:"id"`
	PointStart      string `json:"pointStart" db:"pointStart"`
	PointEnd        string `json:"pointEnd" db:"pointEnd"`
	OrderNumber     int    `json:"orderNumber" db:"orderNumber"`
	ServiceProvider string `json:"serviceProvider" db:"serviceProvider"`
	DateStart       string `json:"dateStart" db:"dateStart"`
	DateFinish      string `json:"dateFinish" db:"dateFinish"`
	CreatedAt       string `json:"createdAt" db:"createdAt"`
}
