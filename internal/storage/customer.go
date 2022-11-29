package storage

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type CustomerStorage struct {
	pool psql.AtomicPoolClient
}

func NewCustomerStorage(pool psql.AtomicPoolClient) *CustomerStorage {
	return &CustomerStorage{pool: pool}
}

func (s *CustomerStorage) Create(ctx context.Context, customer *core.CreateCustomer) (int, error) {
	q := `INSERT INTO customers(first_name, last_name, patronymic, image_path, phone_number, email, password_hash, 
                      delivery_address_id, cart_id, wish_list_id, created_at, updated_at) 
		  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
          RETURNING customer_id`

	var id int

	if err := s.pool.QueryRow(
		ctx,
		q,
		customer.FirstName,
		customer.LastName,
		customer.Patronymic,
		customer.ImagePath,
		customer.PhoneNumber,
		customer.Email,
		customer.PasswordHash,
		customer.DeliveryAddressId,
		customer.CartId,
		customer.WishListId,
		customer.CreatedAt,
		customer.UpdatedAt,
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

func (s *CustomerStorage) FindAccByEmail(ctx context.Context, email string) (*core.SigninCustomer, error) {
	q := `select customer_id, password_hash from customers where email=$1`
	customer := core.SigninCustomer{}

	if err := s.pool.QueryRow(ctx, q, email).Scan(&customer.CustomerId, &customer.PasswordHash); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return nil, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	return &customer, nil
}

func (s *CustomerStorage) FindOne(ctx context.Context, id int) (*core.Customer, error) {
	q := `select customer_id, first_name, last_name, patronymic, image_path, phone_number, email, password_hash, 
       delivery_address_id, cart_id, wish_list_id, created_at, updated_at from customers where customers.customer_id=$1`
	var customer core.Customer

	if err := s.pool.QueryRow(ctx, q, id).Scan(
		&customer.CustomerId,
		&customer.FirstName,
		&customer.LastName,
		&customer.Patronymic,
		&customer.ImagePath,
		&customer.PhoneNumber,
		&customer.Email,
		&customer.PasswordHash,
		&customer.DeliveryAddressId,
		&customer.CartId,
		&customer.WishListId,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return nil, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	return &customer, nil
}

func (s *CustomerStorage) FindAll(ctx context.Context) ([]*core.Customer, error) {
	q := `select customer_id, first_name, last_name, patronymic, image_path, phone_number, email, password_hash, 
       delivery_address_id, cart_id, wish_list_id, created_at, updated_at from customers`

	customers := make([]*core.Customer, 0)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		customer := core.Customer{}

		err := rows.Scan(
			&customer.CustomerId,
			&customer.FirstName,
			&customer.LastName,
			&customer.Patronymic,
			&customer.ImagePath,
			&customer.PhoneNumber,
			&customer.Email,
			&customer.PasswordHash,
			&customer.DeliveryAddressId,
			&customer.CartId,
			&customer.WishListId,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		customers = append(customers, &customer)
	}

	return customers, nil
}
