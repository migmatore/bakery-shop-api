package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/middleware"
	"github.com/migmatore/bakery-shop-api/internal/storage"
)

type StoreStorage interface {
	Create(ctx context.Context, store *core.CreateStore) (int, error)
}

type StoreEmployeeStorage interface {
	Create(ctx context.Context, employee *core.CreateEmployee) (int, error)
}

type StoreService struct {
	transactor      storage.Transactor
	customerStorage StoreStorage
	employeeStorage StoreEmployeeStorage
}

func NewStoreService(
	transactor storage.Transactor,
	customerStorage StoreStorage,
	employeeStorage StoreEmployeeStorage,
) *StoreService {
	return &StoreService{
		transactor:      transactor,
		customerStorage: customerStorage,
		employeeStorage: employeeStorage,
	}
}

func (s *StoreService) Create(ctx context.Context, store *core.CreateStoreDTO) (string, error) {
	var employeeId int
	var storeId int

	err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		var err error
		storeModel := core.NewCreateStoreFromDTO(store, nil, nil)

		storeId, err = s.customerStorage.Create(txCtx, storeModel)
		if err != nil {
			return err
		}

		employeeModel, err := core.NewCreateStoreAdminFromDTO(&store.Creator, nil, storeId)
		if err != nil {
			return err
		}

		employeeId, err = s.employeeStorage.Create(txCtx, employeeModel)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateNewAccessToken(employeeId, false, storeId, true)
	if err != nil {
		return "", nil
	}

	return token, nil
}
