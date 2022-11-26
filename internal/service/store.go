package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/storage"
)

type StoreStorage interface {
	Create(ctx context.Context, store *core.CreateStore) (int, error)
}

type StoreService struct {
	transactor storage.Transactor
	storage    StoreStorage
}

func NewStoreService(transactor storage.Transactor, storage StoreStorage) *StoreService {
	return &StoreService{transactor: transactor, storage: storage}
}

func (s *StoreService) Create(ctx context.Context, store *core.CreateStoreDTO) (string, error) {
	_ = s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		storeModel := core.NewCreateStoreFromDTO(store, nil, nil)

		_, err := s.storage.Create(txCtx, storeModel)
		if err != nil {
			return err
		}

		return nil
	})

	return "", nil
}
