package service

import (
	"context"
	"errors"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/middleware"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeStorage interface {
	Create(ctx context.Context, employee *core.CreateEmployee) (int, error)
	FindAll(ctx context.Context) ([]*core.Employee, error)
	GetAccByEmail(ctx context.Context, email string) (*core.SigninEmployee, error)
}

type EmployeeService struct {
	storage EmployeeStorage
}

func NewEmployeeService(storage EmployeeStorage) *EmployeeService {
	return &EmployeeService{storage: storage}
}

func (s *EmployeeService) Signin(ctx context.Context, employeeAcc *core.SigninEmployeeDTO) (string, error) {
	acc, err := s.storage.GetAccByEmail(ctx, employeeAcc.Email)
	if err != nil {
		return "", errors.New("employee not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(employeeAcc.Password)); err != nil {
		return "", errors.New("incorrect password")
	}

	token, err := middleware.GenerateNewAccessToken(acc.EmployeeId, false, acc.CompanyId, acc.Admin)
	if err != nil {
		return "", errors.New("token generation error")
	}

	return token, nil
}

func (s *EmployeeService) GetAll(ctx context.Context) ([]*core.Employee, error) {
	return s.storage.FindAll(ctx)
}
