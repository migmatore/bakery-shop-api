package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
)

type CustomerStorage interface {
	FindOne(ctx context.Context, id int) *core.GetCustomreDTO
}

type CustomerService struct {
	storage CustomerStorage
}

func NewCustomerService(storage CustomerStorage) *CustomerService {
	return &CustomerService{storage: storage}
}

func (s *CustomerService) GetById(ctx context.Context, id int) *core.GetCustomreDTO {
	return s.storage.FindOne(ctx, id)
}
