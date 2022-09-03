package handler

type CustomerService interface {
}

type CustomerHandler struct {
	service CustomerService
}

func NewCustomerHandler(s CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}
