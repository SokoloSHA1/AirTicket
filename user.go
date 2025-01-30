package airticket

type User struct {
	Id         int    `json:"id" db:"id"`
	FirstName  string `json:"firstName" db:"first_name"`
	LastName   string `json:"lastName" db:"last_name"`
	MiddleName string `json:"middleName" db:"middle_name"`
}

type ReportUser struct {
	UserId     int    `json:"userId" binding:"required"`
	DateFrom   string `json:"dateFrom" binding:"required"`
	DateBefore string `json:"dateBefore" binding:"required"`
}
