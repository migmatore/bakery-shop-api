package core

import "github.com/jackc/pgtype"

type Order struct {
	OrderId           int              `json:"order_id"`
	CustomerId        int              `json:"customer_id"`
	CartId            int              `json:"cart_id"`
	ReceiptId         int              `json:"receipt_id"`
	OrderStatusId     int              `json:"order_status_id"`
	DeliveryAddressId int              `json:"delivery_address_id"`
	DeliveryMethodId  int              `json:"delivery_method_id"`
	OrderDate         pgtype.Timestamp `json:"order_date"`
}
