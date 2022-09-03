package storage

import "gorm.io/gorm"

type CustomerStorage struct {
	db *gorm.DB
}

func NewCustomerStorage(db *gorm.DB) *CustomerStorage {
	return &CustomerStorage{db: db}
}
