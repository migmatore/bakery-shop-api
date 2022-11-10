package storage

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type AddressStorage struct {
	pool psql.AtomicPoolClient
}

func NewAddressStorage(pool psql.AtomicPoolClient) *AddressStorage {
	return &AddressStorage{pool: pool}
}

func (s *AddressStorage) DeliveryAddressCreate(ctx context.Context, deliveryAddress *core.CreateDeliveryAddressDTO) (int, error) {
	q := `INSERT INTO delivery_addresses(region, city, street, house_number, building_number, apartment_number)
		  VALUES ($1, $2, $3, $4, $5, $6)
          RETURNING delivery_address_id`

	var id int
	// TODO add delivery address
	if err := s.pool.QueryRow(
		ctx,
		q,
		deliveryAddress.Region,
		deliveryAddress.City,
		deliveryAddress.Street,
		deliveryAddress.HouseNumber,
		deliveryAddress.BuildingNumber,
		deliveryAddress.ApartmentNumber,
	).Scan(&id); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return 0, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return 0, err
	}

	return id, nil
}
