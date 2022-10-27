package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
)

type ProductStorage interface {
	FindOne(ctx context.Context, id int) (*core.Product, error)
	FindAll(ctx context.Context, queryParams map[string]string) ([]*core.Product, error)
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

	SortOptions := GetSortOptions(queryParams)

	logging.GetLogger(ctx).Infof("%v", SortOptions)

	return s.storage.FindAll(ctx, queryParams)
}

type SortOption struct {
	Column string
	Order  string
}

func GetSortOptions(queryParams map[string]string) []SortOption {
	sortOptions := make([]SortOption, 0)

	//for key, value := range queryParams {
	//	if key == "sort_by" || key == "sort_order" {
	//		sortOptions = append(sortOptions, SortOption{Column: key, Order: value})
	//	}
	//}
	if col, ok := queryParams["sort_by"]; ok || col != "" {
		if order, ok := queryParams["sort_order"]; ok || col != "" {
			sortOptions = append(sortOptions, SortOption{Column: col, Order: order})
		}

		sortOptions = append(sortOptions, SortOption{Column: col, Order: "asc"})
	}

	return sortOptions
}
