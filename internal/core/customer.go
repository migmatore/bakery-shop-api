package core

type Customer struct {
	CustomerId        int    `json:"customer_id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name,omitempty"`
	Patronymic        string `json:"patronymic,omitempty"`
	TelephoneNumber   string `json:"telephone_number"`
	AccountId         int    `json:"account_id"`
	DeliveryAddressId int    `json:"delivery_address_id"`
}

//type Customer struct {
//	ID         int    `json:"id"`
//	FirstName  string `json:"first_name"`
//	LastName   string `json:"last_name"`
//	Patronymic string `json:"patronymic"`
//	City       string `json:"city"`
//	Street     string `json:"street"`
//	House      string `json:"house"`
//	Apartment  string `json:"apartment"`
//}

//	type CreateCustomerDTO struct {
//		FirstName  string `json:"first_name"`
//		LastName   string `json:"last_name"`
//		Patronymic string `json:"patronymic"`
//		City       string `json:"city"`
//		Street     string `json:"street"`
//		House      string `json:"house"`
//		Apartment  string `json:"apartment"`
//	}
type GetCustomreDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
