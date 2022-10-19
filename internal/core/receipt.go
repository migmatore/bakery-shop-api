package core

type Receipt struct {
	ReceiptId       int    `json:"receipt_id"`
	OrderDate       string `json:"order_date"`
	PaymentDetails  string `json:"payment_details"`
	PaymentMethodId int    `json:"payment_method_id"`
}
