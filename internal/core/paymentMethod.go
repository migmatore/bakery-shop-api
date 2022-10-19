package core

type PaymentMethod struct {
	PaymentMethodId int    `json:"payment_method_id"`
	Name            string `json:"name"`
}
