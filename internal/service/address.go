package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
)

type AddressStorage interface {
	CreateDeliveryAddress(ctx context.Context, deliveryAddress *core.CreateDeliveryAddress) (*int, error)
}

type AddressService struct {
	addressStorage AddressStorage
}

func NewAddressService(addressStorage AddressStorage) *AddressService {
	return &AddressService{addressStorage: addressStorage}
}

func (s *AddressService) DeliveryAddressCreate(ctx context.Context, deliveryAddress *core.CreateDeliveryAddressDTO) (*int, error) {
	deliveryAddressModel := core.NewCreateDeliveryAddressFromDTO(deliveryAddress)
	return s.addressStorage.CreateDeliveryAddress(ctx, deliveryAddressModel)
}
