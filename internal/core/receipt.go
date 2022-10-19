package core

import "github.com/jackc/pgtype"

type Receipt struct {
	ReceiptId       int         `json:"receipt_id"`
	OrderDate       pgtype.Date `json:"order_date"`
	PaymentDetails  string      `json:"payment_details"`
	PaymentMethodId int         `json:"payment_method_id"`
}
