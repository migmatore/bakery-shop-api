package service

type Deps struct {
	CustomerStorage CustomerStorage
}

type Service struct {
	Customer *CustomerService
}

func New(deps Deps) *Service {
	return &Service{
		Customer: NewCustomerService(deps.CustomerStorage),
	}
}
