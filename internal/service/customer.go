package service

type CustomerStorage interface {
}

type CustomerService struct {
	storage CustomerStorage
}

func NewCustomerService(storage CustomerStorage) *CustomerService {
	return &CustomerService{storage: storage}
}
