package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
)

type ProductStorage interface {
	FindOne(ctx context.Context, id int) (*core.Product, error)
	FindAll(ctx context.Context) ([]*core.Product, error)
}

type ProductService struct {
	storage ProductStorage
}

func NewProductService(storage ProductStorage) *ProductService {
	return &ProductService{storage: storage}
}

func (s *ProductService) GetOne(ctx context.Context, id int) (*core.Product, error) {
	return s.storage.FindOne(ctx, id)
}

func (s *ProductService) GetAll(ctx context.Context) ([]*core.Product, error) {
	return s.storage.FindAll(ctx)
}
