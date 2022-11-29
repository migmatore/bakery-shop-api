package core

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Customer struct {
	CustomerId        int        `json:"customer_id"`
	FirstName         string     `json:"first_name"`
	LastName          string     `json:"last_name"`
	Patronymic        *string    `json:"patronymic,omitempty"`
	ImagePath         string     `json:"image_path"`
	PhoneNumber       string     `json:"phone_number"`
	Email             string     `json:"email,omitempty"`
	PasswordHash      string     `json:"password_hash,omitempty"`
	DeliveryAddressId *int       `json:"delivery_address_id,omitempty"`
	CartId            int        `json:"cart_id"`
	WishListId        int        `json:"wish_list_id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}

type CreateCustomer struct {
	FirstName         string     `json:"first_name"`
	LastName          string     `json:"last_name"`
	Patronymic        *string    `json:"patronymic,omitempty"`
	ImagePath         string     `json:"image_path"`
	PhoneNumber       string     `json:"phone_number"`
	Email             string     `json:"email,omitempty"`
	PasswordHash      string     `json:"password_hash,omitempty"`
	DeliveryAddressId *int       `json:"delivery_address_id,omitempty"`
	CartId            int        `json:"cart_id"`
	WishListId        int        `json:"wish_list_id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}

type CreateCustomerDTO struct {
	FirstName       string                    `json:"first_name"`
	LastName        string                    `json:"last_name"`
	Patronymic      *string                   `json:"patronymic,omitempty"`
	ImagePath       string                    `json:"image_path"`
	PhoneNumber     string                    `json:"phone_number"`
	Email           string                    `json:"email"`
	Password        string                    `json:"password"`
	DeliveryAddress *CreateDeliveryAddressDTO `json:"delivery_address,omitempty"`
}

type SigninCustomerDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninCustomer struct {
	CustomerId   int
	PasswordHash string
}

func NewCreateCustomerFromDTO(dto *CreateCustomerDTO, deliveryAddressId *int, cartId int, wishListId int) (*CreateCustomer, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	//createdAt, err := time.Parse(time.RFC3339, time.Now())
	//if err != nil {
	//	return nil, err
	//}

	return &CreateCustomer{
		FirstName:         dto.FirstName,
		LastName:          dto.LastName,
		Patronymic:        dto.Patronymic,
		ImagePath:         dto.ImagePath,
		PhoneNumber:       dto.PhoneNumber,
		Email:             dto.Email,
		PasswordHash:      string(hash),
		DeliveryAddressId: deliveryAddressId,
		CartId:            cartId,
		WishListId:        wishListId,
		CreatedAt:         time.Now(),
		UpdatedAt:         nil,
	}, nil
}
