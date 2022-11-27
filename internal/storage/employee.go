package storage

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type EmployeeStorage struct {
	pool psql.AtomicPoolClient
}

func NewEmployeeStorage(pool psql.AtomicPoolClient) *EmployeeStorage {
	return &EmployeeStorage{pool: pool}
}

func (s *EmployeeStorage) Create(ctx context.Context, employee *core.CreateEmployee) (int, error) {
	q := `INSERT INTO employees(first_name, last_name, patronymic, phone_number, email, password_hash, position_id, company_id, admin) 
		  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
          RETURNING employee_id`

	var id int

	if err := s.pool.QueryRow(
		ctx,
		q,
		employee.FirstName,
		employee.LastName,
		employee.Patronymic,
		employee.PhoneNumber,
		employee.Email,
		employee.PasswordHash,
		employee.PositionId,
		employee.CompanyId,
		employee.Admin,
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
