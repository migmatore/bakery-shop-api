package core

type Order struct {
	OrderId           int `json:"order_id"`
	CustomerId        int `json:"customer_id"`
	ProductId         int `json:"product_id"`
	ReceiptId         int `json:"receipt_id"`
	OrderStatusId     int `json:"order_status_id"`
	DeliveryAddressId int `json:"delivery_address_id"`
	DeliveryMethodId  int `json:"delivery_method_id"`
}
