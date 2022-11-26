package core

type Employee struct {
	EmployeeId   int     `json:"employee_id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Patronymic   *string `json:"patronymic,omitempty"`
	PhoneNumber  string  `json:"phone_number"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"password_hash"`
	PositionId   *int    `json:"position_id,omitempty"`
	CompanyId    int     `json:"company_id"`
	Admin        bool    `json:"admin"`
}

type CreateEmployee struct {
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Patronymic   *string `json:"patronymic,omitempty"`
	PhoneNumber  string  `json:"phone_number"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"password_hash"`
	PositionId   *int    `json:"position_id,omitempty"`
	CompanyId    int     `json:"company_id"`
	Admin        bool    `json:"admin"`
}

type CreateEmployeeDTO struct {
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Patronymic  *string `json:"patronymic,omitempty"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	PositionId  *int    `json:"position_id,omitempty"`
	CompanyId   int     `json:"company_id"`
	Admin       bool    `json:"admin"`
}
