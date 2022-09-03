package storage

import "gorm.io/gorm"

type Storage struct {
	Customer *CustomerStorage
}

func New(db *gorm.DB) *Storage {
	return &Storage{Customer: NewCustomerStorage(db)}
}
