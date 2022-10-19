package core

type CustomerAccount struct {
	CustomerAccountId int    `json:"customer_account_id"`
	Email             string `json:"email"`
	PasswordHash      string `json:"password_hash"`
}
