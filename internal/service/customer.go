package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/middleware"
	"golang.org/x/crypto/bcrypt"
)

type CustomerStorage interface {
	FindOne(ctx context.Context, id int) (*core.Customer, error)
	FindAll(ctx context.Context) ([]*core.Customer, error)
	Create(ctx context.Context, customer *core.CreateCustomerWithAccountDTO) (int, error)
}

type CustomerService struct {
	storage CustomerStorage
}

func NewCustomerService(storage CustomerStorage) *CustomerService {
	return &CustomerService{storage: storage}
}

func (s *CustomerService) Signup(ctx context.Context, customer *core.CreateCustomerWithAccountDTO) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	customer.Password = string(hash)

	if customer.DeliveryAddress != nil {
		// create delivery address record
	}

	id, err := s.storage.Create(ctx, customer)
	if err != nil {
		return "", nil
	}
	// TODO move GenerateNewAccessTOken from middleware package
	token, err := middleware.GenerateNewAccessToken(id, true)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func (s *CustomerService) GetById(ctx context.Context, id int) (*core.Customer, error) {
	return s.storage.FindOne(ctx, id)
}

func (s *CustomerService) GetAll(ctx context.Context) ([]*core.Customer, error) {
	return s.storage.FindAll(ctx)
}
