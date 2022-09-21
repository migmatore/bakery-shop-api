package core

type Customer struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Patronymic string `json:"patronymic"`
	City       string `json:"city"`
	Street     string `json:"street"`
	House      string `json:"house"`
	Apartment  string `json:"apartment"`
}

type CreateCustomerDTO struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Patronymic string `json:"patronymic"`
	City       string `json:"city"`
	Street     string `json:"street"`
	House      string `json:"house"`
	Apartment  string `json:"apartment"`
}

type GetCustomreDTO struct {
	ID   int `json:"id"`
	Name string
}
