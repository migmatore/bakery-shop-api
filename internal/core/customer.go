package core

type Customer struct {
	Common
	FirstName  string
	LastName   string
	Patronymic string
	City       string
	Street     string
	House      string
	Apartment  string
}

type CreateCustomerDTO struct {
	FirstName  string
	LastName   string
	Patronymic string
	City       string
	Street     string
	House      string
	Apartment  string
}
