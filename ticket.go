package airticket

type Ticket struct {
	Id              int
	PointStart      string `json:"pointStart"`
	PointEnd        string `json:"pointEnd"`
	OrderNumber     int    `json:"orderNumber"`
	ServiceProvider string `json:"serviceProvider"`
	DateStart       string `json:"dateStart"`
	DateFinish      string `json:"dateFinish"`
	CreatedAt       string `json:"createdAt"`
}
