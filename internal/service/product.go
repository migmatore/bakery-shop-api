package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/api/filter"
	"github.com/migmatore/bakery-shop-api/pkg/api/sort"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
)

type ProductStorage interface {
	FindOne(ctx context.Context, id int) (*core.Product, error)
	FindAll(ctx context.Context, filterOptions []filter.Option, sortOption sort.Option) ([]*core.Product, error)
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

func (s *ProductService) GetAll(ctx context.Context, queryParams map[string]string) ([]*core.Product, error) {
	logging.GetLogger(ctx).Infof("%v", queryParams)

	filterOptions := filter.GetFilterOptions(queryParams)

	logging.GetLogger(ctx).Infof("%v", filterOptions)

	sortOption := sort.GetSortOptions(queryParams)

	logging.GetLogger(ctx).Infof("%s %s", sortOption.Column, sortOption.Order)

	return s.storage.FindAll(ctx, filterOptions, sortOption)
}
