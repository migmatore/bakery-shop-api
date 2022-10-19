package core

type Supplier struct {
	SupplierId      int    `json:"supplier_id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	TelephoneNumber string `json:"telephone_number"`
}
