package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/middleware"
	"github.com/migmatore/bakery-shop-api/internal/storage"
)

type CustomerStorage interface {
	FindOne(ctx context.Context, id int) (*core.Customer, error)
	FindAll(ctx context.Context) ([]*core.Customer, error)
	Create(ctx context.Context, customer *core.CreateCustomer) (int, error)
}

type CustomerDeliveryAddressStorage interface {
	CreateDeliveryAddress(ctx context.Context, deliveryAddress *core.CreateDeliveryAddress) (*int, error)
}

type CustomerCartStorage interface {
	Create(ctx context.Context) (int, error)
}

type CustomerWishListStorage interface {
	Create(ctx context.Context) (int, error)
}

type CustomerService struct {
	transactor      storage.Transactor
	customerStorage CustomerStorage
	addressStorage  CustomerDeliveryAddressStorage
	cartStorage     CustomerCartStorage
	wishListStorage CustomerWishListStorage
}

func NewCustomerService(
	transactor storage.Transactor,
	customerStorage CustomerStorage,
	addressStorage CustomerDeliveryAddressStorage,
	cartStorage CustomerCartStorage,
	wishListStorage CustomerWishListStorage,
) *CustomerService {
	return &CustomerService{
		transactor:      transactor,
		customerStorage: customerStorage,
		addressStorage:  addressStorage,
		cartStorage:     cartStorage,
		wishListStorage: wishListStorage,
	}
}

func (s *CustomerService) Signup(ctx context.Context, customer *core.CreateCustomerDTO) (string, error) {
	var customerId int

	err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		var err error
		var deliveryAddressId *int

		if customer.DeliveryAddress != nil {
			deliveryAddress := core.NewCreateDeliveryAddressFromDTO(customer.DeliveryAddress)

			deliveryAddressId, err = s.addressStorage.CreateDeliveryAddress(txCtx, deliveryAddress)
			if err != nil {
				return err
			}
		}

		cartId, err := s.cartStorage.Create(txCtx)
		if err != nil {
			return err
		}

		wishListId, err := s.wishListStorage.Create(txCtx)
		if err != nil {
			return err
		}

		customerModel, err := core.NewCreateCustomerFromDTO(customer, deliveryAddressId, cartId, wishListId)
		if err != nil {
			return err
		}

		customerId, err = s.customerStorage.Create(txCtx, customerModel)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return "", nil
	}

	// TODO move GenerateNewAccessToken from middleware package
	token, err := middleware.GenerateNewAccessToken(customerId, true)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func (s *CustomerService) GetById(ctx context.Context, id int) (*core.Customer, error) {
	return s.customerStorage.FindOne(ctx, id)
}

func (s *CustomerService) GetAll(ctx context.Context) ([]*core.Customer, error) {
	return s.customerStorage.FindAll(ctx)
}
