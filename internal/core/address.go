package core

type DeliveryAddress struct {
	DeliveryAddressId int     `json:"delivery_address_id"`
	Region            string  `json:"region"`
	City              string  `json:"city"`
	Street            string  `json:"street"`
	HouseNumber       string  `json:"house_number"`
	BuildingNumber    *string `json:"building_number,omitempty"`
	ApartmentNumber   *string `json:"apartment_number,omitempty"`
}

type CreateDeliveryAddress struct {
	Region          string  `json:"region"`
	City            string  `json:"city"`
	Street          string  `json:"street"`
	HouseNumber     string  `json:"house_number"`
	BuildingNumber  *string `json:"building_number,omitempty"`
	ApartmentNumber *string `json:"apartment_number,omitempty"`
}

type CreateDeliveryAddressDTO struct {
	Region          string  `json:"region"`
	City            string  `json:"city"`
	Street          string  `json:"street"`
	HouseNumber     string  `json:"house_number"`
	BuildingNumber  *string `json:"building_number,omitempty"`
	ApartmentNumber *string `json:"apartment_number,omitempty"`
}

func NewCreateDeliveryAddressFromDTO(dto *CreateDeliveryAddressDTO) *CreateDeliveryAddress {
	return &CreateDeliveryAddress{
		Region:          dto.Region,
		City:            dto.City,
		Street:          dto.Street,
		HouseNumber:     dto.HouseNumber,
		BuildingNumber:  dto.BuildingNumber,
		ApartmentNumber: dto.ApartmentNumber,
	}
}

type CompanyAddress struct {
	CompanyId      int     `json:"delivery_address_id"`
	Region         string  `json:"region"`
	City           string  `json:"city"`
	Street         string  `json:"street"`
	HouseNumber    *string `json:"house_number,omitempty"`
	BuildingNumber *string `json:"building_number,omitempty"`
}

type CreateCompanyAddress struct {
	Region         string  `json:"region"`
	City           string  `json:"city"`
	Street         string  `json:"street"`
	HouseNumber    *string `json:"house_number,omitempty"`
	BuildingNumber *string `json:"building_number,omitempty"`
}

type CreateCompanyAddressDTO struct {
	Region         string  `json:"region"`
	City           string  `json:"city"`
	Street         string  `json:"street"`
	HouseNumber    *string `json:"house_number,omitempty"`
	BuildingNumber *string `json:"building_number,omitempty"`
}

func NewCreateCompanyAddressFromDTO(dto *CreateCompanyAddressDTO) *CreateCompanyAddress {
	return &CreateCompanyAddress{
		Region:         dto.Region,
		City:           dto.City,
		Street:         dto.Street,
		HouseNumber:    dto.HouseNumber,
		BuildingNumber: dto.BuildingNumber,
	}
}
