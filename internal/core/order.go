package core

type Order struct {
	OrderId           int `json:"order_id,omitempty"`
	CustomerId        int `json:"customer_id,omitempty"`
	ProductId         int `json:"product_id,omitempty"`
	ReceiptId         int `json:"receipt_id,omitempty"`
	OrderStatusId     int `json:"order_status_id,omitempty"`
	DeliveryAddressId int `json:"delivery_address_id,omitempty"`
	DeliveryMethodId  int `json:"delivery_method_id,omitempty"`
}
