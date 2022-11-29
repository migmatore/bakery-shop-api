package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/api/filter"
	"github.com/migmatore/bakery-shop-api/pkg/api/sort"
)

type ProductStorage interface {
	FindOne(ctx context.Context, id int) (*core.Product, error)
	FindAll(ctx context.Context, filterOptions []filter.Option, sortOption sort.Option) ([]*core.Product, error)
	Patch(ctx context.Context, id int, product *core.PatchProduct) (*core.Product, error)
	Create(ctx context.Context, product *core.CreateProduct) error
	Delete(ctx context.Context, id int) error
}

type ProductEmployeeStorage interface {
	FindOne()
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
	filterOptions := filter.GetFilterOptions(queryParams)
	sortOption := sort.GetSortOptions(queryParams)

	return s.storage.FindAll(ctx, filterOptions, sortOption)
}

func (s *ProductService) Patch(ctx context.Context, id int, product *core.PatchProductDTO) (*core.Product, error) {
	productModel := core.NewPatchProductFromDTO(product)

	return s.storage.Patch(ctx, id, productModel)
}

func (s *ProductService) Create(ctx context.Context, product *core.CreateProductDTO, employeeId int, storeId int) error {
	productModel := core.NewCreateProductFromDTO(product, storeId)

	if err := s.storage.Create(ctx, productModel); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) Delete(ctx context.Context, id int) error {
	_, err := s.storage.FindOne(ctx, id)
	if err != nil {
		return err
	}

	return s.storage.Delete(ctx, id)
}
