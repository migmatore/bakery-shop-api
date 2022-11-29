package core

import "golang.org/x/crypto/bcrypt"

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

type CreateStoreAdminDTO struct {
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Patronymic  *string `json:"patronymic,omitempty"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	PositionId  *int    `json:"position_id,omitempty"`
}

type SigninEmployeeDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninEmployee struct {
	EmployeeId   int
	PasswordHash string
	CompanyId    int
	Admin        bool
}

func NewCreateStoreAdminFromDTO(dto *CreateStoreAdminDTO, positionId *int, companyId int) (*CreateEmployee, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &CreateEmployee{
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		Patronymic:   dto.Patronymic,
		PhoneNumber:  dto.PhoneNumber,
		Email:        dto.Email,
		PasswordHash: string(hash),
		PositionId:   positionId,
		CompanyId:    companyId,
		Admin:        true,
	}, nil
}

func NewCreateEmployeeFromDTO(dto *CreateEmployeeDTO, positionId *int, companyId int) (*CreateEmployee, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &CreateEmployee{
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		Patronymic:   dto.Patronymic,
		PhoneNumber:  dto.PhoneNumber,
		Email:        dto.Email,
		PasswordHash: string(hash),
		PositionId:   positionId,
		CompanyId:    companyId,
		Admin:        dto.Admin,
	}, nil
}
