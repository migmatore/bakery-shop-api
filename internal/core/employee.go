package core

type Employee struct {
	EmployeeId      int     `json:"employee_id"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Patronymic      *string `json:"patronymic,omitempty"`
	TelephoneNumber string  `json:"telephone_number"`
	AccountId       int     `json:"account_id"`
	PositionId      int     `json:"position_id"`
	CompanyId       int     `json:"company_id"`
}
