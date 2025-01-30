package airticket

type User struct {
	Id         int    `json:"id" db:"id"`
	FirstName  string `json:"firstName" db:"firstName"`
	LastName   string `json:"lastName" db:"lastName"`
	MiddleName string `json:"middleName" db:"middleName"`
}

type ReportUser struct {
	UserId     string `json:"userId"`
	DateFrom   string `json:"dateFrom"`
	DateBefore string `json:"dateBefore"`
}
