package core

type Customer struct {
	CustomerId        int     `json:"customer_id"`
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Patronymic        *string `json:"patronymic,omitempty"`
	TelephoneNumber   string  `json:"telephone_number"`
	Email             *string `json:"email,omitempty"`
	PasswordHash      *string `json:"password_hash,omitempty"`
	DeliveryAddressId *int    `json:"delivery_address_id,omitempty"`
}

type CreateCustomer struct {
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Patronymic        *string `json:"patronymic,omitempty"`
	TelephoneNumber   string  `json:"telephone_number"`
	Email             *string `json:"email,omitempty"`
	PasswordHash      *string `json:"password_hash,omitempty"`
	DeliveryAddressId *int    `json:"delivery_address_id,omitempty"`
}

type CreateCustomerWithAccountDTO struct {
	FirstName       string                    `json:"first_name"`
	LastName        string                    `json:"last_name"`
	Patronymic      *string                   `json:"patronymic,omitempty"`
	TelephoneNumber string                    `json:"telephone_number"`
	Email           string                    `json:"email"`
	Password        string                    `json:"password"`
	DeliveryAddress *CreateDeliveryAddressDTO `json:"delivery_address,omitempty"`
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

//	type CreateCustomerAccountDTO struct {
//		FirstName  string `json:"first_name"`
//		LastName   string `json:"last_name"`
//		Patronymic string `json:"patronymic"`
//		City       string `json:"city"`
//		Street     string `json:"street"`
//		House      string `json:"house"`
//		Apartment  string `json:"apartment"`
//	}
//type GetCustomreDTO struct {
//	ID   int    `json:"id"`
//	Name string `json:"name"`
//}
