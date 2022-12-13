package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/api/filter"
	"github.com/migmatore/bakery-shop-api/pkg/api/pagination"
	"github.com/migmatore/bakery-shop-api/pkg/api/sort"
	"math"
)

type ProductStorage interface {
	FindOne(ctx context.Context, id int) (*core.Product, error)
	// TODO Refactor
	FindAll(ctx context.Context, filterOptions []filter.Option, sortOption sort.Option, pag pagination.Pagination) ([]*core.Product, error)
	Patch(ctx context.Context, id int, product *core.PatchProduct) (*core.Product, error)
	Create(ctx context.Context, product *core.CreateProduct) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int, error)
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

func (s *ProductService) GetAll(ctx context.Context, queryParams map[string]string) (*core.ProductPage, error) {
	filterOptions := filter.GetFilterOptions(queryParams)
	sortOption := sort.GetSortOptions(queryParams)
	pag := pagination.GetPaginationOptions(queryParams)

	products, err := s.storage.FindAll(ctx, filterOptions, sortOption, pag)
	if err != nil {
		return nil, err
	}

	total, err := s.storage.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &core.ProductPage{
		Products: products,
		Page:     pag.Page,
		Total:    total,
		LastPage: int(math.Ceil(float64(total / pag.PerPage))),
	}, nil
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
