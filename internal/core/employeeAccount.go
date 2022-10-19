package core

type EmployeeAccount struct {
	EmployeeId   int    `json:"employee_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
