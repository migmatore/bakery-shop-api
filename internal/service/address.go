package service

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
)

type AddressStorage interface {
	DeliveryAddressCreate(ctx context.Context, deliveryAddress *core.CreateDeliveryAddressDTO) (*int, error)
}

type AddressService struct {
	addressStorage AddressStorage
}

func NewAddressService(addressStorage AddressStorage) *AddressService {
	return &AddressService{addressStorage: addressStorage}
}

func (s *AddressService) DeliveryAddressCreate(ctx context.Context, deliveryAddress *core.CreateDeliveryAddressDTO) (*int, error) {
	return s.addressStorage.DeliveryAddressCreate(ctx, deliveryAddress)
}
