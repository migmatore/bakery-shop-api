package core

type DeliveryAddress struct {
	DeliveryAddressId int     `json:"delivery_address_id"`
	Region            string  `json:"region"`
	City              string  `json:"city"`
	Street            string  `json:"street"`
	HouseNumber       string  `json:"house_number"`
	BuildingNumber    *string `json:"building_number,omitempty"`
	ApartmentNumber   string  `json:"apartment_number"`
}

type CreateDeliveryAddressDTO struct {
	Region          string  `json:"region"`
	City            string  `json:"city"`
	Street          string  `json:"street"`
	HouseNumber     string  `json:"house_number"`
	BuildingNumber  *string `json:"building_number,omitempty"`
	ApartmentNumber string  `json:"apartment_number"`
}

type CompanyAddress struct {
	DeliveryAddressId int     `json:"delivery_address_id"`
	Region            string  `json:"region"`
	City              string  `json:"city"`
	Street            string  `json:"street"`
	HouseNumber       *string `json:"house_number,omitempty"`
	BuildingNumber    *string `json:"building_number,omitempty"`
}
