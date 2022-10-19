package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
)

type ProductStorage interface {
	FindAll(ctx context.Context) ([]*core.Product, error)
}

type ProductService struct {
	storage ProductStorage
}

func NewProductService(storage ProductStorage) *ProductService {
	return &ProductService{storage: storage}
}

func (s *ProductService) GetAll(ctx context.Context) ([]*core.Product, error) {
	return s.storage.FindAll(ctx)
}
