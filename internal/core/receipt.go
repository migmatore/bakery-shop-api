package core

type Receipt struct {
	ReceiptId       int     `json:"receipt_id"`
	PaymentDetails  string  `json:"payment_details"`
	PaymentMethodId int     `json:"payment_method_id"`
	Amount          float32 `json:"amount"`
}
