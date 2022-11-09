package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/middleware"
	"golang.org/x/crypto/bcrypt"
)

type CustomerStorage interface {
	WithTransaction(ctx context.Context, fn func(storage CustomerStorage) error) error
	FindOne(ctx context.Context, id int) (*core.Customer, error)
	FindAll(ctx context.Context) ([]*core.Customer, error)
	Create(ctx context.Context, customer *core.CreateCustomer) (int, error)
}

type DeliveryAddressStorage interface {
	DeliveryAddressCreate(ctx context.Context, deliveryAddress *core.CreateDeliveryAddressDTO) (int, error)
}

type CustomerService struct {
	customerStorage CustomerStorage
	addressStorage  DeliveryAddressStorage
}

func NewCustomerService(customerStorage CustomerStorage, addressStorage DeliveryAddressStorage) *CustomerService {
	return &CustomerService{customerStorage: customerStorage, addressStorage: addressStorage}
}

// TODO create transaction for customer
func (s *CustomerService) Signup(ctx context.Context, customer *core.CreateCustomerWithAccountDTO) (string, error) {
	err := s.customerStorage.WithTransaction(ctx, func(storage CustomerStorage) error {

	})

	var deliveryAddressId int

	if customer.DeliveryAddress != nil {
		var err error
		deliveryAddress := core.CreateDeliveryAddressDTO{
			Region:          customer.DeliveryAddress.Region,
			City:            customer.DeliveryAddress.City,
			Street:          customer.DeliveryAddress.Street,
			HouseNumber:     customer.DeliveryAddress.HouseNumber,
			BuildingNumber:  customer.DeliveryAddress.BuildingNumber,
			ApartmentNumber: customer.DeliveryAddress.ApartmentNumber,
		}

		deliveryAddressId, err = s.addressStorage.DeliveryAddressCreate(ctx, &deliveryAddress)
		if err != nil {
			return "", nil
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	strHash := string(hash)

	customerModel := core.CreateCustomer{
		FirstName:         customer.FirstName,
		LastName:          customer.LastName,
		Patronymic:        customer.Patronymic,
		TelephoneNumber:   customer.TelephoneNumber,
		Email:             &customer.Email,
		PasswordHash:      &strHash,
		DeliveryAddressId: &deliveryAddressId,
	}

	id, err := s.customerStorage.Create(ctx, &customerModel)
	if err != nil {
		return "", nil
	}
	// TODO move GenerateNewAccessToken from middleware package
	token, err := middleware.GenerateNewAccessToken(id, true)
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
